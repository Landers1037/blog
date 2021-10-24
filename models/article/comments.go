/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// DB_BLOG_COMMENTS 评论表
type DB_BLOG_COMMENTS struct {
	models.Model
	Name    string `json:"name"`
	User    string `json:"user"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
}
