/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"blog/config"
)

// Conetent2Abs 截取正文内容作为博客摘要
func Conetent2Abs(abs, content string) string {
	if config.Cfg.UseContentAsAbs && abs == "" {
		if len(content) >= config.Cfg.MaxContentLength {
			return content[0:config.Cfg.MaxContentLength]
		}
		return content
	}else if abs == "" {
		return config.Cfg.CustomEmptyAbs
	}else {
		return abs
	}
}
