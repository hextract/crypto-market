FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /stack_connector
COPY stack_connector/go.mod .
COPY stack_connector/go.sum .
RUN go mod download

COPY stack_connector/ .

COPY ../pkg ../pkg

RUN go build -o stack_connector cmd/main.go

EXPOSE $STACK_CONNECTOR_REST_PORT

ENTRYPOINT ["./stack_connector", "--port=8888", "--host=0.0.0.0"]
