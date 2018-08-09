# openRPA Log Server
openRPAのログサーバです。

フレームワークはginを使用しています。

## 環境構築

### 必要なライブラリをインポート
`go get -v ./...`

### MySQLサーバを立ち上げる
`mysql -u root -p`

`> create database openRPA_Log_Server;`

### Envを設定
`export OPEN_RPA_LOG_DB_USER="root"`

`export OPEN_RPA_LOG_DB_PASSWORD="password"`

`export OPEN_RPA_LOG_DB_HOST="localhost"`

`export OPEN_RPA_LOG_DB_NAME="openRPA_Log_Server"`

### ログサーバを実行
`go run main.go`

9000ポートがデフォルトになっています。
変更する場合は、main.goの18行目の`r.Run(":9000")`を変更。

## API 仕様書 

postmanを参照してください(https://documenter.getpostman.com/view/4637669/RWThTLBq)
