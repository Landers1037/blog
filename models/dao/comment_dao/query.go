/*
Name: blog
File: query.go
Author: Landers
*/

package comment_dao

import (
	"blog/config"
	"blog/models"
	"blog/models/article"
)

// CommentQuery 获取评论
// 根据配置支持不同方式获取
func CommentQuery(name string) []article.DB_BLOG_COMMENTS {
	var coms []article.DB_BLOG_COMMENTS
	models.BlogDB.Model(&article.DB_BLOG_COMMENTS{}).Where("name = ?", name).Order(getSortComment()).Find(&coms)
	return coms
}

func CommentQueryAll() []article.DB_BLOG_COMMENTS {
	var coms []article.DB_BLOG_COMMENTS
	models.BlogDB.Model(&article.DB_BLOG_COMMENTS{}).Order(getSortComment()).Find(&coms)
	return coms
}

func getSortComment() string {
	var s string
	switch config.Cfg.SortCommentBy {
	case "id":
		s = "primary_id"
	case "date":
		s = "date"
	case "name":
		s = "name"
	default:
		s = "primary_id"
	}
	if config.Cfg.SortCommentReverse {
		return s + " " + "desc"
	}
	return s
}
