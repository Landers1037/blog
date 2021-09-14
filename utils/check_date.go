/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"time"
)

// CheckUpdateTime 检查上传日期和md内部的创建日期
// 当上传日期大于创建日期时才使用此日期作为更新日期
func CheckUpdateTime(createTime string) (updateTime string, flag bool) {
	var nowTime string
	nowTime = GetDatePlus()
	now, e1 := time.Parse("2006-01-02 15:04:05", nowTime)
	create, e2 := time.Parse("2006-01-02 15:04:05", createTime)
	if e1 != nil || e2 != nil {
		return nowTime, true
	}
	if now.After(create) {
		return nowTime, true
	}
	return createTime, false
}
