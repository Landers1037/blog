/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models/dao/post_dao"
	"blog/models/dao/zhuanlan_dao"
	"blog/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 专栏相关接口
// 继承自dashboard 控制面板的有鉴权 本接口为通用请求接口

// 专栏列表
// 返回专栏id，专栏标题 生成的link 专栏描述 文章数组
func GetZhuanLanList(c *gin.Context) {
	list := zhuanlan_dao.ZhuanLanQueryAll()
	var res []response.RES_ZHUANLAN1
	for _, l := range list {
		var link string
		if l.Name == "" {
			link = string(l.PrimaryID)
		}else {
			link = l.Name
		}
		r := response.RES_ZHUANLAN1{
			Link:    link,
			Title:   l.Title,
			Date:    l.Date,
			Posts:   strings.Split(l.Posts, ","),
			Content: l.Content,
		}
		res = append(res, r)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "get zhuanlan list success",
		"data": res,
	})
}


// 返回指定专栏内容
// 包括专栏id 专栏name如果存在 专栏描述 创建时间 文章列表详情
func GetZhuanLan(c *gin.Context) {
	link := c.Param("link")
	var res response.RES_ZHUANLAN_DETAIL
	zhuanlan := zhuanlan_dao.ZhuanLanQuery(link)
	res.ID = zhuanlan.PrimaryID
	res.Title = zhuanlan.Title
	res.Date = zhuanlan.Date
	res.Content = zhuanlan.Content
	// 只返回uri合法的文章内容其他会被过滤掉
	for _, p := range strings.Split(zhuanlan.Posts, ",") {
		postData, e := post_dao.PostQuery(map[string]interface{}{"name": p})
		if e == nil && postData.Name != "" {
			res.Posts = append(res.Posts, response.Zpost{
				Name:     postData.Name,
				Title:    postData.Title,
				Abstract: postData.Abstract,
				Tags:     postData.Tags,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "get zhuanlan detail success",
		"data": res,
	})
}