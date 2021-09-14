/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package cmd

import "runtime"

func GetGID() int {
	n := runtime.NumGoroutine()
	return n
}
