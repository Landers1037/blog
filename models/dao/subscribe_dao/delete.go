/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package subscribe_dao

import (
	"blog/models"
	"blog/models/subscribe"
)

// SubscribeDelete 取消订阅
// 保证一定成功
func SubscribeDelete(mail string) error {
	var sb subscribe.DB_BLOG_SUBSCRIBE
	e := models.BlogDB.Model(&subscribe.DB_BLOG_SUBSCRIBE{}).Where("mail = ?", mail).First(&sb).Error
	if e != nil {
		return nil
	}
	return models.BlogDB.Delete(&sb).Error
}
