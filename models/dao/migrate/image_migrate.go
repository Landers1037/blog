/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package migrate

import (
	"blog/logger"
	"blog/models"
	"blog/models/article"
	"strings"
)

// 专用于图片资源的地址迁移工作
func ImageMigrate(old, new string) error {
	var posts []article.DB_BLOG_POST
	// 默认会自动对全部的文章遍历
	e := models.BlogDB.Model(&article.DB_BLOG_POST{}).Find(&posts).Error
	if e != nil {
		return e
	}
	// 使用strings自动替换全部符合条件的字符串
	for _, p := range posts {
		logger.BlogLogger.InfoF("开始迁移%s", p.Name)
		var oldAbs, oldContent string
		oldAbs = strings.ReplaceAll(p.Abstract, old, new)
		oldContent = strings.ReplaceAll(p.Content, old, new)
		p.Abstract = oldAbs
		p.Content = oldContent
		e := models.BlogDB.Save(&p).Error
		if e != nil {
			logger.BlogLogger.InfoF("迁移%s失败: %s", p.Name, e.Error())
		}
	}
	return nil
}
