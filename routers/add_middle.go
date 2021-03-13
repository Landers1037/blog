/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package routers

import (
	"blog_br_ng/middleware"
	"github.com/gin-gonic/gin"
)


func addSimpleAuth(r *gin.RouterGroup, simple bool)  {
	if simple {
		r.Use(middleware.SimpleAuth())
	}
}
