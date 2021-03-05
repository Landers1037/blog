/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package middleware
//基于referer的简单验证

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func SimpleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var referer string = c.Request.Referer()
		var host string = c.Request.Host
		//sec ,_:= settings.Cfg.GetSection("nginx")
		//port = sec.Key("PORT").MustString("1325")
		//注意这里是nginx转发的所以不知道真实的host
		ifBlog := strings.Index(referer,"renj.io") != -1
		ifGithub := strings.Index(referer,"github.com") != -1
		ifHost := strings.Index(host,"127.0.0.1") != -1
		if ifBlog || ifGithub && ifHost{
			// 处理请求
			c.Next()
		}else {
			c.Abort()
			//return
		}

	}
}

