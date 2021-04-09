/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package statistics_dao

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

// 更新分享量 尽力保证更新成功所以不返回错误
// 按照逻辑 不存在时只会创建初始值为1的数据
func StatisticShareUpdate(name string, count int) {
	var s article.DB_BLOG_SHARE
	e := models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Where("name = ?", name).First(&s).Error
	if e == nil {
		s.Share = count
		models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Where("name = ?", name).Update("share", s.Share)
	}else {
		models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Create(&article.DB_BLOG_SHARE{
			Name:  name,
			Share: 1,
		})
	}
}

// 更新点赞
func StatisticLikeUpdate(name string, count int) {
	var l article.DB_BLOG_LIKES
	e := models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Where("name = ?", name).First(&l).Error
	if e == nil {
		l.Like = count
		models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Where("name = ?", name).Update("like", l.Like)
	}else {
		models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Create(&article.DB_BLOG_LIKES{
			Name:  name,
			Like:  1,
		})
	}
}