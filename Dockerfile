FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /rebred

COPY . .

RUN go build . && apk --no-cache --update add ca-certificates

CMD go run .