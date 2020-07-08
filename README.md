# portfolio-backend
![ci-master](https://github.com/nekochans/portfolio-backend/workflows/ci-master/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/nekochans/portfolio-backend/badge.svg?branch=master)](https://coveralls.io/github/nekochans/portfolio-backend?branch=master)

GitHub Organization 「nekochans」の説明用Webサイトのバックエンド

# 環境変数の設定

以下の環境変数を設定する必要があります。

```
export GCP_PROJECT_ID="作成したGCPのProjectID"
export DB_USER="nekochans"
export DB_PASSWORD="nekochans(Password2222)"
export DB_NAME="portfolio_backend"
export DB_HOST="portfolio-backend-mysql"
export TEST_DB_USER="nekochans_test"
export TEST_DB_PASSWORD="nekochans(Password2222)"
export TEST_DB_NAME="portfolio_backend_test"
```

これらのパスワードは本番のDBでは決して利用しないで下さい。

[direnv](https://github.com/direnv/direnv) 等を利用すると良いでしょう。

## ローカル上でDockerで動作させる

以下のスクリプトを実行して下さい。

`./docker-compose-up.sh`

[この記事](https://qiita.com/keitakn/items/f46347f871083356149b) のように `delve` を使ってデバックを行う場合は以下のスクリプトを実行して下さい。

`./docker-compose-up-debug.sh`

### マイグレーションの実行

各コンテナが起動した状態で以下を実行して下さい。

```
# データベースにマイグレーションの実行
docker-compose exec go make migrate-up

# マイグレーションをロールバックする
# Are you sure you want to apply all down migrations? [y/N] と2回聞かれるので y でEnterして下さい
docker-compose exec go make migrate-down
```

詳しくは [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) のドキュメントを参照して下さい。

## Docker Hubに反映する

`docker-push.sh` を実行して下さい。

[![dockeri.co](https://dockeri.co/image/nekochans/portfolio-backend-go)](https://hub.docker.com/r/nekochans/portfolio-backend-go)

[![dockeri.co](https://dockeri.co/image/nekochans/portfolio-backend-nginx)](https://hub.docker.com/r/nekochans/portfolio-backend-nginx)

## gcr（Container Registry）に反映する

`docker-push-to-gcr.sh` を実行して下さい。

## ソースコードのフォーマット

`docker-compose exec go sh` でアプリケーション用のコンテナに入ります。

`make lint` を実行して下さい。

もしくは `docker-compose exec go make lint` でも実行出来ます。

lintのルール等は以下を参考にして下さい。

https://golangci-lint.run/usage/linters/

内部でソースコードのフォーマットも行っていますが、自動で修正されない物は自分で修正を行う必要があります。

## テストの実行

`docker-compose exec go sh` でアプリケーション用のコンテナに入ります。

`make test` を実行します。

もしくは `docker-compose exec go make test` でもテストを実行出来ます。

## データベース（MySQL）への接続方法

お使いのPCにMySQLクライアントをインストールして以下のコマンドで接続して下さい。

```
mysql -u root -h 127.0.0.1 -p -P 63306
```

本番環境ではMySQLではなく、AWS、GCPのマネージドサービスを利用します。

## OpenAPIスキーマの更新

APIはスキーマ駆動で開発を行っています。

スキーマは [こちら](https://github.com/nekochans/nekochans-openapi/blob/master/docs/portfolio/openapi.yaml) で管理されています。

スキーマの更新があった場合以下のコマンドでインターフェースの更新を行う必要があります。

```
oapi-codegen -generate types docs/openapi/docs/portfolio/openapi.yaml > infrastructure/openapi/Model.gen.go && \
oapi-codegen -generate chi-server docs/openapi/docs/portfolio/openapi.yaml > infrastructure/openapi/Server.gen.go && \
```
