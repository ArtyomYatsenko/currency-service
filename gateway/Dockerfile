FROM golang:1.23
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gatewayApp ./gateway/cmd/gateway/main.go
CMD ["./gatewayApp --config=./currency/configs"]


