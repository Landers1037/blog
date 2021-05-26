/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package zhuanlan_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
	"blog/models/dao/dberr"
	"blog/utils"
	"strings"
)

// 不使用创建接口
// 创建和更新统一使用本函数
func ZhuanLanUpdate(id int, name, title, date, posts, content string) error {
	// name为空也没事所以直接存储
	// 对于posts为空的情况不会存储
	// 因为可能多篇文章属于不同专栏 所以不做name限制

	// 约定id为0时为创建
	if id == 0 {
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
			logger.BlogLogger.ErrorF("尝试创建专栏失败: %s", e.Error())
			return e
		}
		return nil
	}
	var z article.DB_BLOG_ZHUANLAN
	e := models.BlogDB.Model(article.DB_BLOG_ZHUANLAN{}).Where("primary_id = ?", id).First(&z).Error
	if e != nil {
		// 专栏不存在
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
			logger.BlogLogger.ErrorF("尝试创建专栏失败: %s", e.Error())
			return e
		}
	}else {
		// 更新专栏内容
		if name != "" {
			z.Name = name
		}
		if title != "" {
			z.Title = title
		}
		if date != "" {
			z.Date = date
		}
		if strings.TrimSpace(posts) != "" && posts != "" {
			z.Posts = posts
		}
		if content != "" {
			z.Content = content
		}
		e = models.BlogDB.Save(&z).Error
		if e != nil {
			logger.BlogLogger.ErrorF("尝试更新专栏失败: %s", e.Error())
			return e
		}
	}

	return nil
}
