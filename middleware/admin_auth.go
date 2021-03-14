/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// 控制面板的登陆用户校验
// 规则名称和密码定义在app.ini文件中 在此文件内设置你的用户名和密码
// 前端可以选择记住账号
// 所有的密码以明文方式加载
// cookie based

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check cookie
		adminToken, e := c.Request.Cookie("admin_token")
		if e != nil {
			c.AbortWithStatus(403)
			return
		}
		// parse cookie

		// decrypt cookie data
		if utils.AdminCheck(adminToken.Value) {
			c.Next()
		}else {
			c.AbortWithStatus(403)
			return
		}
		// check if valid

		// return
	}
}
