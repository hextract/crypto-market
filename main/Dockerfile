FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /main
COPY main/go.mod .
COPY main/go.sum .
RUN go mod download

COPY main/ .

COPY ../pkg ../pkg

RUN go build -o main cmd/main.go

EXPOSE $MAIN_REST_PORT

ENTRYPOINT ["./main", "--port=8880", "--host=0.0.0.0", "--tls-certificate", "server.crt", "--tls-key", "server.key"]


