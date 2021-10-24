/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package article

import (
	"blog/models"
)

// Search 搜索
// 模糊搜索 当前支持title name content abstract tag cate
func Search(name string) (s []DB_BLOG_POST) {
	var pattern = "%" + name + "%"
	models.BlogDB.Where("title LIKE ?", pattern, "name LIKE ?", pattern, "abstract LIKE ?", pattern,
		"content LIKE ?", pattern, "tag LIKE ?", pattern, "categories LIKE ?", pattern).Find(&s)

	return
}
