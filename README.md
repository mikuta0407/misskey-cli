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