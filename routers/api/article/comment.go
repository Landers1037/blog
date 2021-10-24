/*
Name: blog
File: comment.go
Author: Landers
*/

package article

import (
	"blog/models/dao/comment_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文章评论
// 数据库已文章的name做为主键所以只需传入name
func GetComments(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "parse name failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get comments",
		"data": comment_dao.CommentQuery(name),
	})
}

type commentBody struct {
	Name    string
	User    string
	Comment string
}

func AddComments(c *gin.Context) {
	var b commentBody
	e := c.BindJSON(&b)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "parse body failed",
			"data": "failed",
		})
		return
	}
	e = comment_dao.CommentAdd(b.Name, b.User, b.Comment)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "add comment failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "add comment success",
		"data": "success",
	})
}
