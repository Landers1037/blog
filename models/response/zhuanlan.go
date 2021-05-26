/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package response

// 返回前端显示用
type RES_ZHUANLAN1 struct {
	Link string `json:"link"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []string `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}

// 专栏详情 包含文章信息
type Zpost struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Abstract string `json:"abstract"`
}

type RES_ZHUANLAN_DETAIL struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []Zpost `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}

// 控制面板用
type RES_ZHUANLAN2 struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []string `json:"posts"`
	Content string `json:"content"` // 专栏的简要描述
}