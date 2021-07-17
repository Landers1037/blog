/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package middleware

import (
	"blog/config"
	"blog/models/dao/statistics_dao"
	"blog/timer"
	"github.com/gin-gonic/gin"
	"time"
)

// 文章开启阅读访问量统计
// 防止频繁写库 按时写入
// 防止并发场景下多次计时 采用每n分钟刷新机制

func PostView() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if config.Cfg.PostViewFlag && name != "" && c.FullPath() == "/api/article/post" {
			timer.UpdatePv(name)
			if time.Now().Minute()%5 == 0 {
				// 只有访问路径中存在name时才会去更新
				statistics_dao.StatisticsViewUpdate(name, timer.GetPv(name))
				timer.ClearPvOne(name)
			}
		}

		c.Next()
	}
}