/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"strings"
)

// 如果是ie返回错误
// 仅拦截ie8-10
// 如需使用在ini配置中开启此项
func AllowIe() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowIe := config.Cfg.AllowIE
		if !allowIe {
			userAgent := c.Request.UserAgent()
			userAgent = strings.ReplaceAll(userAgent, " ", "")
			if strings.Contains(userAgent, "MSIE8") ||
				strings.Contains(userAgent, "MSIE9") ||
				strings.Contains(userAgent, "MSIE10") {
				c.AbortWithStatus(403)
				return
			}
			c.Next()
		}
	}
}
