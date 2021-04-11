/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"blog/logger"
)

func AdminCheck(token string) bool {
	uEncrypt := AdminEncrypt()
	logger.BlogLogger.InfoF("用户鉴权 内部加密密钥%s", uEncrypt)
	if token == uEncrypt {
		return true
	}
	logger.BlogLogger.InfoF("用户鉴权失败 外部加密密钥%s", token)
	return false
}
