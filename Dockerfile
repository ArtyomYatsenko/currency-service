FROM golang:1.23
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cron ./currency/cmd/cron/main.go
CMD ["./cron"]


