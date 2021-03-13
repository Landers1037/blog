/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog_br_ng/models"
	"blog_br_ng/models/article"
)

// 删

// 删除博客场景 无需考虑之前是否存在
func PostDelete(con map[string]interface{}) error {
	e := models.BlogDB.Where(con).Delete(&article.DB_BLOG_POST{}).Error
	return e
}

// 联动删除
// 删除标签

// 删除分类