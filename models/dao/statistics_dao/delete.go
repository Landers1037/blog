/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package statistics_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
)

// 尽力删除模型 不管存不存在数据都会尝试删除
func StatisticViewDel(name string) {
	e := models.BlogDB.Delete(&article.DB_BLOG_VIEWS{}, "name = ?", name).Error
	if e != nil {
		logger.BlogLogger.ErrorF("数据库删除访问量%s失败", name)
	}
}

func StatisticShareDel(name string) {
	e := models.BlogDB.Delete(&article.DB_BLOG_SHARE{}, "name = ?", name).Error
	if e != nil {
		logger.BlogLogger.ErrorF("数据库删除分享量%s失败", name)
	}
}

func StatisticLikeDel(name string) {
	e := models.BlogDB.Delete(&article.DB_BLOG_LIKES{}, "name = ?", name).Error
	if e != nil {
		logger.BlogLogger.ErrorF("数据库删除点赞量%s失败", name)
	}
}