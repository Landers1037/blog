package main

import (
	"blog_br_ng/config"
	"blog_br_ng/models"
	"blog_br_ng/routers"
	"blog_br_ng/utils/cmd"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main()  {
	// 注册额外参数
	flags := flags{}
	initFlag(&flags)
	if flags.flagVersion {
		fmt.Println("blog version: ", "3.2")
		fmt.Println("build: ", "20210312")
		os.Exit(0)
	}
	if flags.flagHelp {
		flag.Usage()
		os.Exit(0)
	}
	// 设置初始化
	configErr := config.InitCfg()
	if configErr != nil {
		os.Exit(1)
	}

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

	app := routers.InitRouter()

	// 测试模式下使用gin自带的web服务器启动
	if flags.flagTest {
		fmt.Println("以测试模式启动")
		fmt.Println("监听端口5000")
		config.Cfg.RunMode = "debug"
		app.Run(":5000")
	}else {
		if flags.flagPort != "" {
			cluster := config.Cfg.Cluster
			if cluster {
				fmt.Println("以集群模式启动")
				ports := strings.Fields(flags.flagPort)
				fmt.Println("监听端口: ", ports)
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
					fmt.Println("子服务运行成功 端口： " + <-portch)
				}
			}else {
				fmt.Println("以单实例模式启动")
				port := strings.Fields(flags.flagPort)[0]
				fmt.Println("监听端口: ", port)
				s := &http.Server{
					Addr: fmt.Sprintf(":%s", port),
					Handler: app,
					ReadTimeout: config.Cfg.ReadTimeout,
					WriteTimeout: config.Cfg.WriteTimeout,
					MaxHeaderBytes: 1 << 20,
				}
				err := s.ListenAndServe()
				if err != nil{
					fmt.Println(err)
				}
			}
		}else {
			// 为空时读取默认值
			port := config.Cfg.HTTPPort
			fmt.Println("以单实例模式启动")
			fmt.Println("监听端口: ", port)
			s := &http.Server{
				Addr: fmt.Sprintf(":%d", port),
				Handler: app,
				ReadTimeout: config.Cfg.ReadTimeout,
				WriteTimeout: config.Cfg.WriteTimeout,
				MaxHeaderBytes: 1 << 20,
			}
			err := s.ListenAndServe()
			if err != nil {
				fmt.Println(err)
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