/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package timer

import (
	"fmt"
	"time"
)

// 启动timer
// 必须在应用启动时执行
func InitTimer() {
	// 字典具有特殊性 需要在这里初始化
	GlobalTmp.TmpUv = TmpUv{View: 0}
	GlobalTmp.TmpPostView = make(map[string]TmpPostView)

	uvTicker := time.NewTicker(TIME_UV * time.Second)
	pvTicker := time.NewTicker(TIME_POST_VIEW * time.Second)
	go func() {
		for {
			select {
			case <-uvTicker.C:
				fmt.Println("uv now: ", GetUv())
				ForceSaveUV()
				ClearUv()
			case <-pvTicker.C:
				fmt.Println("pv now", GetPv("pin"))
				ForceSavePv()
				ClearPv()
			}
		}
	}()
}
