/*
Name: blog
File: delete.go
Author: Landers
*/

package comment_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
)

func CommentDelete(name string, id int) error {
	e := models.BlogDB.Where(map[string]interface{}{"name": name, "primary_id": id}).Delete(&article.DB_BLOG_COMMENTS{}).Error
	if e != nil {
		logger.BlogLogger.InfoF("开始删除评论: %d 错误: %v", id, e)
		return e
	}

	return nil
}