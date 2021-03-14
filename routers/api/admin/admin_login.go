/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package admin

import (
	"blog/config"
	"blog/utils"
	"blog/utils/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginData struct{
	Name string	`json:"name"`
	Passwd string `json:"passwd"`
}

// 使用用户名和密码登录
// 用户名和密码将会和配置文件内的进行比对
func AdminLogin(c *gin.Context) {
	var login LoginData
	e := c.BindJSON(&login)
	if e != nil {
		code := err.ERROR
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"msg": err.GetMsg(code),
			"data": "failed",
		})
	}

	if config.Cfg.AdminName == login.Name &&
		config.Cfg.AdminPwd == login.Passwd {
		c.SetCookie("admin_token", utils.AdminEncrypt(), config.Cfg.CookieMaxAge,
			"/", config.Cfg.AppDomain, false, false)
		c.SetCookie("login_time", utils.GetDateTime(), config.Cfg.CookieMaxAge,
			"/", config.Cfg.AppDomain, false, false)
		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"msg": err.GetMsg(code),
			"data": "success",
		})
	}else {
		code := err.ERROR
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"msg": err.GetMsg(code),
			"data": "failed",
		})
	}
}
