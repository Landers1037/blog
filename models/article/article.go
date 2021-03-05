/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article

import "cloudp/models"

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

type Posts struct {
	Model
	Title string `json:"title"`
	Content string `json:"content"`
	Url string `json:"url"`
}
type Post struct {
	Model
	
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Url string `json:"url"`
}

func Getarticles() (articles []Posts)  {
	models.THISdb.Model(&Posts{}).Find(&articles)
	return
}

func Getarticle(name string) (article Post)  {
	models.THISdb.Model(&Post{}).Where("url = ?",name).First(&article)
	return
}

func Getbrother(name string) (p ,n string)  {
	var a Posts
	var prev Posts
	var next Posts
	models.THISdb.Where("url = ?",name).First(&a)
	models.THISdb.First(&prev,a.ID-1)
	p = prev.Url
	models.THISdb.First(&next,a.ID+1)
	n = next.Url

	return p,n
}