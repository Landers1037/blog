/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/models"
	"blog/models/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文章导出
func ExportPosts(c *gin.Context) {
	var posts []article.DB_BLOG_POST
	e := models.BlogDB.Model(article.DB_BLOG_POST{}).Find(&posts).Error
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "export data fail",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "export data success",
		"data": posts,
	})
}
