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
