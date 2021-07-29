/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/
package models
//数据库的初始化
import (
	"blog/logger"
	"blog/models/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var BlogDB *gorm.DB

func InitDB() {
	var err error
	// 使用sqlite
	BlogDB, err = sql.Sqlite()
	if err != nil {
		logger.BlogLogger.ErrorF("数据库连接失败 %s", err.Error())
		logger.BlogLogger.PanicF("数据库错误 异常退出")
	}

	BlogDB.SingularTable(true)
	BlogDB.DB().SetMaxIdleConns(10)
	BlogDB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	logger.BlogLogger.InfoF("数据库连接已关闭")
	if BlogDB != nil {
		defer BlogDB.Close()
	}
}
