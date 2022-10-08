# Code Level Metrics Example

New Relic CodeStream を使って Code Level Metrics を体験できるリポジトリになっています。

## Requirements
以下の環境や準備が必要になっています。

- Go 1.19.x

- Docker

- [CodeStream Plugin](https://www.codestream.com/)

- New Relic アカウント [Signup Free](https://newrelic.com/signup)

- ※Optional [k6](https://k6.io/)

## How to Use

```shell
export NEW_RELIC_LICENSE_KEY=<your newrelic licensekey>
https://docs.newrelic.com/docs/apis/intro-apis/new-relic-api-keys/
```

```shell
docker compose up --build
```

```shell
curl localhost:8000/example
```

```shell
#ある程度の負荷をかけたい場合にご利用ください
k6 run k6/example.js
```
