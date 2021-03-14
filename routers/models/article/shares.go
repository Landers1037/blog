/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

type DB_BLOG_SHARE struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	Share int `json:"share"`
}