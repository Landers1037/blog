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
	"blog/utils"
	"strings"
)

// 改

// 根据内容的更新 全部更新
// 保证name是存在的情况下才会调用此函数
func PostUpdateAll(postData utils.MdData) error {
	var post = map[string]interface{}{
		"title": postData.Meta.Title,
		"date": postData.Meta.Date,
		"date_plus": postData.Meta.DatePlus,
		"abstract": postData.Abstract,
		"content": postData.Body,
		"tags": strings.Join(postData.Meta.Tags, " "),
		"categories": strings.Join(postData.Meta.Categories, " "),
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", postData.Meta.Name).Updates(post).Error
	logger.BlogLogger.InfoF("开始保存文章 %s 错误: %v", postData.Meta.Name, e)
	if e == nil {
		logger.BlogLogger.InfoF("文章保存完毕 开始更新标签和分类")
		SubTagAdd(postData)
		SubCateAdd(postData)
	}
	return e
}

// 根据字典map来更新需要的值 只有map里定义的值才会被更新
// 首先应该判断这个是否存在
// 值得注意的是 name是唯一的修改重复会报错
// id同样 并且不建议直接修改id
func PostUpdate(name string, data utils.MdData) error {
	// 生成更新字典
	d := map[string]interface{}{
		"title": data.Meta.Title,
		"date": data.Meta.Date,
		"date_plus": data.Meta.DatePlus,
		"tags": strings.Join(data.Meta.Tags, " "),
		"categories": strings.Join(data.Meta.Categories, " "),
		"abstract": data.Abstract,
		"content": data.Body,
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Updates(d).Error
	logger.BlogLogger.InfoF("开始更新文章 %s 错误: %v", name, e)
	if e == nil {
		logger.BlogLogger.InfoF("文章更新完毕 开始更新标签和分类")
		SubTagUpdate(name, data.Meta)
		SubCateUpdate(name, data.Meta)
	}
	return e
}

// 设置pin置顶文章
func PostSetPin(name string) error {
	// 全局只能有一个pin
	var p article.DB_BLOG_POST
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).First(&p).Error
	if e != nil {
		// 不存在pin随便设置
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Update("pin", 1).Error
		return e
	}else {
		// 存在先将原文章改为正常
		p.Pin = 0
		models.BlogDB.Save(&p)
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Update("pin", 1).Error
		if e != nil {
			// 失败回滚
			p.Pin = 1
			models.BlogDB.Save(&p)
		}
		return e
	}
}

// dashboard使用的更新
// 由于支持更换URI 所以区分name newname
// 更新pin时如果为1的已经存在则将其先更新为0
func PostUpdateMap(name, newname, title, date, tags string, pin int) error {
	// 前置处理
	if pin == 1 {
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).Update("pin", 0).Error
		if e != nil {
			logger.BlogLogger.ErrorF("文章重置置顶失败 %s", e.Error())
		}
	}
	var dateSplit string
	if len(strings.Fields(date)) >= 2 {
		dateSplit = strings.Fields(date)[0]
	}else {
		dateSplit = date
	}
	d := map[string]interface{}{
		"name": newname,
		"title": title,
		"date_plus": date,
		"date": dateSplit,
		"tags": tags,
		"update": utils.GetDatePlus(),
		"pin": pin,
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Updates(d).Error
	logger.BlogLogger.InfoF("开始更新文章 %s 错误: %v", name, e)
	if e == nil {
		// 按照逻辑更新成功后 name已经刷新完毕
		logger.BlogLogger.InfoF("文章更新完毕 开始更新标签和分类")
		if name == newname {
			meta := utils.Meta{Name: name, Tags: strings.Fields(tags)}
			SubTagUpdate(name, meta)
			SubCateUpdate(name, meta)
		}else {
			// 删除旧的 创建新的
			SubTagDelOld(name)
			meta := utils.Meta{Name: newname, Tags: strings.Fields(tags)}
			SubTagUpdate(newname, meta)
			SubCateUpdate(newname, meta)
		}
	}
	return e
}

func PostUpdateEditor(name, title, tags, content string) error {
	d := map[string]interface{}{
		"title": title,
		"tags": tags,
		"update": utils.GetDatePlus(),
		"content": content,
	}
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Updates(d).Error
	logger.BlogLogger.InfoF("开始更新文章 %s 错误: %v", name, e)
	if e == nil {
		// 按照逻辑更新成功后 name已经刷新完毕
		logger.BlogLogger.InfoF("文章更新完毕 开始更新标签和分类")
		meta := utils.Meta{Name: name, Tags: strings.Fields(tags)}
		SubTagUpdate(name, meta)
		SubCateUpdate(name, meta)
	}
	return e
}

// dashboard更新文章正文
func PostUpdateContent(name, content string) error {
	d := map[string]interface{}{
		"content": content,
	}
	logger.BlogLogger.InfoF("更新文章%s", name)
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", name).Updates(d).Error
	if e != nil {
		logger.BlogLogger.ErrorF("开始更新文章 %s 错误: %s", name, e.Error())
	}
	return e
}

// 联动更新
// 更新标签
// 因为可能更新name 所以这里单独传入name
func SubTagUpdate(name string, meta utils.Meta) {
	// 查询标签表 不存在则直接添加
	var tmp article.DB_BLOG_TAGS
	e := models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where(map[string]interface{}{"name": name}).First(&tmp).Error
	if e != nil {
		// 不存在记录直接创建
		tags := meta.Tags
		for _, t := range tags {
			if t != "" {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  name,
				})
			}
		}
	}else {
		// 更新不存在的标签
		// 为了保证每次更新的标签不会重复 进行先删除后插入的方式
		tags := meta.Tags
		_ = models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Delete(&article.DB_BLOG_TAGS{}, "name = ?", name).Error
		for _, t := range tags {
			if t != "" {
				models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Create(&article.DB_BLOG_TAGS{
					Tag:   t,
					Name:  name,
				})
			}
		}
	}
}
// 更新分类
func SubCateUpdate(name string, meta utils.Meta) {

}

func SubTagDelOld(name string) {
	e := models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Delete(&article.DB_BLOG_TAGS{}, "name = ?", name).Error
	if e != nil {
		logger.BlogLogger.ErrorF("删除旧文章%s的标签失败: %s", e.Error())
	}
}