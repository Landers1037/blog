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
	"blog/routers/api/admin"
	"blog/routers/api/article"
	"blog/routers/api/dashboard"
	"blog/routers/api/message"
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

	//文章api
	apiArticle := r.Group("/api/article")
	addSimpleAuth(apiArticle)
	{
		apiArticle.GET("/tags", article.GetTags)             // 获取全部标签
		apiArticle.GET("/tag", article.Getarticle_bytag)     // 获取对应标签下的文章
		apiArticle.GET("/posts", article.Getarticles)        // 全部文章列表
		apiArticle.GET("/post", article.Getarticle)          // 指定文章
		apiArticle.GET("/brother", article.GetBrother)       // 指定文章前后篇
		apiArticle.GET("/search", article.Search)            // 搜索
		apiArticle.GET("/comments", article.GetComments)     // 获取评论
		apiArticle.POST("/comments", article.AddComments)    // 发布评论
		apiArticle.GET("/likes", article.GetLikes)           // 获取点赞
		apiArticle.POST("/likes", article.AddLikes)          // 点赞
		apiArticle.GET("/shares", article.GetShares)         // 获取分享
		apiArticle.POST("/shares", article.AddShares)        // 分享
		apiArticle.GET("/views", article.GetViews)           // 获取访问量
		apiArticle.GET("/archive", article.GetArchive)       // 获取归档
		apiArticle.GET("/archives", article.GetArchivePosts) // 获取归档文章列表
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
		apiMes.GET("/message", message.Getmes)   // 获取留言
		apiMes.POST("/message", message.Savemes) // 保存留言
	}
	// 邮件订阅

	// 专栏文章
	apiZhuanlan := r.Group("/api/zhuanlan")
	{
		apiZhuanlan.GET("", article.GetZhuanLanList)   // 全部专栏列表
		apiZhuanlan.GET("/:link", article.GetZhuanLan) // 专栏信息
	}
	// 后台登陆
	apiAdmin := r.Group("/api/admin")
	{
		apiAdmin.POST("/login", admin.AdminLogin)     // 用户登陆
		apiAdmin.PUT("/login", admin.AdminLogin)      // 用户登陆put方式
		apiAdmin.DELETE("/logout", admin.AdminLogout) // 用户登出
	}
	// 后台控制面板
	// 带有权限控制 需要验证cookie
	apiDashboard := r.Group("/api/dashboard")
	addAdminAuth(apiDashboard)
	{
		// 新增类rest接口
		apiDashboard.POST("/post/upload", dashboard.UploadFile)        // 文章上传
		apiDashboard.POST("/post/parse", dashboard.UploadFileCallBack) // 文章上传前的解析回调
		apiDashboard.POST("/post/check", dashboard.CheckFile)          // 文章格式校验
		apiDashboard.POST("/post/export", dashboard.ExportPosts)       // 文章数据导出
		apiDashboard.POST("/post/add", dashboard.NewPost)              // 新增空的文章
		apiDashboard.POST("/db/init", dashboard.InitDB)                // 数据库初始化
		apiDashboard.POST("/db/backup", dashboard.BackUpDB)            // 数据库备份
		apiDashboard.POST("/db/export", dashboard.ExportDataBase)      // 数据库导出

		// 更新类rest接口
		apiDashboard.POST("/post", dashboard.UpdatePost)        // 接受文件上传
		apiDashboard.PUT("/post", dashboard.UpdatePost)         // 文章更新
		apiDashboard.PUT("/tag", dashboard.UpdateTag)           // 文章标签更新
		apiDashboard.PUT("/category", dashboard.UpdateCate)     // 分类更新
		apiDashboard.PUT("/comment", dashboard.UpdateComment)   // 留言更新
		apiDashboard.PUT("/view", dashboard.UpdateView)         // 访问量更新
		apiDashboard.PUT("/share", dashboard.UpdateShare)       // 分享量更新
		apiDashboard.PUT("/like", dashboard.UpdateLike)         // 点赞更新
		apiDashboard.PUT("/message", dashboard.UpdateMessage)   // 留言更新
		apiDashboard.PUT("/subscribe", dashboard.UpdateSub)     // 订阅更新
		apiDashboard.PUT("/zhuanlan", dashboard.UpdateZhuanlan) // 专栏更新

		// 删除类rest接口
		apiDashboard.DELETE("/post", dashboard.DeletePost)
		apiDashboard.DELETE("/tag", dashboard.DeleteTag)
		apiDashboard.DELETE("/category", dashboard.DeleteCate)
		apiDashboard.DELETE("/comment", dashboard.DeleteComment)
		apiDashboard.DELETE("/view", dashboard.DeleteView)
		apiDashboard.DELETE("/share", dashboard.DeleteShare)
		apiDashboard.DELETE("/like", dashboard.DeleteLike)
		apiDashboard.DELETE("/message", dashboard.DeleteMessage)
		apiDashboard.DELETE("/subscribe", dashboard.DeleteSub)
		apiDashboard.DELETE("/zhuanlan", dashboard.DeleteZhuanlan)

		// 获取类rest接口
		apiDashboard.GET("/post", dashboard.GetPost)
		apiDashboard.GET("/tag", dashboard.GetTag)
		apiDashboard.GET("/category", dashboard.GetCate)
		apiDashboard.GET("/comment", dashboard.GetComment)
		apiDashboard.GET("/view", dashboard.GetView)
		apiDashboard.GET("/share", dashboard.GetShare)
		apiDashboard.GET("/like", dashboard.GetLike)
		apiDashboard.GET("/message", dashboard.GetMessage)
		apiDashboard.GET("/subscribe", dashboard.GetSub)
		apiDashboard.GET("/zhuanlan", dashboard.GetZhuanlan)
	}

	return r
}
