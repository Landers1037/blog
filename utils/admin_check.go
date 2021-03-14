/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

func AdminCheck(token string) bool {
	uEncrypt := AdminEncrypt()
	if token == uEncrypt {
		return true
	}
	return false
}
