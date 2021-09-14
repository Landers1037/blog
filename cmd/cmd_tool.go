/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

/*
Package cmd
初始化工具入口
包括：
初始化数据库
初始化配置文件
生成router文件
新增博客文章
修改博客模板
获取博客模板
博客链接迁移
*/
package cmd

import (
	"blog/logger"
	"blog/models/dao/crdb"
	"blog/models/dao/migrate"
	"blog/utils/settings"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"time"
)

// AddToolCmds 添加工具命令
func AddToolCmds() []*cli.Command {
	return []*cli.Command{
		{
			Name:                   "tool",
			Aliases:                []string{"t"},
			Usage:                  "博客配套工具",
			Category:               "Tools of blog",
			Subcommands:            []*cli.Command{
				{
					Name:                   "db",
					Usage:                  "初始化数据库",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						fmt.Println("开始创建数据库")
						db, err := gorm.Open("sqlite3", "blog.db")
						if err != nil {
							fmt.Printf("数据库连接失败: %s\n", err.Error())
							return err
						}
						db.SingularTable(true)
						for table, tableStruct := range crdb.Tables {
							if db.HasTable(tableStruct) {
								fmt.Printf("表%s已经存在\n", table)
								continue
							}
							fmt.Printf("初始化表%s\n", table)
							e := db.CreateTable(tableStruct).Error
							if e != nil {
								fmt.Printf("%s表创建失败%s\n", table, e.Error())
								return e
							}
						}
						return nil
					},
				},
				{
					Name:                   "conf",
					Usage:                  "初始化配置",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						fmt.Println("开始创建配置")
						file := c.String("file")
						err := settings.SaveConfig(file)
						if err == nil {
							fmt.Printf("配置初始化完毕%s\n", file)
						}
						return err
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "file",
							Aliases: []string{"f"},
							Usage:   "指定配置文件",
							Value:   settings.APP_FILE,
						},
					},
				},
				{
					Name:                   "new",
					Usage:                  "新建文档",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						// 第一个参数为文件名
						name := c.Args().First()
						userDir, e := os.UserHomeDir()
						if e != nil {
							fmt.Printf("无法打开用户目录%s\n", e.Error())
							return e
						}
						metaFile := path.Join(userDir, "blog.meta")
						var metainfo string

						_, e = os.Stat(name + ".md")
						if e != nil {
							if os.IsExist(e) {
								return errors.New(fmt.Sprintf("file %s exist", name + ".md"))
							}else {
								// file not exist
								f, e := os.OpenFile(name + ".md", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
								if e != nil {
									return e
								}
								metaDefault := "---\ntitle: %s\nname: %s\ndate: %s\ntags: []\ncategories: []\nabstract: \n---\n<!--more-->"
								_, e = os.Stat(metaFile)
								if e != nil {
									// 不存在时写入默认的
									_ = ioutil.WriteFile(metaFile, []byte(metaDefault), 0644)
								}
								data, e := ioutil.ReadFile(metaFile)
								if e != nil {
									metainfo = metaDefault
								}else {
									metainfo = string(data)
								}
								dateString := time.Now().Format("2006-01-02 15:04:05")
								metainfo = fmt.Sprintf(metainfo,name, name, dateString)
								_, e = f.WriteString(metainfo)
								if e != nil {
									fmt.Printf("文档%s.md创建失败\n", name)
								}else {
									fmt.Printf("文档%s.md创建完毕\n", name)
								}
								return e
							}
						}
						return nil
					},
				},
				{
					Name:                   "temp",
					Usage:                  "修改模板",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						n := c.Bool("n")
						// 保证文件权限保存至~/blog.meta
						userDir, e := os.UserHomeDir()
						if e != nil {
							fmt.Printf("无法打开用户目录%s\n", e.Error())
							return e
						}
						metaFile := path.Join(userDir, "blog.meta")
						metaDefault := "---\ntitle: %s\nname: %s\ndate: %s\ntags: []\ncategories: []\nabstract: \n---\n<!--more-->"
						_, e = os.Stat(metaFile)
						if n && e != nil {
							// 不存在时写入默认的
							fmt.Printf("模板%s不存在 ,开始创建\n", metaFile)
							e = ioutil.WriteFile(metaFile, []byte(metaDefault), 0644)

							return e
						}

						cmd := exec.Command("vi", metaFile)
						cmd.Stdin = os.Stdin
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stdout
						return cmd.Run()
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:  "n",
							Usage: "初始化模板",
							Value: false,
						},
					},
				},
				{
					Name:                   "migrate",
					Usage:                  "迁移链接[旧] [新]",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						args := c.Args().Slice()
						fmt.Println("开始图片链接迁移")
						if len(args) <= 1 {
							fmt.Println("参数不足")
							return nil
						}
						fmt.Printf("开始迁移 旧链接: %s 新链接: %s\n", args[0], args[1])
						e := migrate.ImageMigrate(args[0], args[1])
						if e != nil {
							fmt.Println("迁移失败")
						}else {
							fmt.Println("迁移成功")
						}
						return e
					},
				},
				{
					Name:                   "logo",
					Usage:                  "展示logo",
					Category:               "Tools of blog",
					Action: func(c *cli.Context) error {
						logger.PrintLogo()
						return nil
					},
				},
			},
		},
	}
}
