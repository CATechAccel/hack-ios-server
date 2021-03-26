# hack-ios-server

### mysqlコンテナへの接続方法
```shell
docker-compose exec server sh -c 'mysql -u root -phack-ios-server -h mysql -P 3306 --protocol=tcp'  
```