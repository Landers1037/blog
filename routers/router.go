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
	"blog/routers/api/admin"
	"blog/routers/api/article"
	"blog/routers/api/message"
	"blog/routers/api/robotTXT"
	"blog/routers/api/statistic"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	simpleauth := config.Cfg.SimpleAuth_flag
	gin.SetMode(config.Cfg.RunMode)
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
	addSimpleAuth(apiArticle, simpleauth)
	{
		apiArticle.GET("/tags", article.Gettags)         //获取全部标签
		apiArticle.GET("/tag", article.Getarticle_bytag) //获取对应标签下的文章
		apiArticle.GET("/posts", article.Getarticles)    //全部文章列表
		apiArticle.GET("/post", article.Getarticle)      //指定文章
		apiArticle.GET("/brother", article.Getbrother)   //指定文章前后篇
		apiArticle.GET("/search",article.Search) //搜索
	}
	//系统参数api
	apiSys := r.Group("/api/statistic")
	addSimpleAuth(apiSys, simpleauth)
	{
		apiSys.GET("/routines", statistic.GetRoutines)  //routine数目
		apiSys.GET("/mem", statistic.GetMem)            //占用内存大小
		apiSys.GET("/views", statistic.GetViews)        // uv访问数
		apiSys.GET("/daily", statistic.GetDaily)        //ip当日访问总数，基于nginx日志
		apiSys.GET("/checkredis", statistic.CheckCache) //测试是否命中
	}
	//留言
	apiMes := r.Group("/api")
	addSimpleAuth(apiSys, simpleauth)
	{
		apiMes.GET("/message",message.Getmes)
		apiMes.POST("/message",message.Savemes)
	}
	// 后台登陆
	apiAdmin := r.Group("/api/admin")
	{
		apiAdmin.POST("/login", admin.AdminLogin)
		apiAdmin.PUT("/login", admin.AdminLogin)
		apiAdmin.DELETE("/logout", admin.AdminLogout)
	}
	// 后台控制面板
	apiDashboard := r.Group("/api/dashboard", middleware.AdminAuth())
	{
		apiDashboard.GET("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello this is my blog",
			})
		})
	}

	return r
}

