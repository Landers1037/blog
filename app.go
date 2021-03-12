package main

import (
	"blog/config"
	"blog/models"
	"blog/routers"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main()  {
	// 注册额外参数
	flags := flags{}
	initFlag(&flags)
	if flags.flagVersion {
		fmt.Println("blog version: ", "3.1")
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
			f, _ := os.OpenFile("1.txt", os.O_APPEND|os.O_RDWR, 0666)
			fmt.Fprintln(f,"cluster")
			if cluster {
				fmt.Println("以集群模式启动")
				ports := strings.Fields(flags.flagPort)
				fmt.Println("监听端口: ", ports)
				for _, port := range ports {
					go func(port string) {
						f, _ := os.OpenFile("1.txt", os.O_APPEND|os.O_RDWR, 0666)
						fmt.Fprintln(f,"opop")
						//app := routers.InitRouter()
						//e := app.Run(":%s", port)
						fmt.Fprintln(f,port)
						//s := &http.Server{
						//	Addr: fmt.Sprintf(":%s", port),
						//	Handler: app,
						//	ReadTimeout: config.Cfg.ReadTimeout,
						//	WriteTimeout: config.Cfg.WriteTimeout,
						//	MaxHeaderBytes: 1 << 20,
						//}
						//err := s.ListenAndServe()
						//if err != nil{
						//	fmt.Println(err)
						//}
					}(port)
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
}

func initFlag(flags *flags) {
	ftest := flag.Bool("t", false, "是否开启测试模式")
	fport := flag.String("p", "", "选择监听端口")
	freload := flag.Bool("r", false, "重载配置文件")
	fversion := flag.Bool("v", false, "查看版本信息")
	fhelp := flag.Bool("h", false, "查看帮助信息")

	flag.Parse()
	flags.flagTest = *ftest
	flags.flagPort = *fport
	flags.flagReload = *freload
	flags.flagVersion = *fversion
	flags.flagHelp = *fhelp
}