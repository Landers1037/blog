/*
Name: blog
File: blog_log.go
Author: Landers
*/

package middleware

import (
	"blog/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// BlogLogger 自定义的日志器
// 默认同运行日志共享文件流
func BlogLogger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return formatter(params)
		},
		Output:    logger.BlogLogger.Writer,
		SkipPaths: nil,
	})
}

func formatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	// 强制开启颜色
	statusColor = param.StatusCodeColor()
	methodColor = param.MethodColor()
	resetColor = param.ResetColor()
	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	name := fmt.Sprintf("[\x1b[0;%dm%s\x1b[0m]", 31, "BLOG")
	return fmt.Sprintf("%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		name,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
