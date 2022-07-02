# Proxy Generator

CLI にて NHN Cloud を用いて自動でインスタンスを作成し、グローバル IP をアタッチ
自動で Proxy サーバーを生成し、最終的にブラウザでも使用出来るように IP:Port:User:Pass にして使えるようにします。

## 各ディレクトリ、ファイルの構成

```
.
├── cmd
│   └── main.go // エントリーポイント
├── configs
│   ├── config.go
│   └── config.ini
├── docs
│   ├── README.md
│   ├── memo.md
│   └── question.md
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
