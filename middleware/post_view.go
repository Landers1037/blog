/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"github.com/gin-gonic/gin"
)

// 文章开启阅读访问量统计
// 防止频繁写库 按时写入
// 防止并发场景下多次计时 采用每n分钟刷新机制

func PostView() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}