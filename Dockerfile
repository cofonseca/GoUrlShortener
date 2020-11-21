FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk --no-cache --update add ca-certificates

WORKDIR /rebred

COPY . .

RUN go build .

CMD ./rebred