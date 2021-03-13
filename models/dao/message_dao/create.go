/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"blog_br_ng/models"
	"blog_br_ng/models/message"
	"blog_br_ng/utils"
	"time"
)

// 添加留言
func SaveMessage(mes string) bool {
	m := message.DB_BLOG_MESSAGES{Message:mes, Date: utils.GetDatePlus()}
	e := models.BlogDB.Create(&m).Error
	time.Sleep(1)
	if e != nil {
		return false
	}
	return true
}