/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package sys

import (
	"cloudp/utils/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCount(c *gin.Context)  {
	count := cmd.Sh("ip")
	c.JSON(http.StatusOK,count)
}
