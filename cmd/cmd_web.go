/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package cmd

import (
	"blog/config"
	"blog/logger"
	"blog/models"
	"blog/models/dao/crdb"
	"blog/routers"
	"blog/timer"
	"blog/utils/settings"
	"context"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

// web功能页面入口
func AddWebCmds() []*cli.Command {
	return []*cli.Command{
		{
			Name:                   "web",
			Aliases:                []string{"w", "serve", "server"},
			Usage:                  "启动blog的web服务",
			Description:            "启动web服务",
			Category:               "Run a web service",
			Action:                 nil,
			OnUsageError:           nil,
			Subcommands: []*cli.Command{
				{
					Name:                   "test",
					Aliases:                []string{"t"},
					Usage:                  "测试模式启动服务",
					Category:               "Run a web service",
					Action: func(c *cli.Context) error {
						port := c.Int("port")
						conf := c.String("conf")
						// 初始化日志记录器
						logger.PrintLogo()
						// 设置初始化
						config.InitCfg(conf)
						logger.InitLogger(conf)
						// 初始化数据库
						models.InitDB()
						// init create db
						dbErr := crdb.InitTables()
						if dbErr != nil {
							logger.BlogLogger.PanicF("数据库创建失败")
						}
						// 初始化定时器
						timer.InitTimer()
						// 初始化路由
						config.Cfg.RunMode = "debug"
						app := routers.InitRouter()
						logger.BlogLogger.InfoF("以测试模式启动")
						logger.BlogLogger.InfoF("监听端口: %d", port)
						logger.BlogLogger.InfoF("配置文件: %s", conf)

						return app.Run(fmt.Sprintf(":%d", port))
					},
					OnUsageError:           nil,
					Flags:                  []cli.Flag{
						&cli.IntFlag{
							Name:    "port",
							Aliases: []string{"p"},
							Usage:   "设置启动端口",
							Value:   settings.APP_PORT,
						},
						&cli.StringFlag{
							Name:    "conf",
							Aliases: []string{"c"},
							Usage:   "设置配置文件路径",
							Value:   settings.APP_FILE,
						},
					},
				},
				{
					Name:                   "start",
					Aliases:                []string{"run"},
					Usage:                  "正常模式启动服务",
					Category:               "Run a web service",
					Action: func(c *cli.Context) error {
						port := c.Int("port")
						conf := c.String("conf")
						// 初始化日志记录器
						logger.PrintLogo()
						// 设置初始化
						config.InitCfg(conf)
						logger.InitLogger(conf)
						// 初始化数据库
						models.InitDB()
						// init create db
						dbErr := crdb.InitTables()
						if dbErr != nil {
							logger.BlogLogger.PanicF("数据库创建失败")
						}
						// 初始化定时器
						timer.InitTimer()
						// 初始化路由
						app := routers.InitRouter()
						logger.BlogLogger.InfoF("以单实例模式启动")
						logger.BlogLogger.InfoF("监听端口: %d", port)
						logger.BlogLogger.InfoF("配置文件: %s", conf)
						s := &http.Server{
							Addr: fmt.Sprintf(":%d", port),
							Handler: app,
							ReadTimeout: config.Cfg.ReadTimeout,
							WriteTimeout: config.Cfg.WriteTimeout,
							MaxHeaderBytes: 1 << 20,
						}

						go func() {
							quit := make(chan os.Signal)
							signal.Notify(quit, os.Interrupt)
							// 监听退出信号
							<-quit
							logger.BlogLogger.InfoF("接收到终止信号")
							ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
							defer cancel()
							if err := s.Shutdown(ctx); err != nil {
								logger.BlogLogger.InfoF("服务关闭中")
								logger.BlogLogger.InfoF("即将退出 开始保存临时数据")
								timer.ForceSaveUV()
								timer.ForceSavePv()
								models.CloseDB()
							}
						}()
						err := s.ListenAndServe()
						if err != nil {
							if err == http.ErrServerClosed {
								logger.BlogLogger.InfoF("服务已停止")
							}else {
								logger.BlogLogger.ErrorF(err.Error())
							}

							// 开始写入全部缓存数据
							logger.BlogLogger.InfoF("即将退出 开始保存临时数据")
							timer.ForceSaveUV()
							timer.ForceSavePv()
							models.CloseDB()
						}
						if err == http.ErrServerClosed {
							return nil
						}
						return err
					},
					OnUsageError:           nil,
					Flags:                  []cli.Flag{
						&cli.IntFlag{
							Name:    "port",
							Aliases: []string{"p"},
							Usage:   "设置启动端口",
							Value:   settings.APP_PORT,
						},
						&cli.StringFlag{
							Name:    "conf",
							Aliases: []string{"c"},
							Usage:   "设置配置文件路径",
							Value:   settings.APP_FILE,
						},
					},
				},
				{
					Name:                   "cluster",
					Aliases:                []string{"clu"},
					Usage:                  "集群模式启动服务",
					Category:               "Run a web service",
					Action: func(c *cli.Context) error {
						port := c.String("port")
						conf := c.String("conf")
						// 初始化日志记录器
						logger.PrintLogo()
						// 设置初始化
						config.InitCfg(conf)
						logger.InitLogger(conf)
						// 初始化数据库
						models.InitDB()
						// init create db
						dbErr := crdb.InitTables()
						if dbErr != nil {
							logger.BlogLogger.PanicF("数据库创建失败")
						}
						// 初始化定时器
						timer.InitTimer()
						// 初始化路由
						app := routers.InitRouter()
						// 同时需要clusetr配置开启
						if !config.Cfg.Cluster {
							return errors.New("配置文件需要cluster参数")
						}
						logger.BlogLogger.InfoF("以集群模式启动")
						ports := strings.Fields(port)
						logger.BlogLogger.InfoF("监听端口: %s", strings.Join(ports, " "))
						logger.BlogLogger.InfoF("配置文件: %s", conf)
						portch := make(chan string)
						for _, p := range ports {
							s := &http.Server{
								Addr: fmt.Sprintf(":%s", p),
								Handler: app,
								ReadTimeout: config.Cfg.ReadTimeout,
								WriteTimeout: config.Cfg.WriteTimeout,
								MaxHeaderBytes: 1 << 20,
							}
							go func(s *http.Server, port string) {
								err := s.ListenAndServe()
								portch<-port
								if err != nil{
									if err == http.ErrServerClosed {
										logger.BlogLogger.InfoF("服务已停止")
										models.CloseDB()
									}else {
										logger.BlogLogger.ErrorF(err.Error())
										models.CloseDB()
									}
								}
							}(s, port)
						}

						// 阻塞进程
						for range ports {
							logger.BlogLogger.InfoF("子服务运行成功 端口： " + <-portch)
						}

						return nil
					},
					OnUsageError:           nil,
					Flags:                  []cli.Flag{
						&cli.StringFlag{
							Name:    "port",
							Aliases: []string{"p"},
							Usage:   "设置启动端口",
							Value:   strconv.Itoa(settings.APP_PORT),
						},
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
