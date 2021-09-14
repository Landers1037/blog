/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

// RegisterCLI 注册命令行入口
func RegisterCLI() {
	app := &cli.App{
		Name:                   APP_NAME,
		Usage:                  APP_USAGE,
		Version:                Version,
		Description:            APP_DESCRIPTION,
		Commands:               AddAllCmds(),
		EnableBashCompletion:   false,
		Action:                 nil,
		CommandNotFound: func(c *cli.Context, s string) {
			fmt.Printf("cmd [%s] not found\n", s)
		},
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Authors:                []*cli.Author{
			{Author, Email},
			{Author2, Email2},
		},
		Copyright:              CopyRight,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("app failed to run: %s\n", err.Error())
		os.Exit(1)
	}
}
