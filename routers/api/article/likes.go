/*
Name: blog
File: likes.go
Author: Landers
*/

package article

import (
	"blog/models/dao/like_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type likesData struct {
	Name string
}

// 点赞文章
func GetLikes(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK,gin.H{
			"msg": "get likes failed",
			"data": "failed",
		})
		return
	}
	likeCount := like_dao.LikeQuery(name)
	c.JSON(http.StatusOK,gin.H{
		"msg": "get likes success",
		"data": likeCount,
	})
}

// 点赞会有一个数据库的临时锁
func AddLikes(c *gin.Context) {
	var l likesData
	err := c.BindJSON(&l)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg": "add likes failed",
			"data": "failed",
		})
		return
	}
	err = like_dao.LikeAdd(l.Name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg": "add likes failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"msg": "add likes success",
		"data": "success",
	})
	return
}
