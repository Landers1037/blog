/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package timer

// 内置的全局定时器
// 根据规则定时刷新统计值
// 基于全局的struct 节约内存的占用

const (
	TIME_UV = 600
	TIME_POST_VIEW = 3600
)

type TmpData struct {
	TmpUv TmpUv
	TmpPostView map[string]TmpPostView
}

type TmpUv struct {
	View int
}

type TmpPostView struct {
	Name string
	View int
}
