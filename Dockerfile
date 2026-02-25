FROM golang:1.24-alpine

WORKDIR /workspace

COPY . /workspace

RUN go mod download && go build -o redis-demo ./cmd/server

EXPOSE 8080

CMD ["./redis-demo"]
