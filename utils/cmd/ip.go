/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package cmd
//ip count
import (
	"os/exec"
)
func cmdIp() string {
	cmdout := exec.Command("/home/web/blog_br_ng/shell/ip.sh")
	op,err:= cmdout.Output()
	if err!=nil{
		return "getip failed"
	}
	return string(op)
}

