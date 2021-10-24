/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package zhuanlan_dao

import (
	"blog/models"
	"blog/models/article"
	"strconv"
)

// ZhuanLanQueryAll 专栏列表获取
func ZhuanLanQueryAll() []article.DB_BLOG_ZHUANLAN {
	var zhuanlanList []article.DB_BLOG_ZHUANLAN
	e := models.BlogDB.Model(article.DB_BLOG_ZHUANLAN{}).Find(&zhuanlanList).Error
	if e != nil {
		return zhuanlanList
	}
	return zhuanlanList
}

// ZhuanLanQuery 专栏信息
// 联动获取文章详情
func ZhuanLanQuery(link string) article.DB_BLOG_ZHUANLAN {
	//  因为主键不重复 name不重复 link可能有两种形式
	id, e := strconv.Atoi(link)
	var z article.DB_BLOG_ZHUANLAN
	if e == nil {
		// 传入为id
		models.BlogDB.Model(article.DB_BLOG_ZHUANLAN{}).Where("primary_id = ?", id).First(&z)
	} else {
		// 传入为name
		models.BlogDB.Model(article.DB_BLOG_ZHUANLAN{}).Where("name = ?", link).First(&z)
	}
	return z
}
