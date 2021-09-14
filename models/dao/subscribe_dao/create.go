/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package subscribe_dao

import (
	"blog/models"
	"blog/models/subscribe"
	"blog/utils"
)

// SubscribeAdd 添加逻辑 判断是否是邮箱由前端校验
func SubscribeAdd(mail, period string) error {
	var sb subscribe.DB_BLOG_SUBSCRIBE
	e := models.BlogDB.Model(&subscribe.DB_BLOG_SUBSCRIBE{}).Where("mail = ?", mail).First(&sb).Error
	if e != nil {
		return e
	}
	switch period {
	case "week":
		period = "week"
	case "month":
		period = "month"
	default:
		period = "week"
	}
	return models.BlogDB.Model(&subscribe.DB_BLOG_SUBSCRIBE{}).Create(&subscribe.DB_BLOG_SUBSCRIBE{
		Mail:          mail,
		SubscribeDate: utils.GetDatePlus(),
		Period:        period,
	}).Error
}