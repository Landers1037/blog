/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package timer

import (
	"blog/logger"
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
				logger.BlogLogger.InfoF("定时任务 缓存中全局统计量: %d准备刷新缓存", GetUv())
				ForceSaveUV()
				ClearUv()
			case <-pvTicker.C:
				logger.BlogLogger.InfoF("定时任务 文章访问量准备刷新缓存")
				ForceSavePv()
				ClearPv()
			}
		}
	}()
}
