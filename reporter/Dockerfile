FROM golang:1.23-alpine

WORKDIR /reporter
COPY reporter/ .
RUN go mod tidy
RUN go mod download

RUN go build -o reporter cmd/main.go

EXPOSE $REPORTER_REST_PORT

ENTRYPOINT ./reporter --port=$REPORTER_REST_PORT --host=$REPORTER_HOST