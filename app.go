package main

import (
	"blog/config"
	"blog/logger"
	"blog/models"
	"blog/routers"
	"blog/timer"
	"blog/utils/cmd"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

var version = "3.6"
var build = "20210326"

func main()  {
	// 注册额外参数
	flags := flags{}
	initFlag(&flags)
	if flags.flagVersion {
		fmt.Println("blog version: ", version)
		fmt.Println("build: ", build)
		os.Exit(0)
	}
	if flags.flagHelp {
		flag.Usage()
		os.Exit(0)
	}
	// 初始化日志记录器
	logger.PrintLogo()
	logger.InitLogger()
	// 设置初始化
	config.InitCfg()

	if flags.flagSt != "" {
		// 仅支持linux
		avil1 := runtime.GOOS != "linux"
		avil2 := runtime.GOOS != "unix"
		avil3 := runtime.GOOS != "darwin"
		if avil1 && avil2 && avil3 {
			fmt.Println("你的操作系统不支持此特性")
			os.Exit(0)
		}
		appName := config.Cfg.AppName
		appPid := config.Cfg.APPPid
		appLog := config.Cfg.APPLog
		switch flags.flagSt {
		case "start":
			cmd.StartAPP(appName, appPid, appLog)
		case "stop":
			cmd.StopAPP(appName, appPid)
		case "restart":
			cmd.RestartAPP(appName, appPid, appLog)
		case "status":
			cmd.StatusAPP(appName, appPid)
		default:
			cmd.StatusAPP(appName, appPid)
		}

		os.Exit(0)
	}
	// 初始化数据库
	models.InitDB()
	// 初始化定时器
	timer.InitTimer()
	// 初始化路由
	app := routers.InitRouter()


	// 测试模式下使用gin自带的web服务器启动
	if flags.flagTest {
		logger.BlogLogger.InfoF("以测试模式启动")
		logger.BlogLogger.InfoF("监听端口5000")
		config.Cfg.RunMode = "debug"
		app.Run(":5000")
	}else {
		if flags.flagPort != "" {
			cluster := config.Cfg.Cluster
			if cluster {
				logger.BlogLogger.InfoF("以集群模式启动")
				ports := strings.Fields(flags.flagPort)
				logger.BlogLogger.InfoF("监听端口: %s", strings.Join(ports, " "))
				portch := make(chan string)
				for _, port := range ports {
					app := routers.InitRouter()
					s := &http.Server{
						Addr: fmt.Sprintf(":%s", port),
						Handler: app,
						ReadTimeout: config.Cfg.ReadTimeout,
						WriteTimeout: config.Cfg.WriteTimeout,
						MaxHeaderBytes: 1 << 20,
					}
					go func(s *http.Server, port string) {
						err := s.ListenAndServe()
						portch<-port
						if err != nil{
							fmt.Println(err)
						}
					}(s, port)
				}

				// 阻塞进程
				for _, _ = range ports {
					logger.BlogLogger.InfoF("子服务运行成功 端口： " + <-portch)
				}
			}else {
				logger.BlogLogger.InfoF("以单实例模式启动")
				port := strings.Fields(flags.flagPort)[0]
				logger.BlogLogger.InfoF("监听端口: %s", port)
				s := &http.Server{
					Addr: fmt.Sprintf(":%s", port),
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
					}
				}()
				err := s.ListenAndServe()

				if err != nil {
					logger.BlogLogger.ErrorF(err.Error())
					// 开始写入全部缓存数据
					logger.BlogLogger.InfoF("即将退出 开始保存临时数据")
					timer.ForceSaveUV()
					timer.ForceSavePv()
				}
			}
		}else {
			// 为空时读取默认值
			port := config.Cfg.HTTPPort
			logger.BlogLogger.InfoF("以单实例模式启动")
			logger.BlogLogger.InfoF("监听端口: %d", port)
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
				}
			}()
			err := s.ListenAndServe()

			if err != nil {
				logger.BlogLogger.ErrorF(err.Error())
				// 开始写入全部缓存数据
				logger.BlogLogger.InfoF("即将退出 开始保存临时数据")
				timer.ForceSaveUV()
				timer.ForceSavePv()
			}
		}
	}
}

type flags struct {
	flagTest bool
	flagReload bool
	flagPort string
	flagVersion bool
	flagHelp bool
	flagSt string
}

func initFlag(flags *flags) {
	ftest := flag.Bool("t", false, "是否开启测试模式")
	fport := flag.String("p", "", "选择监听端口")
	freload := flag.Bool("r", false, "重载配置文件")
	fversion := flag.Bool("v", false, "查看版本信息")
	fst := flag.String("s", "", "操作服务进程")
	fhelp := flag.Bool("h", false, "查看帮助信息")

	flag.Parse()
	flags.flagTest = *ftest
	flags.flagPort = *fport
	flags.flagReload = *freload
	flags.flagVersion = *fversion
	flags.flagHelp = *fhelp
	flags.flagSt = *fst
}