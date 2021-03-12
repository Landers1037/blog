/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package middleware

import "github.com/gin-gonic/gin"

//  防止内存碎片化 使用按时写入机制
func Uv() gin.HandlerFunc {
	return func(c *gin.Context) {
		//设计一个缓存的栈
		// 处理请求
		c.Next()
	}
}

