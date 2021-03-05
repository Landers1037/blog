/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article

import "cloudp/models"

type Tags struct {
	Model
	Tag  string `json:"tag"`
	Posturl string `json:"posturl"`
	Title string `json:"title"`
	Content string `json:"content"`

}

func Gettags() (tags []Tags) {
	//返回全部的tags
	models.THISdb.Table("tags").Select("tag").Where("tag != ?","暂时没有标签").Find(&tags)
	//第一个元素
	//db.First(&Tag{})
	return
}

func Getarticle_bytag(tag string) (posts []Tags)  {
	models.THISdb.Model(&Tags{}).Where("tag = ?",tag).Find(&posts)
	return
}
