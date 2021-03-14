/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"time"
)

// 获取准确时间
func GetDatePlus() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取日期 年月日
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

// 获取年月
func GetDate2() string {
	return time.Now().Format("2006-01")
}

// 获取年
func GerDateYear() string {
	return string(time.Now().Year())
}