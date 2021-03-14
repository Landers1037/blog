/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/
package models
//数据库的初始化
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"blog/models/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var BlogDB *gorm.DB

func InitDB() {
	var err error
	// 使用sqlite
	BlogDB, err = sql.Sqlite()
	if err != nil {
		log.Println(err)
	}

	BlogDB.SingularTable(true)
	BlogDB.DB().SetMaxIdleConns(10)
	BlogDB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer BlogDB.Close()
}
