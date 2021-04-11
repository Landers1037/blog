/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/statistics_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 操作统计量

// 访问量
type viewData struct {
	Name string `json:"name"`
	Count int `json:"count"`
}
func UpdateView(c *gin.Context) {
	var data viewData
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update view fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticsViewUpdateFull(data.Name, data.Count)
	c.JSON(http.StatusOK, gin.H{
		"msg": "update view success",
		"data": "success",
	})
}

// 直接清空
type viewData2 struct {
	Name string `json:"name"`
}
func DeleteView(c *gin.Context) {
	var data viewData2
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete view fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticViewDel(data.Name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete view success",
		"data": "success",
	})
}

func GetView(c *gin.Context) {
	name := c.Query("name")
	// 为空时返回列表
	if name == "" {
		res := statistics_dao.StatisticViewQueryAll()
		c.JSON(http.StatusOK, gin.H{
			"msg": "get views success",
			"data": res,
		})
		return
	}
	res := statistics_dao.StatisticViewQuery(name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "get view success",
		"data": res,
	})
}

// 操作分享量
// 懒得归一化结构体了
type shareData struct {
	Name string `json:"name"`
	Count int `json:"count"`
}
func UpdateShare(c *gin.Context) {
	var data shareData
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update share fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticShareUpdate(data.Name, data.Count)
	c.JSON(http.StatusOK, gin.H{
		"msg": "update share success",
		"data": "success",
	})
}

type shareData2 struct {
	Name string `json:"name"`
}
func DeleteShare(c *gin.Context) {
	var data shareData2
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete share fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticShareDel(data.Name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete share success",
		"data": "success",
	})
}

func GetShare(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "get share success",
			"data": 0,
		})
		return
	}
	res := statistics_dao.StatisticShareQuery(name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "get share success",
		"data": res,
	})
}

// 操作点赞数
type likeData struct {
	Name string `json:"name"`
	Count int `json:"count"`
}
func UpdateLike(c *gin.Context) {
	var data likeData
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "update like fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticLikeUpdate(data.Name, data.Count)
	c.JSON(http.StatusOK, gin.H{
		"msg": "update like success",
		"data": "success",
	})
}

type likeData2 struct {
	Name string `json:"name"`
}
func DeleteLike(c *gin.Context) {
	var data likeData2
	e := c.BindJSON(&data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete like fail",
			"data": "fail",
		})
		return
	}
	statistics_dao.StatisticLikeDel(data.Name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete like success",
		"data": "success",
	})
}

func GetLike(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "get like success",
			"data": 0,
		})
		return
	}
	res := statistics_dao.StatisticLikeQuery(name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "get like success",
		"data": res,
	})
}
