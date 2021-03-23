/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package response

type RES_ZHUANLAN struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Posts []string `json:"posts"`
	Content string `json:"content"`
}
