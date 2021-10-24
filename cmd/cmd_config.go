/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

// Package cmd 命令行入口注册
package cmd

import (
	"blog/config"
	"blog/utils/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
)

// AddConfigCmds 配置文件入口
func AddConfigCmds() []*cli.Command {
	return []*cli.Command{
		{
			Name:         "config",
			Aliases:      []string{"conf", "c"},
			Usage:        "应用配置",
			Category:     "App configs",
			Action:       nil,
			OnUsageError: nil,
			Subcommands: []*cli.Command{
				{
					Name:     "show",
					Usage:    "查看应用配置",
					Category: "App configs",
					Action: func(c *cli.Context) error {
						conf := c.String("conf")
						config.InitCfg(conf)

						// 美化输出
						b, e := json.Marshal(&config.Cfg)
						if e != nil {
							return e
						}
						var out bytes.Buffer
						e = json.Indent(&out, b, "", "  ")
						fmt.Println(out.String())
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "conf",
							Aliases: []string{"c"},
							Usage:   "设置配置文件路径",
							Value:   settings.APP_FILE,
						},
					},
				},
				{
					Name:     "router",
					Usage:    "查看静态路由",
					Category: "App configs",
					Action: func(c *cli.Context) error {
						router := c.String("router")
						if _, e := os.Stat(router); os.IsNotExist(e) {
							return nil
						}

						// 美化输出
						res, e := ioutil.ReadFile(router)
						if e != nil {
							return e
						}

						var out bytes.Buffer
						e = json.Indent(&out, res, "", "  ")
						fmt.Println(out.String())
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "router",
							Aliases: []string{"r"},
							Usage:   "设置静态路由路径",
							Value:   "conf/router.json",
						},
					},
				},
				{
					Name:     "edit",
					Usage:    "编辑配置文件",
					Category: "App configs",
					Action: func(c *cli.Context) error {
						conf := c.String("conf")
						if _, e := os.Stat(conf); os.IsNotExist(e) {
							fmt.Printf("配置文件%s不存在\n", conf)
						}

						// 使用vi创建
						cmd := exec.Command("vi", conf)
						cmd.Stdin = os.Stdin
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stdout
						return cmd.Run()
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "conf",
							Aliases: []string{"c"},
							Usage:   "设置配置文件路径",
							Value:   settings.APP_FILE,
						},
					},
				},
			},
		},
	}
}
