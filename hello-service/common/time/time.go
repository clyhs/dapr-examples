package time

import (
	"time"
)

/**
 * @description: 获取当前时间格式化
 * @param {*}
 * @return {*}
 */
func GetNowTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

