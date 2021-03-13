/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package statistic

import (
	"blog_br_ng/utils/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDaily(c *gin.Context)  {
	count := cmd.Sh("ip")
	c.JSON(http.StatusOK,count)
}
