/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package sql

import (
	"cloudp/utils/settings"
	"github.com/jinzhu/gorm"
)


func Sqlite()  (db *gorm.DB,err error) {
	var sq string
	sec,_ := settings.Cfg.GetSection("sqlite")
	sq = sec.Key("DB").String()

	db,_= gorm.Open("sqlite3",sq)
	return db,nil
}

