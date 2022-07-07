# Proxy Generator

## 目標

初期必要設定

- 使用するデフォルトサブネットの ID
- tenant id
- ユーザー名
- パスワード

cli を作成する

## 流れ

ログイン ID とパスワードを検証

インスタンス作成
↓
インスタンスの Port を取得
↓
FloatingIP の作成
↓
FloatingIP をインスタンスにアタッチする
