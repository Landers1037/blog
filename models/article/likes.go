/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog_br_ng/models"
)


type DB_BLOG_LIKES struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	Like int `json:"like"`
}