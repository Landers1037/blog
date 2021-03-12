/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package statistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所有统计访问次数
func GetViews(c *gin.Context)  {
	uv :=  1
	c.JSON(http.StatusOK,uv)
}
