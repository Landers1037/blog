/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package cmd
//ip count
import (
	"os/exec"
)
func cmdIp() string {
	cmdout := exec.Command("/home/web/blog/shell/ip.sh")
	op,err:= cmdout.Output()
	if err!=nil{
		return "getip failed"
	}
	return string(op)
}

