.SILENT: #Убираем вывод самой команды в консоль

lint:
	golangci-lint run ./...

run cron:
	go run ./currency/cmd/cron/main.go
