/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"blog/models"
	"blog/models/message"
	"blog/utils"
)

// 获取最近留言
func GetMessage() []message.DB_BLOG_MESSAGES {
	var list []message.DB_BLOG_MESSAGES
	var orderBy = utils.GetSortMessage()
	models.BlogDB.Model(&message.DB_BLOG_MESSAGES{}).Order(orderBy).Find(&list)
	if len(list)>10 {
		list = list[:10]
	}

	return list
}

// 获取全部留言 设计懒加载方式
func GetAllMessage() []message.DB_BLOG_MESSAGES {
	var list []message.DB_BLOG_MESSAGES
	var orderBy = utils.GetSortMessage()
	models.BlogDB.Model(&message.DB_BLOG_MESSAGES{}).Order(orderBy).Find(&list)

	return list
}

// 懒加载
// page 初始值为0 分页数
func GetMessageOffset(page, perpage int) []message.DB_BLOG_MESSAGES {
	var list []message.DB_BLOG_MESSAGES
	var orderBy = utils.GetSortMessage()
	page = page * perpage
	models.BlogDB.Model(&message.DB_BLOG_MESSAGES{}).Order(orderBy).Offset(page).Limit(perpage).Find(&list)

	return list
}