/*
Name: blog
File: create.go
Author: Landers
*/

package comment_dao

import (
	"blog/models"
	"blog/models/article"
	"blog/utils"
	"errors"
)

// 评论新增(只会存在插入情况)
// date直接来自于服务器
func CommentAdd(name, user, comment string) error {
	var com  = article.DB_BLOG_COMMENTS{
		Name: name,
		User: user,
		Date: utils.GetDatePlus(),
		Comment: comment,
	}

	if models.BlogDB.Model(&article.DB_BLOG_COMMENTS{}).Create(&com).Error != nil {
		return errors.New("add comment failed")
	}
	return nil
}
