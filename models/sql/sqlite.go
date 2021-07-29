/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package sql

import (
	"blog/config"
	"blog/logger"
	"github.com/jinzhu/gorm"
	"os"
	"path"
)

// Sqlite sqlite的连接池
// 对于不存在的数据库或数据库目录会自动创建 仅支持一级目录
func Sqlite() (db *gorm.DB,err error) {
	var sqliteDB string
	sqliteDB = config.Cfg.DB
	err = createDBRoot(sqliteDB)
	if err != nil {
		logger.BlogLogger.PanicF("初始化数据库目录失败%s", err.Error())
		return nil, err
	}
	db, err = gorm.Open("sqlite3", sqliteDB)
	return
}

func createDBRoot(d string) error {
	root := path.Dir(d)
	if _, err := os.Stat(root);os.IsNotExist(err) {
		logger.BlogLogger.InfoF("初始化数据库目录%s", root)
		return os.Mkdir(root, 0644)
	}
	return nil
}