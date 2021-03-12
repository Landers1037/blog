/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func StopAPP(appName, appPid string) {
	// 检查pid是否在运行
	f, e := ioutil.ReadFile(appPid)
	if e != nil {
		fmt.Println(fmt.Sprintf("停止失败 应用: %s PID文件: %s不存在", appName, appPid))
	}
	pid := string(f)
	if checkProc(pid) {
		// 不存在pid文件则默认未启动 使用nohup启动自身实例
		e := exec.Command("bash", "-c", fmt.Sprintf("kill %s", pid)).Run()
		if e != nil {
			fmt.Println(fmt.Sprintf("停止失败 应用: %s PID文件: %s", appName, appPid))
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("停止完毕 应用: %s PID文件: %s PID: %s", appName, appPid, pid))
	} else {
		fmt.Println(fmt.Sprintf("停止完毕 应用: %s PID文件: %s PID: %s", appName, appPid, pid))
	}
}
