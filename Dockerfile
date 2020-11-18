FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /rebred

COPY . .

CMD go run main.go middleware.go