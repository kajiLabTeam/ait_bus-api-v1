include .env

# コンテナ名 envファイルから取得
DB_CONTAINER_NAME = $(MYSQL_CONTAINER_NAME)
# データベース名
DB_NAME = app

network:
	docker network create vol_network
# 開発環境
dev:
	docker-compose up 
dev-d:
	docker-compose up -d
# 本番環境
prod:
	docker-compose up -d vol_golang