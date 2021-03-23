/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"fmt"
	"time"
)

// 获取准确时间 "2006-01-02 15:04:05"
func GetDatePlus() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取准确时间的int 用于前端解析
func GetDateTime() string {

	return fmt.Sprintf("%d", time.Now().Unix())
}

// 获取日期 年月日 "2006-01-02"
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

// 获取年月 "2006-01"
func GetDate2() string {
	return time.Now().Format("2006-01")
}

// 获取年 "2006"
func GerDateYear() string {
	return string(time.Now().Year())
}

// 校验时间字符串合法性
func CheckDate(date string) bool {
	_, e := time.Parse("2006-01-02", date)
	if e != nil {
		return false
	}
	return true
}

func CheckDatePlus(date string) bool {
	_, e := time.Parse("2006-01-02 15:04:05", date)
	if e != nil {
		return false
	}
	return true
}