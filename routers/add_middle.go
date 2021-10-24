/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package routers

import (
	"blog/config"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

// 添加简易refer  host验证
func addSimpleAuth(r *gin.RouterGroup) {
	if config.Cfg.SimpleAuthFlag {
		r.Use(middleware.SimpleAuth())
	}
}

// 添加cookie验证
func addAdminAuth(r *gin.RouterGroup) {
	r.Use(middleware.AdminAuth())
}
