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

func GetRoutines(c * gin.Context)  {
	n := cmd.GetGID()

	c.JSON(http.StatusOK,n)
}

func GetMem(c *gin.Context)  {
	o := cmd.Sh("mem")

	c.JSON(http.StatusOK,o)
}


