/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package zhuanlan_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
)

func ZhuanLanDel(id int) {
	e := models.BlogDB.Delete(&article.DB_BLOG_ZHUANLAN{}, "primary_id = ?", id).Error
	if e != nil {
		logger.BlogLogger.ErrorF("数据库删除专栏%s失败: %s", id, e.Error())
	}
}
