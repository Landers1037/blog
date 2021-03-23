/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package logger

import (
	"blog/utils/settings"
	"fmt"
	"io"
	"os"
)

// 默认只提供到终端的日志记录器
// 如果需要使用其他方式 请使用multiwriter
type Logger struct {
	// os.stdout
	writer io.Writer
	// error || info
	level  string
}

var BlogLogger Logger

// 因为配置文件在日志器后面初始化所以不使用全局的配置
func InitLogger() Logger {
	writers := []io.Writer{
		os.Stdout,
	}
	cf, loadErr := settings.InitCfg()
	if loadErr != nil {
		panic(fmt.Sprintf("加载配置文件失败 %s", loadErr.Error()))
	}
	if cf.APPLogEnable {
		if cf.APPLogFile != "" {
			logFile, e := os.OpenFile(cf.APPLogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
			if e == nil {
				writers = append(writers, logFile)
			}
		}
	}

	BlogLogger = Logger{
		writer: io.MultiWriter(writers...),
		level:  cf.APPLogLevel,
	}

	return BlogLogger
}

// 错误记录
// 所有日志默认都带换行
func (l *Logger) ErrorF(s string, args ...interface{}) {
	if l.level == "error" {
		_, _ = fmt.Fprintf(l.writer, fmtPrefix("error", s), args...)
	}
}

// 提示信息
func (l *Logger) InfoF(s string, args ...interface{}) {
	_, _ = fmt.Fprintf(l.writer, fmtPrefix("info", s), args...)
}

// 异常退出
// 会直接导致程序退出
// 谨慎使用除非遇到必须退出的任务
func (l *Logger) PanicF(s string, args ...interface{}) {
	_, _ = fmt.Fprintf(l.writer, fmtPrefix("panic", s), args...)
	os.Exit(1)
}

// 前缀格式化
// [level] xxxx
func fmtPrefix(level string, log string) string {
	switch level {
	case "error":
		return fmt.Sprintf(Fmt, ERROR, "ERROR", log)
	case "info":
		return fmt.Sprintf(Fmt, INFO, "INFO", log)
	case "panic":
		return fmt.Sprintf(Fmt, PANIC, "PANIC", log)
	default:
		return fmt.Sprintf(Fmt, INFO, "INFO", log)
	}
}