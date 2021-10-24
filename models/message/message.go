/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package message

import (
	"blog/models"
)

// DB_BLOG_MESSAGES 留言表
type DB_BLOG_MESSAGES struct {
	models.Model
	User    string `json:"user"`
	Date    string `json:"date"`
	Message string `gorm:"not null" json:"message"`
}
