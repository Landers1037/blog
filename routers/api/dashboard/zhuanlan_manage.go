/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/zhuanlan_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 更新或者新增一个专栏
// 因为name并不是必定存在所以使用id
type zhuanData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts string `json:"posts"` // 以，分割的字符串
	Content string `json:"content"`
}
func UpdateZhuanlan(c *gin.Context) {
	var z zhuanData
	e := c.BindJSON(&z)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "parse body fail",
			"data": "fail",
		})
		return
	}
	e = zhuanlan_dao.ZhuanLanUpdate(
		z.Id, z.Name, z.Title, z.Date, z.Posts, z.Content)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update zhuanlan fail",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update zhuanlan success",
		"data": "success",
	})
	return
}

// 删除专栏
type zhuanData2 struct {
	Id int `json:"id"`
}
func DeleteZhuanlan(c *gin.Context) {
	var z zhuanData2
	e := c.BindJSON(&z)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete zhuanlan fail",
			"data": "fail",
		})
		return
	}
	zhuanlan_dao.ZhuanLanDel(z.Id)
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete zhuanlan success",
		"data": "success",
	})
}

// 获取指定专栏或者全部专栏信息
// 无参数传递时返回所有专栏信息
func GetZhuanlan(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		data := zhuanlan_dao.ZhuanLanQueryAll()
		c.JSON(http.StatusOK, gin.H{
			"msg": "get zhuanlan success",
			"data": data,
		})
		return
	}
	// 为保证数据统一性 在这里全部使用id作为参数
	data := zhuanlan_dao.ZhuanLanQuery(id)
	c.JSON(http.StatusOK, gin.H{
		"msg": "get zhuanlan success",
		"data": data,
	})
}