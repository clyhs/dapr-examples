package handler

import (
	"log"
	"dapr-examples/hello-service/common/http"
	dapr "github.com/dapr/go-sdk/service/common"
)

// http 路由注册
func InitHttp(s dapr.Service) (err error) {
	log.Println("http router register")

	register := http.NewRouteRegister(s)
	// 注册路由

	// 设备管理
	user := new(User)
	err = register.GET("user/list", user.List)

	return err
}
