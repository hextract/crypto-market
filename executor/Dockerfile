FROM golang:1.23-alpine

WORKDIR /executor
COPY executor/ .
RUN go mod tidy
RUN go mod download

RUN go build -o executor cmd/main.go

EXPOSE $EXECUTOR_REST_PORT

ENTRYPOINT ./executor --port=$EXECUTOR_REST_PORT --host=$EXECUTOR_HOST