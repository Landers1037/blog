/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package routers

import (
	"blog_br_ng/config"
	"blog_br_ng/routers/api/article"
	"blog_br_ng/routers/api/message"
	"blog_br_ng/routers/api/robotTXT"
	"blog_br_ng/routers/api/statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	simpleauth := config.Cfg.SimpleAuth_flag
	gin.SetMode(config.Cfg.RunMode)
	r := gin.New()
	//中间件
	LoadMiddleWare(r)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	//仅供测试

	index := r.Group("/")
	{
		index.GET("robot.txt",robotTXT.GetRobot)
		index.StaticFile("", "dist/index.html")
		index.StaticFile("about", "dist/index.html")
		index.StaticFile("archive", "dist/index.html")
		index.StaticFile("search", "dist/index.html")
		index.StaticFile("overview", "dist/index.html")
		index.StaticFile("message", "dist/index.html")
		index.StaticFile("newui", "dist/index.html")
		index.StaticFile("newui/article", "dist/index.html")
		index.GET("p/:post", func(c *gin.Context) {
			c.File("dist/index.html")
		})
		index.GET("t/:tag", func(c *gin.Context) {
			c.File("dist/index.html")
		})
		index.GET("newui/p:url", func(c *gin.Context) {
			c.File("dist/index.html")
		})
		index.StaticFile("favicon.ico", "dist/favicon.ico")
		index.StaticFile("apple-icon.png", "dist/apple-icon.png")
		index.StaticFile("apple-icon120.png", "dist/apple-icon120.png")
		index.StaticFile("apple-icon152.png", "dist/apple-icon152.png")
		index.StaticFS("js", http.Dir("dist/js"))
		index.StaticFS("css", http.Dir("dist/css"))
		index.StaticFS("fonts", http.Dir("dist/fonts"))
		index.StaticFS("img", http.Dir("./dist/img"))
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
	apiMes := r.Group("/api/")
	addSimpleAuth(apiSys, simpleauth)
	{
		apiMes.GET("/message",message.Getmes)
		apiMes.POST("/message",message.Savemes)
	}
	return r
}

