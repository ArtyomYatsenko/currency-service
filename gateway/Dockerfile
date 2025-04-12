FROM golang:1.23
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway ./gateway/cmd/gateway/main.go
CMD ["./gateway --config=./currency/configs"]


