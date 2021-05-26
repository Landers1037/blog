/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"blog/utils"
)

// 控制面板的登陆用户校验
// 规则名称和密码定义在app.ini文件中 在此文件内设置你的用户名和密码
// 前端可以选择记住账号
// 所有的密码以明文方式加载
// cookie based

// 在http情况下不能进行cookie同源传递
// 使用token验证
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Cfg.StopAdmin {
			c.Next()
		}else {
			// 任意一种方式校验成功都可以

			// check if valid
			if CheckToken(c) {
				c.Next()
			}else {
				c.AbortWithStatus(403)
				return
			}
			// return
		}
	}
}

func checkCookie(c *gin.Context) bool {
	// check cookie
	adminToken, e := c.Request.Cookie("admin_token")
	if e != nil {
		return false
	}
	// parse cookie

	// decrypt cookie data
	if utils.AdminCheck(adminToken.Value) {
		return true
	}else {
		return false
	}
}

func CheckToken(c *gin.Context) bool {
	token := c.GetHeader("admin_token")
	if token == "" {
		return false
	}
	if utils.AdminCheck(token) {
		return true
	}else {
		return false
	}
}