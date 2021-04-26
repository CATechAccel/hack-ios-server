# hack-ios-server
「hack ios server」 のサーバサイドを管理するレポジトリです．

クライアントサイドのレポジトリは[こちら](https://github.com/CATechAccel/hack_iOS)にあります．

API仕様は[こちら](https://github.com/CATechAccel/hack-ios-server/blob/main/api-document.yml)にあります．

## アプリケーションアーキテクチャ
レイヤードアーキテクチャにポートを追加して，技術とアプリケーションを分離するようなアーキテクチャを採用しています．

アプリケーションアーキテクチャ仕様は [docs/architecture.md](https://github.com/CATechAccel/hack-ios-server/blob/main/docs/architecture.md) にあります．

## mysqlコンテナへの接続方法
```shell
docker-compose exec server sh -c 'mysql -u root -phack-ios-server -h mysql -P 3306 --protocol=tcp'
```