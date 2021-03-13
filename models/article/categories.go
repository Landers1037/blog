/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog_br_ng/models"
)
type DB_BLOG_CATES struct {
	models.Model
	Cate string `json:"cate"`
	Name string `json:"name"`
}