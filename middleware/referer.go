/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package middleware
//基于referer的简单验证

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func SimpleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var referer= c.Request.Referer()
		var host = c.Request.Host

		allowRefer := config.Cfg.AppRefer
		allowHost := config.Cfg.AppHost

		ifRefer := check(strings.Fields(allowRefer), referer)
		ifHost := check(strings.Fields(allowHost), host)

		if ifRefer && ifHost{
			// 处理请求
			c.Next()
		}else {
			c.Abort()
			//return
		}

	}
}

func check(strList []string, sub string) bool {
	for _, v := range strList {
		if strings.Contains(sub, v) {
			return true
		}
	}
	return false
}