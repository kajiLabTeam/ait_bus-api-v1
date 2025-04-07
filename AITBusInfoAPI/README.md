# AIT-BusInfo-backend

envファイルをリポジトリのrootディレクトリに置く

.envファイル
```
GOLANG_PORT=8082
MYSQL_PORT=33066
GOLANG_CONTAINER_NAME=vol_golang
MYSQL_CONTAINER_NAME=vol_mysql
```


ネットワーク作成
```
docker network create vol_network
```

go+mysqlのコンテナ起動
```
make dev
```

swaggerの更新
```
swag init --parseDependency --parseInternal
```