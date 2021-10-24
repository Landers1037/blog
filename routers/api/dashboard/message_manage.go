/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/message_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageData struct {
	ID int `json:"id"`
}

type messageUpdateData struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func UpdateMessage(c *gin.Context) {
	var m messageUpdateData
	e := c.BindJSON(&m)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "parse body failed",
			"data": "fail",
		})
		return
	}
	e = message_dao.UpdateMessage(m.ID, m.Message)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "update message failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "update message success",
		"data": "success",
	})
}

func DeleteMessage(c *gin.Context) {
	var m messageData
	e := c.BindJSON(&m)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "parse body failed",
			"data": "fail",
		})
		return
	}
	e = message_dao.DelMessage(m.ID)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "del message failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "del message success",
		"data": "success",
	})
}

func GetMessage(c *gin.Context) {
	message_list := message_dao.GetAllMessage()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get message success",
		"data": message_list,
	})
}
