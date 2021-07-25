/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/message"
)

func UpdateMessage(id int, m string) error {
	e := models.BlogDB.Model(&message.DB_BLOG_MESSAGES{}).Where("primary_id = ?", id).Update("message", m).Error
	if e != nil {
		logger.BlogLogger.ErrorF("更新留言%d失败", id)
	}
	return e
}