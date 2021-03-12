/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog/models"
	"blog/models/article"
	"errors"
	"strings"
)

// 增
// 现阶段没有事务的 所以不保证数据的一致性

// 添加博客文章
// 考虑重复name场景 添加失败
// 联动更改 标签和分类都需要变更
func PostAdd(post article.DB_BLOG_POST) error {
	var tmp article.DB_BLOG_POST
	if models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(map[string]interface{}{"name": post.Name}).First(&tmp).Error != nil {
		return errors.New("article already exists")
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Create(&post).Error
	return e
}

// 标签和分类不作为暴露接口单独使用
// 标签更新失败不会引起博客更新失败 但是会记录日志
func tagAdd(post article.DB_BLOG_POST) {
	// 查询标签表 不存在则直接添加
	var tmp article.DB_BLOG_TAGS
	e := models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where(map[string]interface{}{"name": post.Name}).First(&tmp).Error
	if e != nil {
		tags := strings.Fields(post.Tags)
		for _, t := range tags {
			if t != "" {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  post.Name,
				})
			}
		}
	}else {
		// 更新不存在的标签
		tags := strings.Fields(post.Tags)
		var tmp article.DB_BLOG_TAGS
		for _, t := range tags {
			if t != "" && models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where(map[string]interface{}{"name": post.Name, "tag": t}).First(&tmp).Error != nil {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  post.Name,
				})
			}
		}
	}
}

func cateAdd(post article.DB_BLOG_POST) {

}
