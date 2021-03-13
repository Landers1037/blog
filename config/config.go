/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package config

import (
	"blog_br_ng/utils/settings"
)

var Cfg settings.Cfg

func InitCfg() error {
	var e error
	Cfg, e = settings.InitCfg()
	if e != nil {
		return e
	}
	return nil
}
