# task_management
これはgolang framework "echo"のチュートリアルとして作成したタスク管理システムである。

## 環境
- Mac OS X 10.14.6（Mojave）
- Docker version 19.03.5
- docker-compose version 1.24.1

### go
- golang:1.13.1-alpine

### Vue.js
- vue-cli3
- aixos

### nginx
`nginx.conf`は、`./etc/nginx/nginx.conf`へコピー・編集にて活用する。
80ポートをlistenしたらリバースプロキシとして`http://go:8082`、あるいは、`http://vue:8080`にパスします。これで、Nginx経由でwebアプリにアクセスします。
なお、リクエストのurlパスに`/api/`が含まれている時に、`http://go:8082`へのアクセスが適用されます。
