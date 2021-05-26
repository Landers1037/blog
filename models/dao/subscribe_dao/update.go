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

// 更新访问量
func StatisticsViewUpdate(name string, count int) {
	// 不存在则先创建
	var v article.DB_BLOG_VIEWS
	e := models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Where("name = ?", name).First(&v).Error
	if e == nil {
		v.View = v.View + count
		models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Where("name = ?", name).Update("view", v.View)
	}else {
		models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Create(&article.DB_BLOG_VIEWS{
			Name:  name,
			View:  1,
		})
	}
}