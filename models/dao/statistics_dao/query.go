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

// 获取访问量
func StatisticViewQuery(name string) int {
	var v article.DB_BLOG_VIEWS
	e := models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Where("name = ?", name).First(&v).Error
	if e != nil {
		return 0
	}
	return v.View
}

// 获取当日访问量 依赖于临时内存
func StatisticDailyQuery() int {
	return 0
}

// 获取分享量
func StatisticShareQuery(name string) int {
	var v article.DB_BLOG_SHARE
	e := models.BlogDB.Model(article.DB_BLOG_SHARE{}).Where("name = ?", name).First(&v).Error
	if e != nil {
		return 0
	}
	return v.Share
}

// 获取点赞数
func StatisticLikeQuery(name string) int {
	var v article.DB_BLOG_LIKES
	e := models.BlogDB.Model(article.DB_BLOG_LIKES{}).Where("name = ?", name).First(&v).Error
	if e != nil {
		return 0
	}
	return v.Like
}