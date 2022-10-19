# Proxy Generator

## Feature

NHNCloud にて自動でインスタンスの作成、全件取得、指定したインスタンスを開始、停止などを行うことが可能。
またインスタンスに関してはグローバル IP のアタッチ、デタッチを自動で行います。
グローバル IP がアタッチされているサーバーに関しては、SSH 接続の上、Squid を用いて Proxy サーバーの自動生成を行います。

## Configuration

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

## Requiring Setting

1, 以下の形式に沿った Yaml ファイルを作成する必要があります。

```Yaml
credentials:
  tenantid: <実際のインスタンスの設定から取得>
  username: <NHN CloudのログインID>
  password: <TenantIDに対するPassWord>
  keyname: <ssh接続するためのキー名, NHN Cloud Consoleから取得する>
```

2, ssh 接続を許可するために NHN Cloud のサイトにて 使用する Default のセキュリティグループの 22 ポートを解放する必要があります。

## How To Use

1, 起動後.pem(秘密鍵をドラッグ&ドロップでスタート) // 適切な場所に置かないと権限エラーになる 場所を指定した方がいいかもしれない

```Shell
  proxy-generator-cli create [your nhn cloud private key]
```
