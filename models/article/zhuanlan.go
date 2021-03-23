/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// 专栏
// 精选多篇文章后组成一个专栏
type DB_BLOG_ZHUANLAN struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts string `json:"posts"`
	Content string `json:"content"`
}
