/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package article

import (
	"blog/models/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	str := c.Query("key")
	if str == "" {
		c.JSON(http.StatusOK, "")
	} else {
		list := article.Search(str)
		c.JSON(http.StatusOK, list)
	}
}
