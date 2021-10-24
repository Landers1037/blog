/*
Name: blog
File: query.go
Author: Landers
*/

package like_dao

import (
	"blog/models"
	"blog/models/article"
)

func LikeQuery(name string) int {
	var l article.DB_BLOG_LIKES
	e := models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Where("name = ?", name).First(&l).Error
	if e != nil {
		return 0
	} else {
		return l.Like
	}
}
