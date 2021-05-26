/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package subscribe

import (
	"blog/models"
)

// 订阅
// period周期 month week
type DB_BLOG_SUBSCRIBE struct {
	models.Model
	Mail string `json:"mail"`
	SubscribeDate string `json:"subscribe_date"`
	Period string `json:"period"`
}
