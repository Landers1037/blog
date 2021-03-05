/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package robotTXT

import (
	"github.com/gin-gonic/gin"
)

func GetRobot(c *gin.Context)  {

	c.File("conf/robot.txt")
}
