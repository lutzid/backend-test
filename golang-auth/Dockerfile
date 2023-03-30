FROM golang:alpine

ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache git \
    gcc \
    musl-dev

RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base

/bin/sh -c go get github.com/mattn/go-sqlite3

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o binary

ENTRYPOINT ["/app/binary"]