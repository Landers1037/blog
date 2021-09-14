/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package cmd

func Sh(cmd string) string {
	var res string
	switch cmd {
		case "mem": res = cmdMem()
		case "ip": res = cmdIp()

	}
	return res
}
