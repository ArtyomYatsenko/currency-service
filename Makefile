.PHONY: lint cron gateway

lint:
	@golangci-lint run ./...

cron:
	@go run ./currency/cmd/cron/main.go

gateway:
	@go run ./gateway/cmd/gateway/main.go
