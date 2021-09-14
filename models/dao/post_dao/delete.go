/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
)

// 删

// PostDelete 删除博客场景 无需考虑之前是否存在
// 因为name是唯一的所以用name作为约束
func PostDelete(name string) error {
	e := models.BlogDB.Where("name = ?", name).Delete(&article.DB_BLOG_POST{}).Error
	logger.BlogLogger.InfoF("开始删除文章: %s 错误: %v", name, e)
	if e == nil {
		logger.BlogLogger.InfoF("文章删除完毕 开始删除标签和分类")
		SubTagsDel(name)
		SubCateDel(name)
	}
	return e
}

// SubTagsDel 联动删除
// 删除标签
func SubTagsDel(name string) {
	_ = models.BlogDB.Where("name = ?", name).Delete(&article.DB_BLOG_TAGS{}).Error
}

// SubCateDel 删除分类
func SubCateDel(name string) {
	_ = models.BlogDB.Where("name = ?", name).Delete(&article.DB_BLOG_CATES{}).Error
}