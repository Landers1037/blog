/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: cloudp
*/
package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	Cluster bool

	PageSize int
	JwtSecret string
	Uv_flag bool
	SimpleAuth_flag bool
	UseRedis bool

) //一次定义多个变量

var (
	RedisHost string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout  time.Duration
	Expires int
	PostsTimeout int
)
func init()  {
	var err error
	Cfg,err = ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}//出现错误时执行
	loadMode()
	loadServer()
	loadApp()
	loadMiddle()
	RedisSetting()
}

func loadMode()  {
	//决定debug模式的开关
	RunMode = Cfg.Section("mode").Key("RUN_MODE").MustString("release")
}

func loadServer()  {
	//服务器相关配置
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	HTTPPort = server.Key("HTTP_PORT").MustInt(6000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	Cluster = server.Key("CLUSTER").MustBool(false)
}

func loadApp() {
	//应用的配置
	app, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = app.Key("PAGE_SIZE").MustInt(8)
}

func loadMiddle()  {
	//中间件
	mid  ,_:= Cfg.GetSection("middle")
	Uv_flag = mid.Key("UV").MustBool(true)
	SimpleAuth_flag = mid.Key("SIMPLEAUTH").MustBool(true)
	UseRedis = mid.Key("USEREDIS").MustBool(true)
}

func RedisSetting()  {
	//是否使用redis缓存
	redis := Cfg.Section("redis")
	RedisHost = redis.Key("Host").MustString("127.0.0.1:6379")
	Password = redis.Key("Password").MustString("")
	MaxIdle = redis.Key("MaxIdle").MustInt(30)
	MaxActive = redis.Key("MaxActive").MustInt(30)
	IdleTimeout  = time.Duration(redis.Key("IdleTimeout").MustDuration(200)) * time.Second
	Expires = redis.Key("EXPIRES").MustInt(600)
	PostsTimeout = redis.Key("POSTSTIMEOUT").MustInt(60)
}