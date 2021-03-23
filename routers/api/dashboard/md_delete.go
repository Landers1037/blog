/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/post_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文章的删除
func DeletePost(c *gin.Context) {
	type data struct{
		Name string `json:"name"`
	}
	var d data
	e := c.BindJSON(&d)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "parse body failed",
			"data": "fail",
		})
		return
	}
	e = post_dao.PostDelete(d.Name)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete post failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete post success",
		"data": "success",
	})
	return
}