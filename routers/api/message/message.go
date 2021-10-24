/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package message

import (
	"blog/models/dao/message_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	Mes string `json:"mes"`
}

func Savemes(c *gin.Context) {
	var m message
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(http.StatusForbidden, "error")
	} else {
		flag := message_dao.SaveMessage(m.Mes)
		if !flag {
			c.JSON(http.StatusOK, "error")
		} else {
			c.JSON(http.StatusOK, "saved")
		}

	}
}

func Getmes(c *gin.Context) {
	data := message_dao.GetMessage()
	c.JSON(http.StatusOK, data)
}
