FROM golang:1.23-alpine

WORKDIR /receiver
COPY receiver/ .
RUN go mod tidy
RUN go mod download

RUN go build -o receiver cmd/main.go

EXPOSE $RECEIVER_REST_PORT

ENTRYPOINT ./receiver --port=$RECEIVER_REST_PORT --host=$RECEIVER_HOST