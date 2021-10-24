/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// DB_BLOG_CATES 分类表
type DB_BLOG_CATES struct {
	models.Model
	Cate string `json:"cate"`
	Name string `json:"name"`
}
