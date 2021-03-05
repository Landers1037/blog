/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package sys

import (
	"cloudp/models/sys"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUv(c *gin.Context)  {
	uv :=  sys.GetuvFromdb()
	c.JSON(http.StatusOK,uv)
}
