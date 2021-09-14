/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"blog/config"
)

// GetSortMessage 留言使用id user date排序
// 默认id
func GetSortMessage() string {
	var sort string
	switch config.Cfg.SortMessageBy {
	case "id", "ID":
		sort = "primary_id"
	case "user":
		sort = "user"
	case "date":
		sort = "date"
	default:
		sort = "primary_id"
	}

	if config.Cfg.SortMessageReverse {
		return sort + " " + "desc"
	}
	return sort
}
