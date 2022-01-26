package handler

import (
	"dapr-examples/hello-service/common/http"
	"dapr-examples/hello-service/common/response"
	"dapr-examples/hello-service/model"
	"dapr-examples/hello-service/service"
	resp "dapr-examples/hello-service/common/model/response"
)

type User struct {
}

func (c *User) List(r *http.Request) *response.Response  {
	var searchParams *model.UserParam
	if err := r.Parse(&searchParams); err != nil {
		return response.Fail(-1, err.Error())
	}
	err, list, total:= service.GetUserList(searchParams)
	if err != nil {
		return response.Fail(-1, err.Error())
	}
	return response.OK(resp.PageResult{
		List:     list,
		Total:    total,
		Page:     searchParams.Page,
		PageSize: searchParams.PageSize,
	})
}