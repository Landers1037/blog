/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package statistic

import (
	"blog/models/dao/statistics_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所有统计访问次数
func GetViews(c *gin.Context) {
	uv := statistics_dao.StatisticViewQuery("all")
	c.JSON(http.StatusOK, uv)
}
