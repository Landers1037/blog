/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

// GetSortPost 获取post的排序标准
func GetSortPost(sortBy string, reverse bool) string {
	var sort string
	switch sortBy {
	case "id", "ID":
		sort = "primary_id"
	case "date":
		sort = "date_plus"
	case "update", "update_date":
		sort = "update"
	case "title":
		sort = "title"
	case "name":
		sort = "name"
	default:
		sort = "primary_id"
	}

	if reverse {
		return sort + " " + "desc"
	}
	return sort
}
