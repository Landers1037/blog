/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package message

import (
	message2 "cloudp/models/message"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	Mes string `json:"mes"`
}

func Savemes(c *gin.Context)  {
	var m message
	err := c.BindJSON(&m)
	if err != nil{
		c.JSON(http.StatusForbidden,"error")
	}else {
		flag := message2.SaveMs(m.Mes)
		if flag{
			c.JSON(http.StatusOK,"error")
		}else {
			fmt.Println("ok")
			c.JSON(http.StatusOK,"saved")
		}

	}
}

func Getmes(c *gin.Context) {
	data := message2.GetMes()
	c.JSON(http.StatusOK,data)
}