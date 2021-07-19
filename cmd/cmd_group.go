/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package cmd

import (
	"github.com/urfave/cli/v2"
)

// 整合全部的cmds分组
func AddAllCmds() []*cli.Command {
	group :=  []*cli.Command{}
	group = append(group, AddWebCmds()...)
	group = append(group, AddConfigCmds()...)
	group = append(group, AddToolCmds()...)
	group = append(group, AddServiceCmds()...)

	return group
}
