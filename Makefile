run:
	go run cmd/main.go -mode=debug -db=test.db -port=8000
prod:
	go run cmd/main.go -mode=release -db=test.db -port=8000
swagger:
	~/go/bin/swag init -g ./cmd/main.go
dev:
	~/go/bin/air -c .air.toml