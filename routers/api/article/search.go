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

func Search(c *gin.Context)  {
	str := c.Query("title")
	if str == ""{
		c.JSON(http.StatusOK,"")
	}else {
		list := article.Search(str)
		c.JSON(http.StatusOK,list)
	}
}
