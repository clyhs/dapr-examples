go build -o de-hello-service
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o de-hello-service

swag init -o ./swagger

docker build -t abigfish/hello-service .

docker run -d --name=de-hello-service -e CONFIG_NAME="config" -p 3000:3000 -v /Users/chenliyu/go/src/dapr-examples/hello-service/config:/config abigfish/hello-service

dapr run --app-id hello-service --app-protocol grpc --app-port 3000 --dapr-http-port 3501 --log-level debug --components-path ./config go run main.go
dapr run --app-id hello-service --app-protocol grpc --app-port 3000 --dapr-http-port 3501 --log-level debug --components-path ./config ./de-hello-service

http://127.0.0.1:3501/v1.0/invoke/hello-service/method/swagger.json
```
version: "3"
services:
  de-hello-service-dapr:
    image: daprio/daprd:edge
    network_mode: "host"
    command: ["./daprd",
        "-app-id", "hello-service",
        "-app-protocol", "grpc",
        "-app-port", "3000",
        "-dapr-grpc-port", "3801",
        "-dapr-http-port", "3501",
        "-enable-metrics=false",
        "-log-level", "info",
        "-placement-host-address", "placement:50005"]
    volumes:
        - "/Users/chenliyu/go/src/dapr-examples/hello-service/components/:/components"
        - "/etc/localtime:/etc/localtime:ro"
    restart: always
    depends_on:
        - de-hello-service

  de-hello-service:
    image: abigfish/hello-service
    network_mode: "host"
    environment:
        CONFIG_NAME: "config"
    volumes:
        - "/etc/localtime:/etc/localtime:ro"
        - "/Users/chenliyu/go/src/dapr-examples/hello-service/config:/config"
    restart: always
    ports:
        - 3000:3000
```

docker-compose up -d

