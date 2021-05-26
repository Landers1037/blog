/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package config

import (
	"blog/logger"
	"blog/utils/settings"
)

var Cfg settings.Cfg

func InitCfg() {
	var e error
	Cfg, e = settings.InitCfg()
	if e != nil {
		logger.BlogLogger.PanicF("程序配置文件读取失败 %s", e.Error())
	}
}
