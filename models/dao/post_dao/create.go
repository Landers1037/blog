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
	"blog/models/dao/dberr"
	"blog/utils"
	"strings"
)

// 增
// 现阶段没有事务的 所以不保证数据的一致性

// 添加博客文章
// 考虑重复name场景 添加失败
// 联动更改 标签和分类都需要变更
func PostAdd(postData utils.MdData) error {
	var tmp article.DB_BLOG_POST
	var post  = article.DB_BLOG_POST{
		Name: postData.Meta.Name,
		Title: postData.Meta.Title,
		Date: postData.Meta.Date,
		DatePlus: postData.Meta.DatePlus,
		Abstract: postData.Abstract,
		Content: postData.Body,
		Tags: strings.Join(postData.Meta.Tags, " "),
		Categories: strings.Join(postData.Meta.Categories, " "),
		Pin: 0,
	}

	if models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(map[string]interface{}{"name": post.Name}).First(&tmp).Error == nil {
		return dberr.ERR_ARTICLE_EXIST
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Create(&post).Error
	logger.BlogLogger.InfoF("开始保存文章: %s 错误: %v", post.Name, e)
	if e == nil {
		logger.BlogLogger.InfoF("文章保存完毕 开始更新标签和分类")
		SubTagAdd(postData)
		SubCateAdd(postData)
	}
	return e
}

// 标签和分类不作为暴露接口单独使用
// 标签更新失败不会引起博客更新失败 但是会记录日志
func SubTagAdd(postData utils.MdData) {
	// 查询标签表 不存在则直接添加
	var tmp article.DB_BLOG_TAGS
	e := models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where(map[string]interface{}{"name": postData.Meta.Name}).First(&tmp).Error
	if e != nil {
		// 不存在记录直接创建
		tags := postData.Meta.Tags
		for _, t := range tags {
			if t != "" {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  postData.Meta.Name,
				})
			}
		}
	}else {
		// 更新不存在的标签
		// 为了保证每次更新的标签不会重复 进行先删除后插入的方式
		tags := postData.Meta.Tags
		_ = models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Delete(&article.DB_BLOG_TAGS{}, "name = ?", postData.Meta.Name).Error
		for _, t := range tags {
			if t != "" {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  postData.Meta.Name,
				})
			}
		}
	}
}

func SubCateAdd(postData utils.MdData) {

}
