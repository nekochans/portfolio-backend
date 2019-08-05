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

## Docker Hubに反映する

`docker-push.sh` を実行して下さい。

[![dockeri.co](https://dockeri.co/image/nekochans/portfolio-backend-go)](https://hub.docker.com/r/nekochans/portfolio-backend-go)

[![dockeri.co](https://dockeri.co/image/nekochans/portfolio-backend-nginx)](https://hub.docker.com/r/nekochans/portfolio-backend-nginx)

## gcr（Container Registry）に反映する

`docker-push-to-gcr.sh` を実行して下さい。
