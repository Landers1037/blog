/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package article

import "blog/models"

// DB_BLOG_TAGS 标签表
type DB_BLOG_TAGS struct {
	models.Model
	Tag string `json:"tag"`
	Name string `json:"name"`
}
