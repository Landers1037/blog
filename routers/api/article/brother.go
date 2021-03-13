/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package article

import (
	"blog_br_ng/models/dao/post_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getbrother(c *gin.Context)  {
	name := c.Query("name")
	p,n := post_dao.Getbrother(name)
	var data = []string{p,n}
	c.JSON(http.StatusOK,data)
}
