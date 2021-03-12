/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package sql

import (
	"blog/config"
	"github.com/jinzhu/gorm"
)


func Sqlite()  (db *gorm.DB,err error) {
	var sqliteDB string
	sqliteDB = config.Cfg.DB
	db, err = gorm.Open("sqlite3",sqliteDB)
	return
}

