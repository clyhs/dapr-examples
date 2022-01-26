package dapr

import (
	"dapr-examples/hello-service/common/config"
	"os"
	"log"
	"dapr-examples/hello-service/handler"
	"dapr-examples/hello-service/model"
	// "dapr-examples/hello-service/service"
	"net/http"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"dapr-examples/hello-service/swagger"
)

func Start() {
	// serviceName := config.String("server.name")
	serviceAddress := os.Getenv("SERVICE_ADDRESS")
	if len(serviceAddress) < 2 {
		serviceAddress = config.String("server.address")
	}

	/*
	logcfg.Config(
		logcfg.AddHook(&logcfg.LogHook{
			Service:   serviceName,
			IP:        "127.0.0.1",
			IsConsole: true,
		}),
	)*/

	s, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	// 注册 swagger 日志接口
	if err := swagger.InitHttp(s); err != nil {
		log.Fatalf("swagger register failed: %v", err)
	}

	// 注册 http 路由
	if err := handler.InitHttp(s); err != nil {
		log.Fatalf("handler register failed: %v", err)
	}
	// 注册 DB
	model.InitDB()

	// 初始化车厅对象
	//service.InitSupportDevice()

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}