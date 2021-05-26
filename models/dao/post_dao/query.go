/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package post_dao

import (
	"blog/config"
	"blog/models"
	"blog/models/article"
	"blog/models/response"
	"blog/utils"
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
//
// 默认以id方式逆序排序 通过配置文件可以自定义
func PostQueryAll(con map[string]interface{}) ([]response.RES_POST, error) {
	var posts []response.RES_POST
	var orderBy = utils.GetSortPost(config.Cfg.SortPostBy, config.Cfg.SortPostReverse)
	if len(con) > 0 {
		var pin article.DB_BLOG_POST
		var normal []article.DB_BLOG_POST
		e1 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).First(&pin).Error
		e2 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where(con).Not("pin = ?", 1).Order(orderBy).Find(&normal).Error
		if e1 != nil {
			for _, v := range normal {
				// 对返回的摘要进行重置 摘要不存在默认使用自定义空摘要 否则使用截取
				posts = append(posts, response.RES_POST{
					ID:       v.PrimaryID,
					Name:     v.Name,
					Title:    v.Title,
					Date:     v.Date,
					Abstract: utils.Conetent2Abs(v.Abstract, v.Content),
				})
			}
		}else {
			posts = append(posts, response.RES_POST{
				ID:       pin.PrimaryID,
				Name:     pin.Name,
				Title:    pin.Title,
				Date:     pin.Date,
				Abstract: utils.Conetent2Abs(pin.Abstract, pin.Content),
			})
			for _, v := range normal {
				posts = append(posts, response.RES_POST{
					ID:       v.PrimaryID,
					Name:     v.Name,
					Title:    v.Title,
					Date:     v.Date,
					Abstract: utils.Conetent2Abs(v.Abstract, v.Content),
				})
			}
		}

		return posts, e2
	}

	var pin article.DB_BLOG_POST
	var normal []article.DB_BLOG_POST

	e1 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Where("pin = ?", 1).First(&pin).Error
	e2 := models.BlogDB.Model(&article.DB_BLOG_POST{}).Order(orderBy).Not("pin = ?" ,1).Find(&normal).Error
	// 不存在pin使用默认顺序
	if e1 != nil {
		for _, v := range normal {
			posts = append(posts, response.RES_POST{
				ID:       v.PrimaryID,
				Name:     v.Name,
				Title:    v.Title,
				Date:     v.Date,
				Abstract: utils.Conetent2Abs(v.Abstract, v.Content),
			})
		}
	}else {
		posts = append(posts, response.RES_POST{
			ID:       pin.PrimaryID,
			Name:     pin.Name,
			Title:    pin.Title,
			Date:     pin.Date,
			Abstract: utils.Conetent2Abs(pin.Abstract, pin.Content),
		})
		for _, v := range normal {
			posts = append(posts, response.RES_POST{
				ID: 	  v.PrimaryID,
				Name:     v.Name,
				Title:    v.Title,
				Date:     v.Date,
				Abstract: utils.Conetent2Abs(v.Abstract, v.Content),
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
//
// 根据自定义规则当按照其他方式排序时 此时的上下篇不再以id作为媒介
// 当前的憨批办法是全查
func Getbrother(name string) (p ,n string)  {
	var count int
	var current article.DB_BLOG_POST
	var allPost []response.RES_POST_BROTHER

	models.BlogDB.Where("name = ?", name).First(&current)
	_ = models.BlogDB.Model(&article.DB_BLOG_POST{}).Count(&count)
	// 统一的情况 根据文章数组来确定
	// 使用高级查询 避免查询占用过多内存
	orderBy := utils.GetSortPost(config.Cfg.SortPostBy, config.Cfg.SortPostReverse)
	var tmp1 article.DB_BLOG_POST
	var tmp2 []article.DB_BLOG_POST
	e := models.BlogDB.Select("name, title, pin").Where("pin = ?", 1).First(&tmp1).Error
	models.BlogDB.Select("name, title, pin").Not("pin = ?", 1).Order(orderBy).Find(&tmp2)
	if e != nil {
		for _, p := range tmp2 {
			allPost = append(allPost, response.RES_POST_BROTHER{
				Name:  p.Name,
				Title: p.Title,
				Pin:   p.Pin,
			})
		}
	}else {
		allPost = append(allPost, response.RES_POST_BROTHER{
			Name:  tmp1.Name,
			Title: tmp1.Title,
			Pin:   tmp1.Pin,
		})
		for _, p := range tmp2 {
			allPost = append(allPost, response.RES_POST_BROTHER{
				Name:  p.Name,
				Title: p.Title,
				Pin:   p.Pin,
			})
		}
	}

	// count大小应该和长度保持一致
	for index, post := range allPost {
		if post.Name == current.Name {
			// 特殊情况1 当前文章为第一篇
			if index == 0 && index+1 < count {
				return current.Name, allPost[index+1].Name
			}else if index == 0 && index+1 >= count {
				return current.Name, current.Name
			}else if index == count-1 && index-1 >=0 {
				// 情况2 当前文章为最后一篇
				return allPost[index-1].Name, current.Name
			}else if index == count-1 && index-1 <0 {
				return current.Name, current.Name
			}else {
				return allPost[index-1].Name, allPost[index+1].Name
			}
		}
	}
	// 找不到此文章 此情况其实不存在
	return "", ""
}


func Getbrother2(name string) (p ,n string)  {
	var count int
	var a article.DB_BLOG_POST
	var prev article.DB_BLOG_POST
	var next article.DB_BLOG_POST
	models.BlogDB.Where("name = ?",name).First(&a)
	_ = models.BlogDB.Model(&article.DB_BLOG_POST{}).Count(&count)
	if a.ID == count-1 || a.Pin == 1 {
		// 最后一篇文章
		e := models.BlogDB.Model(&article.DB_BLOG_POST{}).First(&prev, "pin = ?", 1).Error
		if e != nil {
			//  不存在pin时返回第一篇
		}
		p = prev.Name
	}else {
		models.BlogDB.First(&prev, "id = ?", a.ID+1)
		p = prev.Name
	}

	if a.Pin == 1 {
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