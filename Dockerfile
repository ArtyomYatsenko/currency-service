FROM golang:1.23
WORKDIR /app
COPY . .
RUN go build -o cron ./currency/cmd/cron/main.go
CMD ["./cron"]