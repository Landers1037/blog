/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models/dao/post_dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPost(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		data, e := post_dao.PostQueryAll(map[string]interface{}{})
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "get post failed",
				"data": []string{},
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "get post success",
			"data": data,
		})
	}else {
		// 获取文章的详细内容
		data, e := post_dao.PostQuery(map[string]interface{}{"name": name})
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "get post failed",
				"data": "",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "get post success",
			"data": data,
		})
	}
}
