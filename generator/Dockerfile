FROM golang:1.24-alpine

WORKDIR /generator
COPY generator/ .
RUN go mod tidy
RUN go mod download

RUN go build -o generator cmd/main.go

EXPOSE $GENERATOR_REST_PORT
EXPOSE $MATCHING_ENGINE_PORT

ENTRYPOINT ./generator --port=$GENERATOR_REST_PORT --host=$GENERATOR_HOST