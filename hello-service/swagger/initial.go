package swagger

import (
	dapr "github.com/dapr/go-sdk/service/common"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/swag"
	"dapr-examples/hello-service/common/http"
	"dapr-examples/hello-service/common/response"
)

func InitHttp(s dapr.Service) (err error) {
	log.Println("swagger router register")

	// 注册 swagger 接口地址
	register := http.NewRouteRegister(s)
	err = register.GET("swagger.json", func(r *http.Request) *response.Response {
		jsonStr, err := swag.ReadDoc()
		if err != nil {
			return response.Fail(-1,err.Error())
		}
		return response.OK(jsonStr)
	})
	return err
}
