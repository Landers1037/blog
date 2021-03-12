/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"blog/models"
)

// 总views name=ALL/all
type DB_BLOG_VIEWS struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	View int `json:"view"`
}

// 获取全局总访问量
func GetViewAllFromdb() (uv DB_BLOG_VIEWS) {
	models.BlogDB.First(&uv,DB_BLOG_VIEWS{Name: "all"})
	return
}

// 获取文章的访问量
func GetViewFromdb(name string) (uv DB_BLOG_VIEWS) {
	// 保证了一定查询成功 不成功返回空结构体
	var postv DB_BLOG_VIEWS
	e := models.BlogDB.First(&postv, DB_BLOG_VIEWS{Name: name}).Error
	if e != nil {
		return DB_BLOG_VIEWS{View: 0}
	}
	return postv
}

// 全局的总访问量
func AddViewAll(num int) {
	var uv DB_BLOG_VIEWS
	//每5个请求为一个缓存
	e := models.BlogDB.Select("view").First(&uv, DB_BLOG_VIEWS{Name: "all"}).Error
	if e != nil {
		return
	}
	num = num+uv.View
	models.BlogDB.First(&uv, DB_BLOG_VIEWS{Name: "all"}).Update("view", num)
	return
}

// 添加文章的view
func AddView(num int, name string) {
	// 为避免高并发频繁写入 做时间间隔桶机制
	var uv DB_BLOG_VIEWS
	e := models.BlogDB.Select("view").First(&uv, DB_BLOG_VIEWS{Name: name}).Error
	if e != nil {
		return
	}
	num = num+uv.View
	models.BlogDB.First(&uv, DB_BLOG_VIEWS{Name: name}).Update("view", num)
	return
}