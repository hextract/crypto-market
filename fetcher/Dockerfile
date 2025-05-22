FROM golang:1.23-alpine

WORKDIR /fetcher
COPY fetcher/ .
RUN go mod tidy
RUN go mod download

RUN go build -o fetcher cmd/main.go

EXPOSE $FETCHER_REST_PORT

ENTRYPOINT ./fetcher --port=$FETCHER_REST_PORT --host=$FETCHER_HOST