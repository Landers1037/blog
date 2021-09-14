/*
Name: blog
File: build_archive.go
Author: Landers
*/

package utils

import (
	"blog/models/response"
	"sort"
	"strings"
)

type archives struct {
	Date string `json:"date"`
	Count int `json:"count"`
}

// BuildArchive 构建日期归档的数组
func BuildArchive(posts []response.RES_POST) []archives {
	flagMap := make(map[string]int)
	for _, p := range posts {
		date := strings.Split(p.Date, "-")
		if len(date) > 1 {
			// 保留年月
			ym := strings.Join(date[0:2], "-")
			// 存入标志map中 存在key时结构体值加1
			if _, ok := flagMap[ym]; !ok {
				flagMap[ym] = 1
			}else {
				flagMap[ym] += 1
			}
		}
	}
	// 对字典key排序后输出到结构体数组
	var keys []string
	for key := range flagMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// 生成结构体
	var ar []archives
	for _, k := range keys {
		ar = append(ar, archives{
			Date:  k,
			Count: flagMap[k],
		})
	}

	return ar
}
