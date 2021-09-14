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
	"strings"
)

// StartAPP start app using nohup
func StartAPP(appName, appPid, appLog string) {
	// 检查pid是否在运行
	f, e := ioutil.ReadFile(appPid)
	pid := string(f)
	// 这里必须先关闭文件流
	if e != nil || !checkProc(pid) {
		// 不存在pid文件则默认未启动 使用nohup启动自身实例
		e := exec.Command("bash", "-c", fmt.Sprintf("nohup ./%s web start > %s 2>&1 & echo $! > %s", appName, appLog, appPid)).Run()
		if e != nil {
			fmt.Println(fmt.Sprintf("启动失败 应用: %s PID文件: %s", appName, appPid))
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("启动完毕 应用: %s PID文件: %s", appName, appPid))
		fmt.Println(fmt.Sprintf("日志记录于: %s", appLog))
	}else if checkProc(pid) {
		// pid正在运行
		fmt.Println("应用正在运行")
	}else {

	}
}

func checkProc(pid string) bool {
	pid = strings.Trim(pid, "\n")
	pid = strings.TrimSpace(pid)
	e := exec.Command("bash", "-c", fmt.Sprintf("cat /proc/%s/status", pid)).Run()
	if e != nil {
		return false
	}
	return true

}
