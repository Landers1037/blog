/*
Name: blog
File: create.go
Author: Landers
*/

package like_dao

import (
	"blog/models"
	"blog/models/article"
	"sync"
)

var likeLock sync.Mutex

// LikeAdd 新增点赞
func LikeAdd(name string) error {
	var l article.DB_BLOG_LIKES
	e := models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Where("name = ?", name).First(&l).Error
	if e != nil {
		l.Like = 1
		l.Name = name
		return models.BlogDB.Model(&article.DB_BLOG_LIKES{}).Create(&l).Error
	}else {
		likeLock.Lock()
		defer likeLock.Unlock()
		l.Like = l.Like + 1
		return models.BlogDB.Save(&l).Error
	}
}
