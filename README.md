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

