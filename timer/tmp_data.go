/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package timer

import (
	"blog/models/dao/statistics_dao"
)

// 临时数据
var GlobalTmp TmpData

func GetUv() int {
	return GlobalTmp.TmpUv.View
}

func GetPv(name string) int {
	if v, ok := GlobalTmp.TmpPostView[name];ok {
		return v.View
	}else {
		return 0
	}
}

func ForceSaveUV() {
	var count int
	statistics_dao.StatisticsViewUpdate("all", count)
}

func ForceSavePv() {
	for _, k := range GlobalTmp.TmpPostView {
		statistics_dao.StatisticsViewUpdate(k.Name, k.View)
	}
}

func UpdateUv() {
	GlobalTmp.TmpUv.View += 1
}

func UpdatePv(name string) {
	if v, ok := GlobalTmp.TmpPostView[name];ok {
		v.View += 1
	}else {
		GlobalTmp.TmpPostView[name] = TmpPostView{
			Name: name,
			View: 1,
		}
	}
}

func ClearUv() {
	GlobalTmp.TmpUv.View = 0
}

func ClearPvOne(name string) {
	if v, ok := GlobalTmp.TmpPostView[name];ok {
		v.View = 0
	}
}

func ClearPv() {
	// 为避免重新内存浪费 删除所有键值对
	GlobalTmp.TmpPostView = map[string]TmpPostView{}
}