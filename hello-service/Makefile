SRV=hello-service
VERSION=latest
APP_PORT=3000
DAPR_HTTP_PORT=3501
DAPR_GRPC_PORT=3801
LOG_LEVEL=info
AUTHOR=abigfish
#


.PHONY: build 
build:
	go env -w CGO_ENABLED=0
	go env -w GOOS=linux
	go env -w GOARCH=amd64
	go build -a -ldflags '-extldflags "-static"' -o de-${SRV} main.go
	go env -w CGO_ENABLED=1
	go env -w GOOS=darwin

.PHONY: docker
docker:
	docker build . -t ${AUTHOR}/${SRV}:${VERSION}

.PHONY: swagger
swagger:
	# 生成 swagger 文档
	swag init -o ./swagger

start:
	rm -rf logs
	dapr run --app-id ${SRV} \
         --app-protocol grpc \
         --app-port ${APP_PORT} \
         --dapr-http-port ${DAPR_HTTP_PORT} \
		 --dapr-grpc-port ${DAPR_GRPC_PORT} \
         --log-level ${LOG_LEVEL} \
         --components-path ./config \
         go run main.go