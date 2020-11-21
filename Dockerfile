FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

WORKDIR /rebred

RUN apk --no-cache --update add ca-certificates && go build .

CMD ./rebred