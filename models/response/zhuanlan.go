/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package response

// RES_ZHUANLAN1 返回前端显示用
type RES_ZHUANLAN1 struct {
	Link string `json:"link"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []string `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}

// Zpost 专栏详情 包含文章信息
type Zpost struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Abstract string `json:"abstract"`
	Tags string `json:"tags"`
}

type RES_ZHUANLAN_DETAIL struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []Zpost `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}

// RES_ZHUANLAN2 控制面板用
type RES_ZHUANLAN2 struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []string `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}