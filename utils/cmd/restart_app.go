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

// RestartAPP 重启服务 先停止后启动
func RestartAPP(appName, appPid, appLog string) {
	// 如果未运行则直接启动
	f, e := ioutil.ReadFile(appPid)
	if e != nil {
		e := exec.Command("bash", "-c", fmt.Sprintf("nohup ./%s web start > %s 2>&1 & echo $! > %s", appName, appLog, appPid)).Run()
		if e != nil {
			fmt.Println(fmt.Sprintf("启动失败 应用: %s PID文件: %s", appName, appPid))
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("启动完毕 应用: %s PID文件: %s", appName, appPid))
		fmt.Println(fmt.Sprintf("日志记录于: %s", appLog))
	}else {
		pid := string(f)
		if checkProc(pid) {
			// 停止并启动
			e := exec.Command("bash", "-c", fmt.Sprintf("kill %s", pid)).Run()
			if e != nil {
				fmt.Println(fmt.Sprintf("停止失败 应用: %s PID文件: %s", appName, appPid))
				os.Exit(1)
			}
			fmt.Println(fmt.Sprintf("停止完毕 应用: %s PID文件: %s PID: %s", appName, appPid, pid))

			e = exec.Command("bash", "-c", fmt.Sprintf("nohup ./%s > %s 2>&1 & echo $! > %s", appName, appLog, appPid)).Run()
			if e != nil {
				fmt.Println(fmt.Sprintf("启动失败 应用: %s PID文件: %s", appName, appPid))
				os.Exit(1)
			}
			fmt.Println(fmt.Sprintf("启动完毕 应用: %s PID文件: %s", appName, appPid))
			fmt.Println(fmt.Sprintf("日志记录于: %s", appLog))
		}else {
			// 启动
			e := exec.Command("bash", "-c", fmt.Sprintf("nohup ./%s > %s 2>&1 & echo $! > %s", appName, appLog, appPid)).Run()
			if e != nil {
				fmt.Println(fmt.Sprintf("启动失败 应用: %s PID文件: %s", appName, appPid))
				os.Exit(1)
			}
			fmt.Println(fmt.Sprintf("启动完毕 应用: %s PID文件: %s", appName, appPid))
			fmt.Println(fmt.Sprintf("日志记录于: %s", appLog))
		}
	}

}
