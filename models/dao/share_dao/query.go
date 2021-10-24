/*
Name: blog
File: query.go
Author: Landers
*/

package share_dao

import (
	"blog/models"
	"blog/models/article"
)

func ShareQuery(name string) int {
	var s article.DB_BLOG_SHARE
	e := models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Where("name = ?", name).First(&s).Error
	if e != nil {
		return 0
	} else {
		return s.Share
	}
}
