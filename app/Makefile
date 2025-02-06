include .env
migrate-up:
	migrate --path database/migrations --database 'postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' -verbose up
migrate-down:
	migrate --path database/migrations --database 'postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' -verbose down
migrate-create:
	migrate create -ext sql -dir database/migrations -seq $(SEQ_NAME)
