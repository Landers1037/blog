/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package sys

import "cloudp/models"

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}
type Uv struct {
	Model

	Count int `json:"count"`
	Name string `json:"name"`

}

func GetuvFromdb() (uv Uv) {
	models.THISdb.First(&uv,1)
	return
}

func AddUv(num int) {
	var uv Uv
	//每5个请求为一个缓存
	models.THISdb.Select("count").First(&uv)
	num = num+uv.Count
	models.THISdb.First(&uv,1).Update("count",num)
	return
}