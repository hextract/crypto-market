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

ENTRYPOINT ["./main", "--port=$MAIN_REST_PORT", "--host=$MAIN_HOST"]


