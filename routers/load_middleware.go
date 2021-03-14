/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package routers

import (
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func LoadMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(middleware.Cors())
	r.Use(middleware.AllowIe())
	//统计类的中间件分路由组编写
	r.Use(middleware.Uv())
	r.Use(middleware.PostView())
	//r.Use(middleware.SimpleAuth())
	r.Use(gin.Recovery())
}
