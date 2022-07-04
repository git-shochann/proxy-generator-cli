# Proxy Generator

Web 上で NHNCloud の API を用いて自動でインスタンスの作成、全件取得、指定したインスタンスを開始、停止
またグローバル IP のアタッチ、デタッチを行います。
グローバル IP がアタッチされているサーバーに関しては、SSH 接続の上、Squid を用いて Proxy サーバーの自動生成を行います。

## 各ディレクトリ、ファイルの構成

```Markdown
.
├── cmd
│   └── main.go // エントリーポイント
├── configs
│   ├── config.go
│   └── config.ini
├── docs
│   ├── README.md
│   ├── memo.md
├── go.mod
├── go.sum
├── image-list.txt
├── internal
│   ├── floating-ip.go // floatingIP 関連
│   ├── image.go // imageID を取得
│   ├── instance.go // インスタンスを生成, 停止, 取得
│   └── token.go // API 使用準備, token 生成
├── makefile // 便利コマンド群
└── pkg
└── random.go // ランダムの名前を生成
```
