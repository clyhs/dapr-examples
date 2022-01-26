package model

import "dapr-examples/hello-service/common/model/request"

type User struct {
	ID          uint   `json:"id" gorm:"primarykey;comment:ID"`
	Name        string `json:"name" gorm:"type:varchar(100);comment:名称" validate:"required"`
}

type UserParam struct {
	request.PageInfo
	request.SortInfo

	ID     string `json:"id"`
	Name   string `json:"name"`
}