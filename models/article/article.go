/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package article

import "blog_br_ng/models"

type DB_BLOG_POST struct {
	models.Model
	ID int `gorm:"unique;not null" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	DatePlus string `json:"date_plus"`
	Update string `json:"update"`
	Abstract string `json:"abstract"`
	Content string `json:"content"`
	Tags string `json:"tags"`
	Categories string `json:"categories"`
	Pin int `json:"pin"`
}