/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package middleware

import (
	"cloudp/models/sys"
)
import "github.com/gin-gonic/gin"

var tmp  = 0
func Uv() gin.HandlerFunc {
	return func(c *gin.Context) {
		//设计一个缓存的栈
		if tmp>5{
			tmp = 0
			sys.AddUv(5)
		}else {
			tmp++
		}
		// 处理请求
		c.Next()
	}
}

