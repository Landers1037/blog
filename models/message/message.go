/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package message

import (
	"cloudp/models"
	"time"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}
type message struct {
	Model
	Mes string `json:"mes"`
}

func GetMes() (meslist []string) {
	var list []message
	models.THISdb.Model(&message{}).Order("id desc").Find(&list)
	if len(list)>10{
		tmp := list[:10]
		for i:=0;i<len(tmp);i++{
			meslist = append(meslist, tmp[i].Mes)
		}
	}else {
		tmp := list
		for i:=0;i<len(tmp);i++{
			meslist = append(meslist, tmp[i].Mes)
		}
	}

	return
}

func SaveMs(mes string)  bool{
	m := message{Mes:mes}
	models.THISdb.Create(&m)
	time.Sleep(1)
	e := models.THISdb.NewRecord(m)
	return e

}