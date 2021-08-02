/*
Name: blog
File: views.go
Author: Landers
*/

package article

import (
	"blog/models/dao/statistics_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetViews 文章访问量
func GetViews(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK,gin.H{
			"msg": "get views failed",
			"data": 0,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"msg": "get views success",
		"data": statistics_dao.StatisticViewQuery(name),
	})
}
