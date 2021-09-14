/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package crdb

import (
	"blog/config"
	"blog/logger"
	"blog/models"
)

// InitTables 启动时初始化所有数据库表
// 默认创建
func InitTables() error {
	for table, tableStruct := range Tables {
		if !config.Cfg.HideDBLog {
			logWithDBConfig("info", "初始化表%s", table)
		}

		if models.BlogDB.HasTable(tableStruct) {
			logWithDBConfig("info", "已存在，跳过")
		}else {
			e := models.BlogDB.CreateTable(tableStruct).Error
			if e != nil {
				logWithDBConfig("error", "创建表失败%s", e.Error())
				return e
			}
			logWithDBConfig("info", "表创建完毕")
		}
	}
	return nil
}

func logWithDBConfig(t, s string, args ...interface{}) {
	if !config.Cfg.HideDBLog {
		switch t {
		case "error":
			logger.BlogLogger.ErrorF(s, args...)
		case "info":
			logger.BlogLogger.InfoF(s, args...)
		default:
			logger.BlogLogger.InfoF(s, args...)
		}
	}
}