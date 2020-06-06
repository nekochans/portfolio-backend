# portfolio-backend
GitHub Organization 「nekochans」の説明用Webサイトのバックエンド

# 環境変数の設定

以下の環境変数を設定する必要があります。

```
export GO111MODULE=on
export GCP_PROJECT_ID=作成したGCPのProjectID
```

[direnv](https://github.com/direnv/direnv) 等を利用すると良いでしょう。

## ローカル上でDockerで動作させる

以下のスクリプトを実行して下さい。

`./docker-compose-up.sh`

[この記事](https://qiita.com/keitakn/items/f46347f871083356149b) のように `delve` を使ってデバックを行う場合は以下のスクリプトを実行して下さい。

`./docker-compose-up-debug.sh`

### マイグレーションの実行

`docker-compose exec go sh` でアプリケーション用のコンテナに入ります。

`/go/app` で以下を実行します。

```
# データベースにマイグレーションの実行
migrate -source file://./_sql -database 'mysql://nekochans:nekochans(Password2222)@tcp(portfolio-backend-mysql:3306)/portfolio_backend' up

# テスト用のデータベースにマイグレーションの実行
migrate -source file://./_sql -database 'mysql://nekochans_test:nekochans(Password2222)@tcp(portfolio-backend-mysql:3306)/portfolio_backend_test' up
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

lintのルール等は以下を参考にして下さい。

https://golangci-lint.run/usage/linters/

内部でソースコードのフォーマットも行っていますが、自動で修正されない物は自分で修正を行う必要があります。

## テストの実行

`docker-compose exec go sh` でアプリケーション用のコンテナに入ります。

`make test` を実行します。

## データベース（MySQL）への接続方法

お使いのPCにMySQLクライアントをインストールして以下のコマンドで接続して下さい。

```
mysql -u root -h 127.0.0.1 -p -P 63306
```

本番環境ではMySQLではなく、AWS、GCPのマネージドサービスを利用します。
