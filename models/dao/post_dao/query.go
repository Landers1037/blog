/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog/models"
	"blog/models/article"
	"blog/models/response"
)

// 查
// 博客查询字典聚合单个
// 为防止指针逃逸 全部以返回值的形式响应
func PostQuery(con map[string]interface{}) (article.DB_BLOG_POST, error) {
	var post article.DB_BLOG_POST
	if len(con) > 0 {
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(con).First(&post).Error
		return post, e
	}
	// 没有条件时默认查找第一个
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).First(&post).Error
	return post, e
}

// 符合博客排序规范 按照id排序
// 因为pin文章id为0 且pin=1 所以对结果进行拼接
// pin规则重新设计 不存在pin置顶文章时 默认按照id排序先后
// 存在pin文章时 结果为pin + 剩下文章按照id排序 保证pin文章随时可以修改 只需要修改pin为1 全局只有一个pin为1的文章
func PostQueryAll(con map[string]interface{}) ([]response.RES_POST, error) {
	var posts []response.RES_POST
	if len(con) > 0 {
		var pin article.DB_BLOG_POST
		var normal []article.DB_BLOG_POST
		e1 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).First(&pin).Error
		e2 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(con).Not("pin = ?", 1).Order("id desc").Find(&normal).Error
		if e1 != nil {
			for _, v := range normal {
				posts = append(posts, response.RES_POST{
					Name:     v.Name,
					Title:    v.Title,
					Date:     v.Date,
					Abstract: v.Abstract,
				})
			}
		}else {
			posts = append(posts, response.RES_POST{
				Name:     pin.Name,
				Title:    pin.Title,
				Date:     pin.Date,
				Abstract: pin.Abstract,
			})
			for _, v := range normal {
				posts = append(posts, response.RES_POST{
					Name:     v.Name,
					Title:    v.Title,
					Date:     v.Date,
					Abstract: v.Abstract,
				})
			}
		}

		return posts, e2
	}

	var pin article.DB_BLOG_POST
	var normal []article.DB_BLOG_POST

	e1 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).First(&pin).Error
	e2 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Order("id desc").Not("pin = ?" ,1).Find(&normal).Error
	// 不存在pin使用默认顺序
	if e1 != nil {
		for _, v := range normal {
			posts = append(posts, response.RES_POST{
				Name:     v.Name,
				Title:    v.Title,
				Date:     v.Date,
				Abstract: v.Abstract,
			})
		}
	}else {
		posts = append(posts, response.RES_POST{
			Name:     pin.Name,
			Title:    pin.Title,
			Date:     pin.Date,
			Abstract: pin.Abstract,
		})
		for _, v := range normal {
			posts = append(posts, response.RES_POST{
				Name:     v.Name,
				Title:    v.Title,
				Date:     v.Date,
				Abstract: v.Abstract,
			})
		}
	}

	return posts, e2
}

// todo 暂时不做limit 按照切片返回来做 后期加上

// 获取上下篇
// id越小日期越靠前
// 因为最新文章在最前 所以上一篇是id+1 下一篇是id-1
// 预定义的规则 最后一篇的pre一定是pin文章 pin文章的是它本身
func Getbrother(name string) (p ,n string)  {
	var count int
	var a article.DB_BLOG_POST
	var prev article.DB_BLOG_POST
	var next article.DB_BLOG_POST
	models.BlogDB.Where("name = ?",name).First(&a)
	_ = models.BlogDB.Model(&article.DB_BLOG_POST{}).Count(&count)
	if a.ID == count-1 || a.ID == 0 {
		// 最后一篇文章
		p = "pin"
	}else {
		models.BlogDB.First(&prev, "id = ?", a.ID+1)
		p = prev.Name
	}

	if a.ID == 0 {
		// 置顶的下一篇为最新文章
		models.BlogDB.First(&next, "id = ?", count-1)
		n = next.Name
	} else {
		models.BlogDB.First(&next, "id = ?", a.ID-1)
		n = next.Name
	}

	return p,n
}

// 获取博客标签
func TagQuery(name string) ([]article.DB_BLOG_TAGS, error) {
	var tags []article.DB_BLOG_TAGS
	e := models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where(map[string]interface{}{"name": name}).Find(&tags).Error
	if e != nil {
		return tags, e
	}
	return tags, nil
}

// 获取博客分类
func CateQuery(name string) ([]article.DB_BLOG_CATES, error) {
	var cates []article.DB_BLOG_CATES
	e := models.BlogDB.Model(&article.DB_BLOG_CATES{}).Where(map[string]interface{}{"name": name}).Find(&cates).Error
	if e != nil {
		return cates, e
	}
	return cates, nil
}