/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article

import (
	"cloudp/models/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getbrother(c *gin.Context)  {
	name := c.Query("name")
	p,n := article.Getbrother(name)
	var data = []string{p,n}
	c.JSON(http.StatusOK,data)
}
