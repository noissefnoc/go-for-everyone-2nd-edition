# みんなのGo言語 第2版

## 1章 Goによるチーム開発の始め方とコードを書く上での心得

### 1.1 開発環境の構築

#### REPL

[motemem/gore](https://github.com/motemem/gore) を使う。補助として
* [mdempsky/gocode](https://github.com/mdempsky/gocode)
* [k0kubun/pp](https://github.com/k0kubun/pp)

も使うと良い

#### Goプロジェクトコード管理

[motemen/ghq](https://github.com/motemen/ghq) を使う (すでに使ってた)。


#### Goプロジェクトコード検索

[peco/peco](https://github.com/peco/peco) を使う。

Bashの設定だと

```bash
function ghql() {
  local selected_file=$(ghq list --full-path | peco --query "$LBUFFER")
  if [ -n "$selected_file" ]; then
    if [ -t 1 ]; then
      echo ${selected_file}
      cd ${selected_file}
      pwd
    fi
  fi
}

bind -x '"\201": ghql'
bind '"\C-g":"\201\C-m"'
```

あたりを入れる。


### 1.2 エディタと開発環境

この章は Goland 使っているので、Goland の公式ドキュメントを読むようにする。


### 1.3 Go を始める

#### Tour of Go と公式ドキュメント

Tour of Go と 公式ドキュメントをまず当たるように書いている。

#### プロジェクトを始める & ディレクトリ構成

この辺、 Perl でいうところの `Module::Starter` みたいなのがあると嬉しいんだけど、 [Songmu/godzil](https://github.com/Songmu/godzil) のようなオーサリングツールを現在模索中。


### 1.4 Goらしいコードを書く

* `error` をちゃんと使う。1.14で `try` の導入見送りになったっぽいので多値返却でいく
* 正規表現はできるだけ使わず、使っても `regexp.MustCompile` で確定させておく
* map を避ける。できるだけ構造体を使って型を定義する
* reflect を避ける。できるだけ型をつける
* 巨大な struct を作らず継承させようとしない。これは API の JSON Response を分割して作る時に思う
* 並行処理を使いすぎない
* Go のコードを読もう
    * Go のバージョンは古いけど「[GoのためのGo](https://motemen.github.io/go-for-go-book/https://motemen.github.io/go-for-go-book/)」も参考になるかも
* テストとCI。 `go vet` や `golint` などでのチェックを入れる
* ビルドとデプロイ。ビルド時の埋め込みやフラグ分岐など
* モニタリング


## 2章 マルチプラットフォームで動作する社内ツールの作り方

私がWindows 持ってないので、この章の話ちゃんと検証できてないところがある。


### 2.1 Goで社内ツールを作る理由

* ランタイムが不要なので配布しやすいが、その分クロスコンパイルに考慮する必要がある
* cgoやOS間の差異を吸収しない記述などがあるので、その点に気をつける必要がある

### 2.2 守るべき暗黙のルール

* OS内のファイルパス (`path/filepath`)とURLのパス(`path`)の違いに注意
    * 間違えると場合によってディレクトリトラバーサルなどを起こす場合がある
* パスセパレーターがWindowsとmacOS/Linuxで異なるのでそこに注意
* Windowsのファイル操作に関するタイミングが異なって `defer` のタイミングを厳密にした方がいい
    * この話、 `ioutil/TempFile` でツイート見た記憶がある
* 基本的に内部処理はUTF-8を使う
    * これは昔BSD(EUC-JP)とLinux(UTF-8)環境が混在してたときに確かにやった
    
### 2.3 TUIもWindowsで動かしたい

Windows OS持ってないので確認できず

### 2.4 OS固有の処理への対応

* OSの分岐を見る環境変数や `runtime` パッケージ
* ファイル名のサフィックスに対応しているOSやアーキテクチャ名をつけておくと、 `go build` 時に対応した環境ファイルでビルドされる
* ファイル中のコメントでも `// +build` とすることで、対応を分岐することもできる
* `pkg-config` を使うオプションもある

### 2.5 頑張るよりも周りのツールに頼る

* デーモナイズの話
    * Perlをdaemon toolsで起動してた時代を思い出す
* Windowsの場合は [nssm](https://nssm.cc/) を使うらしい

### 2.6 シングルバイナリにこだわる

* 静的ファイルをバイナリに含める方法。以下が紹介されている
    * [rakyll/statik](https://github.com/rakyll/statik)
    * [packr/packr](https://github.com/packr/packr)

### 2.7 Windowsアプリケーションの作成

Windows持ってないので試せてません。

### 2.8 設定ファイルの取り扱い

* JSONかYAMLで作る
* XDG Base Directory Specificationによる規定で `$HOME/.config/(アプリケーション名)` に配置する
* Go 1.12からは `os.UserHomeDir()` でホームディレクトリがcgoの状態に関わらず取れる
* が、Windowsの設定ファイルは `%APPDATA%` 配下にあったほうがよいので、実行環境と環境変数を

### 2.9 社内ツールのその先に

OSS化できたらいいなぁ


## 3章 実用的なアプリケーションを作るために

### 3.1 はじめに

実用的なアプリケーションの話

* どのような機能を持っているかが容易に調べられる
* パフォーマンスが良い
* 多様な入出力を扱える
* 人間にとって扱いやすい形式で入出力できる
* 想定外の場合に安全に処理を停止できる

仕事で時間をかけて作らないと対応しづらいのと、個人のユースケースだけだと不測の事態があまり出づらいのがなぁ。

### 3.2 バージョン管理

バージョンの埋め込み。`-ldflag` と `-X` で該当のパッケージ変数をビルド時に変更できる仕組みで変更可能。

### 3.3 効率的なI/O処理

* バッファリングの話。比較的低級な言語でないと使う機会がないのでつい忘れがち

### 3.4 乱数を使う

* 精度が高い乱数が欲しいなら `crypto/rand` を使う(サンプルコード略)

### 3.5 人間が扱いやすい形式の数値

* [dustin/go-humanize](https://github.com/dustin/go-humanize) の紹介 (サンプルコード略)
* byte変換や時間、順序数、3桁カンマなど

### 3.6 Goから外部コマンドを実行する

* `os/exec` を使うが、 `Output()` , `CombineOutput()` を使う標準の方法だと以下制約がある
    * 出力がコマンドが実行した後に一度にまとめてメモリ上に返される
    * コマンドに対して標準入力を与えることができない
* 入出力ブロックしないように処理を進めていくには標準入出力と標準エラー出力でそれぞれ goroutine を動かして、それぞれがブロックしないようにするなど考慮する
* 外部コマンド文字列を入力としてとる場合のバリデーション用のパーザとして `mattn/go-shellwords` がある

### 3.7 タイムアウトする

* `context` パッケージを使う

### 3.8 goroutine の停止

この辺あまりよく理解できていない。いくつかサンプル書いてテスト

* channelを使用する方法。一度closeしたチャンネルをサイド使おうとするとpanicが起こるので注意
* contextを使用する方法。selectして確認する

### 3.9 シグナルを扱う

* チャンネルで受け取る
* `os.Signal` インターフェースを持つものを作れば独自のシグナル定義もできる


## 4章 コマンドラインツールを作る

### 4.1 なぜGoでCLIツールを書くのか？

* 配布のしやすさ
* 複数プラットフォームへの対応のしやすさ
* パフォーマンス

ランタイムなしにワンバイナリで配布ができるのは作者として楽ではあるのだけど、ライブラリを使っていると知らないうちに cgo 周りで Windows 動かない場合があって難しいところがある。

### 4.2 デザイン

* CLIツールのインターフェース
    * シングルコマンドパターン：標準の `flag` パッケージを使う
    * サブコマンドパターン：サードパーティーのものを使う。詳しくは 4.4
* リポジトリ構成
    * バイナリをメインにするのか、ライブラリをメインにするのか
    * メインにする方が第一階層
        * ライブラリがメインで、コマンドがサブの場合は `PROJECT_ROOT/cmd/COMMAND_NAME/main.go` として配置

### 4.3 flag パッケージ

* 標準パッケージの `flag` パッケージの紹介
* `flag.Type` (ポインタ渡し) または `flag.TypeVal` (値渡し) をした後に `flag.Parse()` で値を取得
* ショートオプションとロングオプションを設定したい場合は両方とも記述が必要
* フラグの記載場所。著者は非パッケージスコープでの定義を推奨
* `flag` の設定を変えることで出力先を変更するなどが可能なので、テストのときに利用する
* コマンドライン引数をパースするような型をカスタムで作成することも可能(カンマ区切り値が例に挙げられている)
* サードパーティーのフラグのパッケージとして以下が例示されている
    * [spf13/pflag](https://github.com/spf13/pflag)
    * [jessevdk/go-flags](https://github.com/jessevdk/go-flags)
    * [alecthomas/kingpin](https://github.com/alecthomas/kingpin)

### 4.4 サブコマンドを持ったCLIツール

* サードパッケージ製のパッケージ紹介
    * [urfave/cli](https://github.com/urfave/cli)
    * [spf13/cobra](https://github.com/spf13/cobra)
    * [docopt/docopt.go](https://github.com/docopt/docopt.go)
    * [mitchellh/cli](https://github.com/mitchellh/cli)
    * [google/subcommands](https://github.com/google/subcommands)
* 著者の使っている [mitchellh/cli](https://github.com/mitchellh/cli) を詳しく説明
    * サブコマンドをインターフェースとして定義

### 4.5 使いやすく、メンテナンスしやすいツール

* 使いやすいツール
    * ステータスコードを返却する。 `os.Exit(statusCode int)` を使う
    * `log.Fatal` など、 `defer` の期待通りの挙動にならないものがある
* 標準出力と標準エラー出力
    * `io.Writer` を引数に取る `fmt` などの関数で出力先を変更できる
* エラーは `error` を介してコンテキストを付与する
* メンテナンスしやすいツール
    * CLIのテスト
        * ステータスコードの確認
        * メッセージの出力
    * `*bytes.Buffer` に書き出すことでテストしやすくする


## 5章 The Dark Arts Of Reflection

### 5.1 動的な方の判別

* 型の検出と型アサーションの限界
    * 型アサーションだと、想定しうる型を全て列記しなければならない
    * これに操作するには `reflect` を使う

### 5.2 `reflect` パッケージ

* `reflect` パッケージを使って型を判定し、型にあった対応をする
* 異なった型に対する処理を実行すると `panic` が起こる
* `reflect.StructTag` で `struct` のタグを判定できる。
    * JSONなんかのタグはこんな感じでパースされてるのか

### 5.3 reflect の利用例

* 値を動的に生成したり、スコープの話だったり
* 「ポインタとinterfaceの値」の「Setできる値」のところ、前節でpanic起こしてたコードだ
* 動的に `select` を構築することができる

### 5.4 reflect のパフォーマンスとまとめ

* reflectは実行時に逐次確認をするので、パフォーマンスが悪くなる
* 本当に必要なときにしか使わない方がよい


## 6章 Goのテストに関するツールセット

### 6.1 Goにおけるテストのあり方

* `testing` パッケージの使い方

### 6.2 testing パッケージ入門

* テストファイルにおける命名規則やマジックコメントについて

### 6.3 ベンチマーク入門

* `go test -bench` の使い方と見方

### 6.4 テストの実践的なテクニック

* Table Driven Test
    * 入力と期待値を持つ構造体の配列のテーブルをループしてテストを実行する
* 構造体の比較は `reflect.DeepEqual` で行う
* `-race` フラグでRace Detector を使う
* `TestMain` を使うことで `xUnit` の `startUp` / `shutDown` のような機能を使う
* マジックコメントでタグをつけることができる。テストの種類を分けるのにも使える
* テストにおける変数または手続きの置き換え (パッケージ変数の置き換え)
* インターフェースのモッキング
* `net/http/httptest` パッケージ
* テストカバレッジは `-coverprofile` や `-covermode` を使う


## 7章 データベースの取り扱い

### 7.1 Go におけるデータベースの取り扱い

* `database/sql` の働き

### 7.2 database/sql を使ってデータベースに接続する

* [lib/pq](https://github.com/lib/pq) を使ってデータベースに接続

### 7.3 Exec 命令の実行

* プレースホルダーの話
* 実行結果の話

### 7.4 Query 命令の実行

* 行イテレーターが帰ってくるので、ループしながら割り当て

### 7.5 Go におけるデータベースの型

* データベースカラムの型のGoの型へのマッピングの話

### 7.6 ORMを使ったデータベースの扱い方

* 次の節で [go-gorp/gorp](https://github.com/go-gorp/gorp) を使うための説明

### 7.7 REST サーバを作る

* Webフレームワークとして [labstack/echo](https://github.com/labstack/echo) を使う
* Validatorとして [go-playground/validator.v9](https://gopkg.in/go-playground/validator.v9) を使う
 