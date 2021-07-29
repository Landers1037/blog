/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/
package settings

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"sync"
	"time"
)

const (
	APP_NAME = "blog"
	APP_FILE = "conf/app.ini"
	APP_PORT = 5000
	APP_LOG = "data/blog.log"
	APP_LOG_LEVEL = "error"
	APP_DB = "data/blog.db"
	APP_PID = "data/blog.pid"
	APP_DEFAULT_ADMIN = "admin"
	APP_DEFAULT_PASSWD = "12345"
)

type Cfg struct {
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Cluster      bool
	StaticRouter bool
	HideDBLog    bool

	PageSize 			int
	MessageSize 		int
	SortPostBy 			string
	SortPostReverse 	bool
	SortMessageBy 		string
	SortMessageReverse 	bool
	SortCommentBy 		string
	SortCommentReverse 	bool
	UseContentAsAbs 	bool
	MaxContentLength 	int
	CustomEmptyAbs 		string
	FakeStaticUrl 		bool
	ZhuanLanId		 	bool
	JwtSecret 			string
	AppRefer 			string
	AppHost  			string
	AppDomain 			string
	AllowIE 			bool

	UvFlag 			bool
	SimpleAuthFlag 	bool
	CORSFlag 		bool
	PostViewFlag 	bool
	UseRedis 		bool
	TryFile 		bool
	TryFileIndex 	string

	AdminName 	 string
	AdminPwd  	 string
	CookieMaxAge int
	StopAdmin 	 bool

	DB string
	MySQLType string
	MySQLHost string
	MySQLUserName 	string
	MySQLPassWd   	string
	MySQLName     	string
	MySQLTablePrefix string

	RedisHost 		string
	Password 		string
	MaxIdle 		int
	MaxActive 		int
	IdleTimeout 	time.Duration
	Expires 		int
	PostsTimeout 	int

	AppName 		string
	APPPid 			string
	APPLog 			string
	APPLogLevel 	string
	APPLogFile 		string
	APPLogEnable 	bool
}

// 定义ini的映射结构体
type iniCfg struct {
	Mode Mode `ini:"mode"`
	Run Run `ini:"run"`
	App App `ini:"app"`
	Server Server `ini:"server"`
	Admin Admin `ini:"admin"`
	Mysql Mysql `ini:"mysql"`
	Sqlite Sqlite `ini:"sqlite"`
	Middle Middle `ini:"middle"`
	Redis Redis `ini:"redis"`
}

type Mode struct {
	RunMode string `ini:"RUN_MODE"`
}

type Run struct {
	AppName string `ini:"APP_NAME"`
	APPPid string `ini:"APP_PID"`
	APPLog string `ini:"APP_LOG"`
	APPLogLevel string `ini:"APP_LOG_LEVEL"`
	APPLogFile string `ini:"APP_LOG_FILE"`
	APPLogEnable bool `ini:"APP_LOG_ENABLE"`
}

type App struct {
	PageSize int `ini:"PAGE_SIZE"`
	MessageSize int `ini:"MESSAGE_SIZE"`
	SortPostBy string `ini:"SORT_POST_BY"`
	SortPostReverse bool `ini:"SORT_POST_REVERSE"`
	SortMessageBy string `ini:"SORT_MESSAGE_BY"`
	SortMessageReverse bool `ini:"SORT_MESSAGE_REVERSE"`
	SortCommentBy string `ini:"SORT_COMMENT_BY"`
	SortCommentReverse bool `ini:"SORT_COMMENT_REVERSE"`
	UseContentAsAbs bool `ini:"USE_CONTENT_AS_ABS"`
	MaxContentLength int `ini:"MAX_CONTENT_LENGTH"`
	CustomEmptyAbs string `ini:"CUSTOM_EMPTY_ABS"`
	FakeStaticUrl bool `ini:"FAKE_STATIC_URL"`
	ZhuanLanId bool `ini:"ZHUANLAN_ID"`
	JwtSecret string `ini:"JWT_SECRET"`
	AppRefer string `ini:"APP_REFER"`
	AppHost  string `ini:"APP_HOST"`
	AppDomain string `ini:"APP_DOMAIN"`
	AllowIE bool `ini:"APP_ALLOW_IE"`
}

type Server struct {
	HTTPPort int `ini:"HTTP_PORT"`
	ReadTimeout time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
	Cluster bool `ini:"CLUSTER"`
	StaticRouter bool `ini:"STATIC_ROUTER"`
	HideDBLog bool `ini:"HIDE_DB_LOG"`
}

type Admin struct {
	AdminName string `ini:"USERNAME"`
	AdminPwd  string `ini:"PASSWD"`
	CookieMaxAge int `ini:"PASSWD"`
	StopAdmin bool `ini:"STOP_ADMIN"`
}

type Mysql struct {
	MySQLType string `ini:"TYPE"`
	MySQLHost string `ini:"USER"`
	MySQLUserName string `ini:"PASSWORD"`
	MySQLPassWd   string `ini:"HOST"`
	MySQLName     string `ini:"NAME"`
	MySQLTablePrefix string `ini:"TABLE_PREFIX"`
}

type Sqlite struct {
	DB string `ini:"DB"`
}

type Middle struct {
	UvFlag 			bool `ini:"UV"`
	SimpleAuthFlag 	bool `ini:"SIMPLEAUTH"`
	CORSFlag 		bool `ini:"CORS"`
	PostViewFlag 	bool `ini:"POSTVIEW"`
	UseRedis 		bool `ini:"USEREDIS"`
	TryFile 		bool `ini:"TRY_FILE"`
	TryFileIndex 	string `ini:"TRY_FILE_INDEX"`
}

type Redis struct {
	RedisHost string `ini:"Host"`
	Password string `ini:"Password"`
	MaxIdle int `ini:"MaxIdle"`
	MaxActive int `ini:"MaxActive"`
	IdleTimeout time.Duration `ini:"IdleTimeout"`
	Expires int `ini:"EXPIRES"`
	PostsTimeout int `ini:"POSTSTIMEOUT"`
}

// 防止重复引用 使用初始化加载方式
// err分为读取失败 初始化失败
// p不为空时优先判断是否存在 不存在会直接退出 这个操作不会触发配置的自动生成
func InitCfg(p string) (Cfg, error) {
	var config Cfg
	var confFile string

	if p == "" {
		confFile = APP_FILE
	}else if _, e := os.Stat(p);os.IsNotExist(e) {
		log.Fatalf("Fail to parse '%s', not exist\n", p)
	}else {
		confFile = p
	}

	// 开始加载配置 此时失败会自动创建
	if _, e := os.Stat(confFile);os.IsNotExist(e) {
		return Cfg{}, e
	}
	cfg, err := ini.Load(confFile)
	if err != nil {
		log.Fatalf("Fail to parse '%s': %s\n", APP_FILE, err.Error())
		return config, err
	}//出现错误时执行
	loadMode(&config, cfg)
	loadServer(&config, cfg)
	loadApp(&config, cfg)
	loadMiddle(&config, cfg)
	RedisSetting(&config, cfg)
	loadDatabase(&config, cfg)
	loadRunningFile(&config, cfg)
	loadAdmin(&config, cfg)

	return config, nil
}

// 配置序列化
func SaveConfig(f string) error {
	var cf string
	if f == "" {
		cf = APP_FILE
	}else {
		cf = f
	}
	lock := sync.Mutex{}
	if _, e := os.Stat(cf); os.IsNotExist(e) {
		log.Printf("%s不存在 即将初始化\n", cf)
		cfg := ini.Empty()
		defaultCfg := &iniCfg{
			Mode:   Mode{RunMode: "release"},
			Run:    Run{
				AppName:      APP_NAME,
				APPPid:       APP_PID,
				APPLog:       APP_LOG,
				APPLogLevel:  APP_LOG_LEVEL,
				APPLogFile:   APP_LOG,
				APPLogEnable: true,
			},
			App:    App{
				PageSize:           8,
				MessageSize:        5,
				SortPostBy:         "id",
				SortPostReverse:    true,
				SortMessageBy:      "id",
				SortMessageReverse: true,
				SortCommentBy:      "id",
				SortCommentReverse: true,
				UseContentAsAbs:    true,
				MaxContentLength:   120,
				CustomEmptyAbs:     "<code>Sorry</code>该文章暂无概述",
				FakeStaticUrl:      false,
				ZhuanLanId:         true,
				JwtSecret:          "12345",
				AppRefer:           "",
				AppHost:            "127.0.0.1",
				AppDomain:          "",
				AllowIE:            false,
			},
			Server: Server{
				HTTPPort:     5000,
				ReadTimeout:  60,
				WriteTimeout: 60,
				Cluster:      false,
				StaticRouter: false,
				HideDBLog:    true,
			},
			Admin:  Admin{
				AdminName:    APP_DEFAULT_ADMIN,
				AdminPwd:     APP_DEFAULT_PASSWD,
				CookieMaxAge: 3600,
				StopAdmin:    false,
			},
			Mysql:  Mysql{
				MySQLType:        "mysql",
				MySQLHost:        "127.0.0.1",
				MySQLUserName:    "root",
				MySQLPassWd:      "123456",
				MySQLName:        "blog",
				MySQLTablePrefix: "blog_",
			},
			Sqlite: Sqlite{DB: APP_DB},
			Middle: Middle{
				UvFlag:         true,
				SimpleAuthFlag: false,
				CORSFlag:       true,
				PostViewFlag:   true,
				UseRedis:       false,
				TryFile:        false,
				TryFileIndex:   "",
			},
			Redis:  Redis{
				RedisHost:    "127.0.0.1:6379",
				Password:     "",
				MaxIdle:      30,
				MaxActive:    30,
				IdleTimeout:  200,
				Expires:      60,
				PostsTimeout: 10,
			},
		}
		err := cfg.ReflectFrom(defaultCfg)
		if err != nil {
			return err
		}
		lock.Lock()
		if _, errConf := os.Stat("conf");os.IsNotExist(errConf) && cf == APP_FILE {
			e = os.Mkdir("conf", 0644)
			if e != nil {
				return e
			}
		}

		e = cfg.SaveTo(cf)
		defer lock.Unlock()
		return e
	}
	return nil
}

func loadMode(config *Cfg, c *ini.File)  {
	//决定debug模式的开关
	config.RunMode = c.Section("mode").Key("RUN_MODE").MustString("release")
}

func loadServer(config *Cfg, c *ini.File)  {
	//服务器相关配置
	server, err := c.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %s\n", err.Error())
	}
	config.HTTPPort = server.Key("HTTP_PORT").MustInt(6000)
	config.ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	config.WriteTimeout =  time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	config.Cluster = server.Key("CLUSTER").MustBool(false)
	config.StaticRouter = server.Key("STATIC_ROUTER").MustBool(false)
	config.HideDBLog = server.Key("HIDE_DB_LOG").MustBool(false)
}

func loadApp(config *Cfg, c *ini.File) {
	//应用的配置
	app, err := c.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %s\n", err.Error())
	}

	config.JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	config.PageSize = app.Key("PAGE_SIZE").MustInt(8)
	config.MessageSize = app.Key("MESSAGE_SIZE").MustInt(5)
	config.SortPostBy = app.Key("SORT_POST_BY").MustString("id")
	config.SortPostReverse = app.Key("SORT_POST_REVERSE").MustBool(true)
	config.SortMessageBy = app.Key("SORT_MESSAGE_BY").MustString("id")
	config.SortMessageReverse = app.Key("SORT_MESSAGE_REVERSE").MustBool(true)
	config.SortCommentBy = app.Key("SORT_COMMENT_BY").MustString("id")
	config.SortCommentReverse = app.Key("SORT_COMMENT_REVERSE").MustBool(true)

	config.UseContentAsAbs = app.Key("USE_CONTENT_AS_ABS").MustBool(false)
	config.MaxContentLength = app.Key("MAX_CONTENT_LENGTH").MustInt(30)
	config.CustomEmptyAbs = app.Key("CUSTOM_EMPTY_ABS").MustString("")
	config.FakeStaticUrl = app.Key("FAKE_STATIC_URL").MustBool(false)
	config.ZhuanLanId = app.Key("ZHUANLAN_ID").MustBool(true)
	config.AppDomain = app.Key("APP_DOMAIN").MustString("")
	config.AppRefer = app.Key("APP_REFER").MustString("")
	config.AppHost = app.Key("APP_REFER").MustString("")
	config.AllowIE = app.Key("APP_ALLOW_IE").MustBool(true)
}

func loadMiddle(config *Cfg, c *ini.File)  {
	//中间件
	mid  ,_:= c.GetSection("middle")
	config.UvFlag = mid.Key("UV").MustBool(true)
	config.SimpleAuthFlag = mid.Key("SIMPLEAUTH").MustBool(true)
	config.UseRedis = mid.Key("USEREDIS").MustBool(true)
	config.CORSFlag = mid.Key("CORS").MustBool(true)
	config.PostViewFlag = mid.Key("POSTVIEW").MustBool(true)
	config.TryFile = mid.Key("TRY_FILE").MustBool(false)
	config.TryFileIndex = mid.Key("TRY_FILE_INDEX").MustString("")
}

func RedisSetting(config *Cfg, c *ini.File)  {
	//是否使用redis缓存
	redis := c.Section("redis")
	config.RedisHost = redis.Key("Host").MustString("127.0.0.1:6379")
	config.Password = redis.Key("Password").MustString("")
	config.MaxIdle = redis.Key("MaxIdle").MustInt(30)
	config.MaxActive = redis.Key("MaxActive").MustInt(30)
	config.IdleTimeout  = time.Duration(redis.Key("IdleTimeout").MustDuration(200)) * time.Second
	config.Expires = redis.Key("EXPIRES").MustInt(600)
	config.PostsTimeout = redis.Key("POSTSTIMEOUT").MustInt(60)
}

func loadDatabase(config *Cfg, c *ini.File) {
	// sqlite3
	sqlite := c.Section("sqlite")
	config.DB = sqlite.Key("DB").MustString("")
	// mysql
	mysql := c.Section("mysql")
	config.MySQLType = mysql.Key("TYPE").MustString("mysql")
	config.MySQLHost = mysql.Key("HOST").MustString("")
	config.MySQLName = mysql.Key("NAME").MustString("")
	config.MySQLUserName = mysql.Key("USER").MustString("")
	config.MySQLPassWd = mysql.Key("PASSWORD").MustString("")
	config.MySQLTablePrefix = mysql.Key("TABLE_PREFIX").MustString("")
}

func loadRunningFile(config *Cfg, c *ini.File) {
	running := c.Section("run")
	config.AppName = running.Key("APP_NAME").MustString("blog")
	config.APPPid = running.Key("APP_PID").MustString("blog.pid")
	config.APPLog = running.Key("APP_LOG").MustString("blog.log")
	config.APPLogLevel = running.Key("APP_LOG_LEVEL").MustString("error")
	config.APPLogFile = running.Key("APP_LOG_FILE").MustString("app.log")
	config.APPLogEnable = running.Key("APP_LOG_ENABLE").MustBool(false)
}

func loadAdmin(config *Cfg, c *ini.File) {
	admin := c.Section("admin")
	config.AdminName = admin.Key("USERNAME").MustString("")
	config.AdminPwd = admin.Key("PASSWD").MustString("")
	config.CookieMaxAge = admin.Key("COOKIE_MAX_AGE").MustInt(1800)
	config.StopAdmin = admin.Key("STOP_ADMIN").MustBool(false)
}