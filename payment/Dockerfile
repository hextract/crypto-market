FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /payment
COPY payment/go.mod .
COPY payment/go.sum .
RUN go mod download

COPY payment/ .

COPY ../pkg ../pkg

RUN go build -o payment cmd/main.go

EXPOSE $PAYMENT_REST_PORT

ENTRYPOINT ["./payment", "--port=8888", "--host=0.0.0.0"]

