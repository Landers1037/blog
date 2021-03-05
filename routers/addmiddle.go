/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package routers

import (
	"cloudp/middleware"
	"github.com/gin-gonic/gin"
)


func add(r *gin.RouterGroup,uv,simple bool)  {
	if uv{
		r.Use(middleware.Uv())
	}
	if simple {
		r.Use(middleware.SimpleAuth())
	}
}
