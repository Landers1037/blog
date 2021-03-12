/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

type DB_BLOG_COMMENTS struct {
	models.Model
	Name string `json:"name"`
	Date string `json:"date"`
	Comment int `json:"comment"`
}