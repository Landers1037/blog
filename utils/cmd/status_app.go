/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// StatusAPP 查看状态 基于pid文件
// pid找不到才会去寻找程序自身
func StatusAPP(appName, appPid string) {
	f, e := ioutil.ReadFile(appPid)
	if e != nil {
		fmt.Println(fmt.Sprintf("应用: %s PID文件: %s不存在", appName, appPid))
		fmt.Println("开始搜寻应用进程")
		opt, e := exec.Command("bash", "-c", fmt.Sprintf("pgrep %s", appName)).Output()
		if e != nil {
			fmt.Println(fmt.Sprintf("状态获取失败 应用: %s", appName))
		}else {
			if len(opt) == 0 {
				fmt.Println(fmt.Sprintf("状态获取成功 应用: %s未运行", appName))
			}else {
				fmt.Println(fmt.Sprintf("状态获取成功 应用: %s运行于PID: %s", appName, string(opt)))
			}
		}
	}else {
		pid := string(f)
		if checkProc(pid) {
			fmt.Println(fmt.Sprintf("状态获取成功 应用: %s运行于PID: %s", appName, pid))
		}else {
			fmt.Println(fmt.Sprintf("状态获取成功 应用: %s未运行", appName))
		}
	}
}
