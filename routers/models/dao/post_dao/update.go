/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog/models"
	"blog/models/article"
)

// 改

// 根据字典map来更新需要的值 只有map里定义的值才会被更新
// 首先应该判断这个是否存在
// 值得注意的是 name是唯一的修改重复会报错
// id同样 并且不建议直接修改id
func PostUpdate(con map[string]interface{}, data map[string]interface{}) error {
	if len(data) <= 1 {
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(con).Update(data).Error
		return e
	}else {
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(con).Updates(data).Error
		return e
	}
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

// 联动更新
// 更新标签

// 更新分类