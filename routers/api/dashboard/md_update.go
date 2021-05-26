/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/post_dao"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// 针对文章的内容更新
// 只更新有值的内容 支持3种方式 type=file type=args type=editor
type postData struct {
	Name string `json:"name"`
	NewName string `json:"newname"`
	Title string `json:"title"`
	Date string `json:"date"`
	Tags string `json:"tags"`
	Pin int `json:"pin"`
}

type postData2 struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Tags string `json:"tags"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	uploadType := c.Query("type")
	name := c.Query("name")
	if uploadType == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "type is empty",
			"data": "fail",
		})
		return
	}
	if uploadType == "file" {
		if name == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "name is empty",
				"data": "fail",
			})
			return
		}
		file, e := c.FormFile("uploadmd")
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse file failed",
				"data": "fail",
			})
			return
		}
		if file.Size >= 1024 * 100 || !strings.Contains(file.Filename, ".md") {
			c.JSON(http.StatusOK, gin.H{
				"msg": "file invalid",
				"data": "fail",
			})
			return
		}
		f, e := file.Open()
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse file failed",
				"data": "fail",
			})
			return
		}
		// 无需解析直接存库
		b, e := ioutil.ReadAll(f)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse file failed",
				"data": "fail",
			})
			return
		}
		e = post_dao.PostUpdateContent(name, string(b))
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse file failed",
				"data": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "update post success",
			"data": "success",
		})

	}else if uploadType == "args" {
		var d postData
		e := c.BindJSON(&d)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse body failed",
				"data": "fail",
			})
			return
		}
		e = post_dao.PostUpdateMap(
			d.Name,
			d.NewName,
			d.Title,
			d.Date,
			d.Tags,
			d.Pin,
			)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "update post failed",
				"data": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "update post success",
			"data": "success",
		})
	}else if uploadType == "editor" {
		var d postData2
		e := c.BindJSON(&d)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "parse body failed",
				"data": "fail",
			})
			return
		}
		e = post_dao.PostUpdateEditor(
			d.Name,
			d.Title,
			d.Tags,
			d.Content,
		)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "update post failed",
				"data": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "update post success",
			"data": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "type is invalid",
			"data": "fail",
		})
		return
	}
}
