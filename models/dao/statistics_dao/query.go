/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package statistics_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
	"blog/models/response"
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

// 用于响应后台 需要附带post的title

func StatisticViewQueryAll() []response.ViewPosts {
	var res []response.ViewPosts
	var v []article.DB_BLOG_VIEWS
	e := models.BlogDB.Model(&article.DB_BLOG_VIEWS{}).Find(&v).Error
	if e != nil {
		logger.BlogLogger.ErrorF("获取全部访问量统计失败: %s", e.Error())
	}
	var tmpres []response.ViewPosts
	for _, view := range v {
		// 对all单独判断
		if view.Name != "all" {
			var p article.DB_BLOG_POST
			e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", view.Name).Find(&p).Error
			if e == nil {
				tmpres = append(tmpres, response.ViewPosts{
					Name:  view.Name,
					View:  view.View,
					Title: p.Title,
				})
			}
		}else if view.Name == "all" {
			res = append(res, response.ViewPosts{
				Name: view.Name,
				View: view.View,
				Title: "全局统计量",
			})
		}
		// 非法文章不记录
	}
	res = append(res, tmpres...)
	return res
}

// 获取当日访问量 依赖于临时内存
func StatisticDailyQuery() int {
	return 0
}

// 获取分享量
func StatisticShareQuery(name, dtype string) interface{} {
	if name == "all" {
		var v []article.DB_BLOG_SHARE
		var count int
		e := models.BlogDB.Model(article.DB_BLOG_SHARE{}).Find(&v).Error
		if e != nil {
			if dtype == "data" {
				return v
			}
			return 0
		}
		if dtype == "data" {
			return v
		}
		for _, s := range v {
			count += s.Share
		}
		return count
	}else {
		var v article.DB_BLOG_SHARE
		e := models.BlogDB.Model(article.DB_BLOG_SHARE{}).Where("name = ?", name).First(&v).Error
		if e != nil {
			if dtype == "data" {
				return v
			}
			return 0
		}
		if dtype == "data" {
			return v
		}
		return v.Share
	}
}

// 获取点赞数
func StatisticLikeQuery(name, dtype string) interface{} {
	if name == "all" {
		var v []article.DB_BLOG_LIKES
		var count int
		e := models.BlogDB.Model(article.DB_BLOG_LIKES{}).Find(&v).Error
		if e != nil {
			if dtype == "data" {
				return v
			}
			return 0
		}
		if dtype == "data" {
			return v
		}
		for _, l := range v {
			count += l.Like
		}
		return count
	}else {
		var v article.DB_BLOG_LIKES
		e := models.BlogDB.Model(article.DB_BLOG_LIKES{}).Where("name = ?", name).First(&v).Error
		if e != nil {
			if dtype == "data" {
				return v
			}
			return 0
		}
		if dtype == "data" {
			return v
		}
		return v.Like
	}
}

// 获取评论数
func StatisticCommentQuery(name string) int {
	var v int
	if name == "all" {
		e := models.BlogDB.Model(article.DB_BLOG_COMMENTS{}).Count(&v).Error
		if e != nil {
			return 0
		}
		return v
	}else {
		e := models.BlogDB.Model(article.DB_BLOG_COMMENTS{}).Where("name = ?", name).Count(&v).Error
		if e != nil {
			return 0
		}
		return v
	}
}