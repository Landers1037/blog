/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"blog/models"
	"blog/models/message"
)

func DelMessage(id int) error {
	// 尽力删除模型
	e := models.BlogDB.Delete(&message.DB_BLOG_MESSAGES{}, "primary_id = ?", id).Error
	return e
}