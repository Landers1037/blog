/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// DB_BLOG_SHARE 分享量表
type DB_BLOG_SHARE struct {
	models.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Share int    `json:"share"`
}
