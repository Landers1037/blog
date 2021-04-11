/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"blog/config"
	"blog/models"
	"blog/models/article"
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

// 文章导出
// 分为全部导出和单个文章的导出
// 不传递参数时为全部导出
// 生成文件的逻辑由前端处理
func ExportPosts(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
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
		return
	}
	var post article.DB_BLOG_POST
	e := models.BlogDB.Model(article.DB_BLOG_POST{}).Where("name = ?", name).First(&post).Error
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "export data fail",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "export data success",
		"data": post,
	})
	return
}

// 数据库导出
func ExportDataBase(c *gin.Context) {

}

// 数据库备份
// 全局的备份路径在预定义里为/home/backup
// 防止io出错 使用shell完成
func BackUpDB(c *gin.Context) {
	date := utils.GetDate()
	cmd := fmt.Sprintf("cp %s /home/backup/blog%s.db",
		config.Cfg.DB, date)
	e := exec.Command("bash", "-c", cmd).Run()
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "backup db failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "backup db success",
		"data": "success",
	})
}