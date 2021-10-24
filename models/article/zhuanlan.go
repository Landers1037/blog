/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// DB_BLOG_ZHUANLAN 专栏
// 精选多篇文章后组成一个专栏
type DB_BLOG_ZHUANLAN struct {
	models.Model
	Name    string `gorm:"unique" json:"name"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Posts   string `json:"posts"`
	Content string `json:"content"`
}

// posts在数据库中以，分割
