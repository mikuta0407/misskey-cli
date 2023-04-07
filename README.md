# misskey-cli

MisskeyのCLIクライアントです。

まだ制作中ですが、とりあえず使えるようになりました。

![misskey-cli](https://user-images.githubusercontent.com/13357430/194720200-8dbf0394-9d4b-4e84-ad91-739eb0fec1c4.png)

[twty](https://github.com/mattn/twty)が便利だったのでMisskeyでも似たようなことをしたくて作りました。

## とりあえずの使い方

### 事前準備

1. `misskey-cli.toml` を作成します。以下はデフォルトの場所(オプション不要で読まれる場所)
  - Windows: `%APPDATA%\misskey-cli.toml`
  - Linux: `~/.config/misskey-cli.toml`
  - macOS: `~/mikuta0407/Library/Application Support/misskey-cli.toml`
1. misskeyのサイトで、設定→その他設定→API→アクセストークンの作成 で、APIのアクセストークンを作成します(コピーして下さい)
2. 以下の内容にします。(token欄に2番で作成したAPIアクセストークンを入力します)
  ```
  [[Instance]]
    host = "https://example.com"
    name = "インスタンス名(自由)"
    token = "API Token"
  ```

  複数インスタンス書く場合はこれを羅列します。(tomlのテーブルです)

### コマンド

- タイムラインを見る: `tl`
  - `./misskey-cli tl -i インスタンス名` (インスタンス名はtomlの「name」項目の値)
    - `-m local`でローカルTL(デフォルト)
    - `-m global`でグローバルTL
    - `-m home`でホームTL
    - `-l 20`で20件表示(デフォルト10件)
- ストリーミング: `stream`
  - `./misskey-cli stream -i hogeinstance local`
- 新規投稿/リプライ/投稿の削除: `note`
  - `./misskey-cli note -i hogeinstance ねこです`: hogeinstanceインスタンスに「ねこです」と投稿
  - `./misskey-cli note -i hogeinstance -r 9012abcdef "よろしくおねがいします。 ねこでした`: ノートIDが9012abcdefの投稿に返信
  - `./misskey-cli note -i hogeinstance -d 9012abcdef`: ノートIDが9012abcdefの投稿を削除
- リノート: `renote`
  - `./misskey-cli renote -i hogeinstance 9012abcdef`: ノートIDが9012abcdefの投稿をリノート


### その他挙動について

- `-i`でインスタンスを指定しなかった場合は、toml内の一番上のインスタンスを自動的に利用します
- `--config`でtomlファイルを指定できます。指定がない場合に`~/.config/misskey-cli.toml`を読むようになっています。
- 未知の挙動があるかもしれません。サーバーを破壊することは無いと思いますが、責任は負いません。

## できること・できないことまとめ

### できること

- インスタンス設定関連
  - 接続情報を保存して利用する(tomlで保存)
  - インスタンスを切り替えて操作する
    - `-i`でtoml内の`name`で指定した名前を指定します
- タイムライン
  - **Stream APIを用いてタイムラインをストリーミングで見る**
  - homeタイムラインを見る
  - localタイムラインを見る
  - globalタイムラインを見る
  - 投稿主がローカルか他サーバかがわかる
  - 一度に表示する数の制御ができる
  - 添付があるかないかわかる
  - 投稿がRenoteかどうか分かる
  - リプライ元の投稿がわかる
  - 投稿時間がわかる
- 投稿関連
  - 投稿する
    - publicのみ
  - リプライする
  - 投稿の削除をする
  - renoteする
 
## まだできないこと

- ~~Windows(cmd/ps)での動作~~ ← mattnさん、yulogさんのおかげで動作するようになりました!!!
- インスタンス設定関連
  - コマンドから新たにアカウントを追加する (優先度: 最高)
- タイムライン関連
  - 投稿の詳細(特定の1件)を見る (優先度: 中)
  - リアクションを確認する (優先度: 低)
  - Unicode対応絵文字を表示する (優先度: 中)
- 通知(全部)
- 投稿関連
  - リアクションをする/外す (優先度: 中)
  - 公開範囲を設定をする (優先度: 低)
  - NSFW設定をする (優先度: 低)
  - ファイルを投稿する (優先度: 低)
- 検索(全部)
- ユーザー関連(全部)
  - フォロー/アンフォローする (優先度: 中)
  - 特定のユーザーの投稿をみる (優先度: 低)
  - その他全部
- その他「できること」以外の内容
