/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package response

// RES_POST 专用于响应的结构体
type RES_POST struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Abstract string `json:"abstract"`
	Tags string `json:"tags"`
}

// RES_POST_MORE 文章详情
type RES_POST_MORE struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Abstract string `json:"abstract"`
	Content string `json:"content"`
	Tags string `json:"tags"`
	Categories string `json:"categories"`
}

// RES_POST_BROTHER 仅用于获取上下篇
type RES_POST_BROTHER struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Pin int `json:"pin"`
}