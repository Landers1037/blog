/*
Name: blog
File: shres.go
Author: Landers
*/

package article

import (
	"blog/models/dao/share_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type shareData struct {
	Name string
}

func GetShares(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get shares failed",
			"data": "failed",
		})
		return
	}
	shareCount := share_dao.ShareQuery(name)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get shares success",
		"data": shareCount,
	})
}

func AddShares(c *gin.Context) {
	var s shareData
	err := c.BindJSON(&s)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "add shares failed",
			"data": "failed",
		})
		return
	}
	err = share_dao.ShareAdd(s.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "add shares failed",
			"data": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "add shares success",
		"data": "success",
	})
	return
}
