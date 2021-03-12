/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package tag_dao

import (
	"blog/models"
	"blog/models/article"
	"blog/models/response"
)

// 标签的dao

// 获取当前存在的全部标签 注意去重
func TagQueryAll() []response.RES_TAG {
	var tags []article.DB_BLOG_TAGS
	var res []response.RES_TAG
	models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where("tag != ?", "暂时没有标签").Find(&tags)
	for _, t := range tags {
		if len(res) == 0 {
			res = append(res, response.RES_TAG{Tag: t.Tag})
		}else {
			for i, r := range res {
				if r.Tag == t.Tag {
					break
				}
				if i == len(res)-1 {
					res = append(res, response.RES_TAG{
						Tag: t.Tag,
					})
				}
			}
		}
	}
	return res
}

// 获取对应标签的全部文章
func QueryTagWithPosts(tag string) ([]response.RES_POST)  {
	var res []response.RES_POST
	var tags []article.DB_BLOG_TAGS
	models.BlogDB.Model(&article.DB_BLOG_TAGS{}).Where("tag = ?",tag).Find(&tags)

	// 在文章表中按照tag中的post name查询
	for _, t := range tags {
		var p article.DB_BLOG_POST
		models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("name = ?", t.Name).First(&p)
		res = append(res, response.RES_POST{
			Name:     p.Name,
			Title:    p.Title,
			Date:     p.Date,
			Abstract: p.Abstract,
		})
	}

	return res
}