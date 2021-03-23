/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 模拟try file机制
// 对于使用前端处理路由的分离式设计程序使用
// 默认api的路由为本服务的 其他路由不由本服务处理
func TryFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 对于/进行屏蔽
		if c.Request.URL.Path == "/" {
			c.Status(http.StatusOK)
		}
		if config.Cfg.TryFile {
			path := c.Request.URL.Path
			if strings.Contains(path, "/api") {
				c.Next()
			}else {
				c.Next()
				if c.Writer.Status() == http.StatusNotFound {
					c.Status(http.StatusFound)
					c.File(config.Cfg.TryFileIndex)
				}
			}
		}
	}
}
