/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package admin

import (
	"blog/models"
)

// DB_BLOG_ADMIN 管理员表 暂时使用明文方式存储密码
type DB_BLOG_ADMIN struct {
	models.Model
	UserName string `gorm:"unique;not null" json:"user_name"`
	PassWd string `json:"passwd"`
	Date string `json:"date"`
}

// todo 多用户时进行用户分权