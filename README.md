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

`go run main.go tl (--config config.toml) -i インスタンス名`

でローカルTLが引っ張れます。

` go run main.go note (--config config.toml) -i インスタンス名 "こんにちは"`

で「こんにちは」が投稿されます。

--configオプションを指定しない場合は~/.config/misskey-cli.tomlを読むようになっています。

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
  - Renoteかどうか分かる
  - リプライ元の投稿がわかる
  - 投稿時間がわかる
- 投稿関連
  - 投稿できる
    `go run main.go note -i hoge "ねこですよろしくおねがいします"`
    ‐ 投稿するとちゃんとIDが表示される)
  - リプライができる
    `go run main.go note -r 90ab12cd34 "ねこでした"`
  - 投稿の削除ができる
    `go run main.go note -d 90ab12cd34`
  - renoteできる
    `go run main.go renote 90ab12cd34`

## できないこと
- `misskey-cli tl ...`みたいにつかう(未ビルドのため) (リリース時に解決)
- インスタンス関連
  - コマンドから新たにアカウントを追加する (優先度: 最高)
  - デフォルト設定(-iがないときに自動で事前に指定したインスタンスを利用する) (優先度: 最高)
- タイムライン
  - 投稿の詳細(特定の1件)を見る (優先度: 中)
  - リアクションを確認する (優先度: 中)
- 投稿関連
  - リアクションをする/外す (優先度: 中)
- 検索
- ユーザー関連
  - フォロー/アンフォローする (優先度: 中)
  - 特定のユーザーの投稿をみる (優先度: 高)
- Stream APIを利用した自動更新(watchコマンドで擬似的に可能ではある) (優先度: 低)