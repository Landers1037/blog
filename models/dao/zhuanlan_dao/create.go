/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package zhuanlan_dao

import (
	"blog/models"
	"blog/models/article"
	"blog/models/dao/dberr"
	"blog/utils"
	"strings"
)

// ZhuanLanAdd 保存专栏
func ZhuanLanAdd(name, title, posts, content string) error {
	// name为空也没事所以直接存储
	// 对于posts为空的情况不会存储
	// 因为可能多篇文章属于不同专栏 所以不做name限制
	if strings.TrimSpace(posts) == "" || posts == "" {
		return dberr.ERR_EMPTY_POSTS
	}
	z := article.DB_BLOG_ZHUANLAN{
		Name:    name,
		Title:   title,
		Date:    utils.GetDatePlus(),
		Posts:   posts,
		Content: content,
	}
	e := models.BlogDB.Model(article.DB_BLOG_ZHUANLAN{}).Create(&z).Error
	if e != nil {
		return e
	}
	return nil
}
