/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package routers

import (
	"cloudp/middleware"
	"cloudp/routers/api/article"
	"cloudp/routers/api/message"
	"cloudp/routers/api/robotTXT"
	"cloudp/routers/api/sys"
	"cloudp/utils/settings"
	"github.com/gin-gonic/gin"
)

var (
	uv_flag bool
	simpleauth bool
)

func InitRouter() *gin.Engine {
	uv_flag = settings.Uv_flag
	simpleauth = settings.SimpleAuth_flag
	gin.SetMode(settings.RunMode)
	r := gin.New()
	//中间件
	r.Use(gin.Logger())
	r.Use(middleware.Cors())
	//统计类的中间件分路由组编写
	//r.Use(middleware.Uv())
	//r.Use(middleware.SimpleAuth())
	r.Use(gin.Recovery())


	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	//仅供测试

	robot := r.Group("/")
	{
		robot.GET("/robot.txt",robotTXT.GetRobot)
	}
	//文章api
	apiArticle := r.Group("/api/article")
	add(apiArticle,uv_flag,simpleauth)
	{
		apiArticle.GET("/tags", article.Gettags)         //获取全部标签
		apiArticle.GET("/tag", article.Getarticle_bytag) //获取对应标签下的文章
		apiArticle.GET("/posts", article.Getarticles)    //全部文章列表
		apiArticle.GET("/post", article.Getarticle)      //指定文章
		apiArticle.GET("/brother", article.Getbrother)   //指定文章前后篇
		apiArticle.GET("/se",article.Search) //搜索
	}
	//系统参数api
	apiSys := r.Group("/api/sys")
	add(apiSys,false,simpleauth)
	{
		apiSys.GET("/routines", sys.GetRoutines) //routine数目
		apiSys.GET("/mem", sys.GetMem) //占用内存大小
		apiSys.GET("/uv", sys.GetUv) // uv访问数
		apiSys.GET("/count", sys.GetCount) //ip当日访问总数，基于nginx日志
		apiSys.GET("/checkredis",sys.CheckCache) //测试是否命中
	}
	//留言
	apiMes := r.Group("/api/")
	add(apiSys,uv_flag,simpleauth)
	{
		apiMes.GET("/message",message.Getmes)
		apiMes.POST("/message",message.Savemes)
	}
	return r
}

