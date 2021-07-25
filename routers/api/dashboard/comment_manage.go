/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/comment_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 博客评论留言管理
type commentUpdate struct {
	Name string
	ID int
	Comment string
}

func UpdateComment(c *gin.Context) {
	var com commentUpdate
	err := c.BindJSON(&com)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update comment failed",
			"data": "failed",
		})
		return
	}
	err = comment_dao.CommentUpdate(com.Name, com.ID, com.Comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update comment failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update comment success",
		"data": "success",
	})
	return
}

type commentDel struct {
	Name string
	ID int
}

func DeleteComment(c *gin.Context) {
	var com commentDel
	err := c.BindJSON(&com)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete comment failed",
			"data": "failed",
		})
		return
	}
	err = comment_dao.CommentDelete(com.Name, com.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete comment failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete comment success",
		"data": "success",
	})
	return
}

func GetComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get comments success",
		"data": comment_dao.CommentQueryAll(),
	})
}