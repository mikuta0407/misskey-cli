# misskey-cli

まだ制作中ですが、

現在のコミットでは、

tomlファイルに
```
[[Instance]]
  host = "https://example.com"
  name = "インスタンス名(自由)"
  token = "API Token"

```

と書いて準備した後、

`go run main.go tl -c config.toml -i インスタンス名`

でTLが引っ張れます。

` go run main.go note -c config.toml -i インスタンス名 "こんにちは"`

で「こんにちは」が投稿されます。

-cオプションは、将来的には指定無しの場合は~/.config/misskey-cli.tomlを読むようにしたいなと思っています。

## できること
- インスタンス関連
  - インスタンスを切り替えて操作する
  - インスタンスの情報をtomlから引っ張ってくる
- タイムライン
  - homeタイムラインを見る
    `go run main.go tl -i hoge -m home`
  - localタイムラインを見る
    `go run main.go tl -i hoge -m local`
  - globalタイムラインを見る
    `go run main.go tl -i hoge -m global`
  - 投稿主がローカルか他サーバかがわかる
  - 一度に表示する数の制御ができる
    `go run main.go tl -i hoge -l 20`
  - 添付があるかないかわかる
- 投稿関連
  - 投稿できる
    `go run main.go note -i hoge "ねこですよろしくおねがいします"`
    ‐ 投稿するとちゃんとIDが表示される)

## できないこと
- `misskey-cli tl ...`みたいにつかう(未ビルドのため)
- インスタンス関連
  - コマンドから新たにアカウントを追加する
- タイムライン
  - 投稿の詳細を見る
- 投稿関連
  - リプライする
  - renoteする
- 検索
- ユーザー関連
  - フォローする
- Stream APIを利用した自動更新(watchコマンドで擬似的に可能)