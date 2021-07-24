/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package article
////文章列表api
import (
	"blog/middleware"
	"blog/models/dao/tag_dao"
	"blog/utils/err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Getarticle(c *gin.Context){
	//单个文章的获取，因为使用前端渲染的方式所以数据库里只保存文章的文件名称
	//文件名称唯一且等于本地的文件名
	name := c.Query("name")
	if name != ""{
		data := middleware.Cache(name)
		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"msg": err.GetMsg(code),
			"data": data,
		})
	}else {
		c.JSON(http.StatusOK,"")
	}
	//文章的获取应该是传输md文件流

}

func Getarticles(c *gin.Context)  {
	//获取文章的列表
	//data := article.Getarticles()
	page := c.Query("p")
	if page != ""{
		p ,_:= strconv.Atoi(page)
		data ,length:= middleware.PostCache(p)

		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg": err.GetMsg(code),
			"data": data,
			"len": length,

		})
	}else {
		code := err.SUCCESS
		data, _ := middleware.PostCache(0)
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg": err.GetMsg(code),
			"data": data,

		})
	}

}

func Getarticle_bytag(c *gin.Context)  {
	//获取对应tag的文章
	tag := c.Query("tag")
	data := tag_dao.QueryTagWithPosts(tag)
	code := err.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg": err.GetMsg(code),
		"data": data,

	})

}
