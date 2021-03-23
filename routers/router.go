/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package routers

import (
	"blog/config"
	"blog/logger"
	"blog/middleware"
	"blog/routers/api/admin"
	"blog/routers/api/article"
	"blog/routers/api/dashboard"
	"blog/routers/api/message"
	"blog/routers/api/robotTXT"
	"blog/routers/api/statistic"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Cfg.RunMode)
	logger.BlogLogger.InfoF("当前服务运行模式: %s", config.Cfg.RunMode)
	r := gin.New()
	//中间件
	LoadMiddleWare(r)
	addStaticRouter(r)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello this is my blog",
		})
	})
	//仅供测试

	robot := r.Group("/")
	{
		robot.GET("/robot.txt",robotTXT.GetRobot)
	}
	//文章api
	apiArticle := r.Group("/api/article")
	addSimpleAuth(apiArticle)
	{
		apiArticle.GET("/tags", article.Gettags)         // 获取全部标签
		apiArticle.GET("/tag", article.Getarticle_bytag) // 获取对应标签下的文章
		apiArticle.GET("/posts", article.Getarticles)    // 全部文章列表
		apiArticle.GET("/post", article.Getarticle)      // 指定文章
		apiArticle.GET("/brother", article.Getbrother)   // 指定文章前后篇
		apiArticle.GET("/search",article.Search)         // 搜索
	}
	//系统参数api
	apiSys := r.Group("/api/statistic")
	addSimpleAuth(apiSys)
	{
		apiSys.GET("/routines", statistic.GetRoutines)  // routine数目
		apiSys.GET("/mem", statistic.GetMem)            // 占用内存大小
		apiSys.GET("/views", statistic.GetViews)        // uv访问数
		apiSys.GET("/daily", statistic.GetDaily)        // ip当日访问总数，基于nginx日志
		apiSys.GET("/checkredis", statistic.CheckCache) // 测试是否命中
	}
	//留言
	apiMes := r.Group("/api")
	addSimpleAuth(apiSys)
	{
		apiMes.GET("/message",message.Getmes)   // 获取留言
		apiMes.POST("/message",message.Savemes) // 保存留言
	}
	// 邮件订阅

	// 专栏文章
	apiZhuanlan := r.Group("/zhuanlan")
	{
		apiZhuanlan.GET("")
	}
	// 后台登陆
	apiAdmin := r.Group("/api/admin")
	{
		apiAdmin.POST("/login", admin.AdminLogin) // 用户登陆
		apiAdmin.PUT("/login", admin.AdminLogin)  // 用户登陆put方式
		apiAdmin.DELETE("/logout", admin.AdminLogout) // 用户登出
	}
	// 后台控制面板
	// 带有权限控制 需要验证cookie
	apiDashboard := r.Group("/api/dashboard", middleware.AdminAuth())
	addAdminAuth(apiDashboard)
	{
		// 新增类rest接口
		apiDashboard.POST("/post/upload", dashboard.UploadFile)
		apiDashboard.POST("/post/check", dashboard.CheckFile)
		apiDashboard.POST("/post/export", dashboard.CheckFile)
		apiDashboard.POST("/db/backup", dashboard.CheckFile)
		apiDashboard.POST("/db/export", dashboard.CheckFile)

		// 更新类rest接口
		apiDashboard.PUT("/post", dashboard.CheckFile)
		apiDashboard.PUT("/tag", dashboard.CheckFile)
		apiDashboard.PUT("/category", dashboard.CheckFile)
		apiDashboard.PUT("/comment", dashboard.CheckFile)
		apiDashboard.PUT("/view", dashboard.CheckFile)
		apiDashboard.PUT("/share", dashboard.CheckFile)
		apiDashboard.PUT("/message", dashboard.CheckFile)
		apiDashboard.PUT("/subscribe", dashboard.CheckFile)
		apiDashboard.PUT("/zhuanlan", dashboard.CheckFile)

		// 删除类rest接口
		apiDashboard.DELETE("/post", dashboard.DeletePost)
		apiDashboard.DELETE("/tag", dashboard.CheckFile)
		apiDashboard.DELETE("/category", dashboard.CheckFile)
		apiDashboard.DELETE("/comment", dashboard.CheckFile)
		apiDashboard.DELETE("/view", dashboard.CheckFile)
		apiDashboard.DELETE("/share", dashboard.CheckFile)
		apiDashboard.DELETE("/message", dashboard.CheckFile)
		apiDashboard.DELETE("/subscribe", dashboard.CheckFile)
		apiDashboard.DELETE("/zhuanlan", dashboard.CheckFile)

		// 获取类rest接口
		apiDashboard.GET("/post", dashboard.CheckFile)
		apiDashboard.GET("/tag", dashboard.CheckFile)
		apiDashboard.GET("/category", dashboard.CheckFile)
		apiDashboard.GET("/comment", dashboard.CheckFile)
		apiDashboard.GET("/view", dashboard.CheckFile)
		apiDashboard.GET("/share", dashboard.CheckFile)
		apiDashboard.GET("/message", dashboard.CheckFile)
		apiDashboard.GET("/subscribe", dashboard.CheckFile)
		apiDashboard.GET("/zhuanlan", dashboard.CheckFile)
	}

	return r
}

