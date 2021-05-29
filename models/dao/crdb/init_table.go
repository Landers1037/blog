/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package crdb

import (
	"blog/logger"
	"blog/models"
)

// 启动时初始化所有数据库表
// 默认创建
func InitTables() error {
	for table, tableStruct := range Tables {
		logger.BlogLogger.InfoF("初始化表%s", table)
		if models.BlogDB.HasTable(tableStruct) {
			logger.BlogLogger.InfoF("已存在，跳过")
		}else {
			e := models.BlogDB.CreateTable(tableStruct).Error
			if e != nil {
				logger.BlogLogger.ErrorF("创建表失败%s", e.Error())
				return e
			}
			logger.BlogLogger.InfoF("表创建完毕")
		}
	}
	return nil
}
