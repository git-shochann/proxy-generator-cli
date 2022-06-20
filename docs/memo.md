# Proxy Generator

## MEMO

token.go -> API 呼び出し時に必要なトークンを発行。
instance.go -> 1 個のインスタンスを作成する。
floating_ip.go -> グローバル IP をアタッチする。
proxy.sh -> サーバーに ssh 接続を行い、Proxy サーバーを構築する。

IP:22:centos にて ssh キーを用いて接続 -> ドキュメント確認

## 目標にしたい

ユーザー側で選んで作成できる Automata の Generator のような機能に近づける。

初期必要設定としては

- 使用するデフォルトサブネットの ID
- tenant id
- ユーザー名
- パスワード
