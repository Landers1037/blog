/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"blog/config"
	"crypto/sha256"
	"fmt"
)

// 内部实现逻辑很简单只需要比对加密后的串是否一致
// 加密方式sha256
func AdminEncrypt() string {
	// 默认为每天更换一次
	date := GetDate()
	sha := sha256.New()
	sha.Write([]byte(fmt.Sprintf("%s%sAt%s", config.Cfg.AdminName, config.Cfg.AdminPwd, date)))
	res := sha.Sum(nil)
	return fmt.Sprintf("%x", res)
}
