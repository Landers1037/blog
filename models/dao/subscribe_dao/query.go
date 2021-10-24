/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package subscribe_dao

import (
	"blog/models"
	"blog/models/article"
)

// StatisticViewQuery 获取访问量
func StatisticViewQuery(name string) int {
	var v article.DB_BLOG_VIEWS
	e := models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Where("name = ?", name).First(&v).Error
	if e != nil {
		return 0
	}
	return v.View
}

// StatisticDailyQuery 获取当日访问量 依赖于临时内存
func StatisticDailyQuery() int {
	return 0
}
