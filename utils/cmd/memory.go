/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package cmd

import (
	"os/exec"
)
func cmdMem() string {
	cmdout := exec.Command("/home/web/blog/shell/mem.sh")
	op,err:= cmdout.Output()
	if err!=nil{
		return "getmem failed"
	}
	return string(op)
}
