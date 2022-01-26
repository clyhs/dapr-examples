package service

import (
	"dapr-examples/hello-service/model"
	"gorm.io/gorm"
)

func parseUserFilter(db *gorm.DB, searchParams *model.UserParam) *gorm.DB {
	// 名称
	if searchParams.Name != "" {
		db = db.Where("name = ?", searchParams.Name)
	}

	return db
}

func GetUserList(searchParams *model.UserParam) (err error, list []model.User, total int64) {
	// 创建db
	db := model.DB.Model(&model.User{})
	// 条件过滤
	db = parseUserFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		orderBy := "id asc"
		if searchParams.OrderBy != "" {
			db.Order(searchParams.OrderBy)
		}
		db.Order(orderBy)
		// 分页
		if searchParams.PageSize > 0 {
			limit := searchParams.PageSize
			offset := searchParams.PageSize * (searchParams.Page - 1)
			db.Limit(limit).Offset(offset)
		}
		err = db.Find(&list).Error
	}
	return err, list, total
}
