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
	"time"
)

type Cfg struct {
	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	Cluster bool

	PageSize int
	MessageSize int
	JwtSecret string
	AppRefer string
	AppHost  string
	AppDomain string
	AllowIE   bool

	Uv_flag bool
	SimpleAuth_flag bool
	CORS_flag bool
	PostView_flag bool
	UseRedis bool

	DB string
	MySQLType string
	MySQLHost string
	MySQLUserName string
	MySQLPassWd   string
	MySQLName     string
	MySQLTablePrefix string

	RedisHost string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout  time.Duration
	Expires int
	PostsTimeout int
}

// 防止重复引用 使用初始化加载方式
func InitCfg() (Cfg, error) {
	var config Cfg
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
		return config, err
	}//出现错误时执行
	loadMode(&config, cfg)
	loadServer(&config, cfg)
	loadApp(&config, cfg)
	loadMiddle(&config, cfg)
	RedisSetting(&config, cfg)
	loadDatabase(&config, cfg)

	return config, nil
}

func loadMode(config *Cfg, c *ini.File)  {
	//决定debug模式的开关
	config.RunMode = c.Section("mode").Key("RUN_MODE").MustString("release")
}

func loadServer(config *Cfg, c *ini.File)  {
	//服务器相关配置
	server, err := c.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	config.HTTPPort = server.Key("HTTP_PORT").MustInt(6000)
	config.ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	config.WriteTimeout =  time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	config.Cluster = server.Key("CLUSTER").MustBool(false)
}

func loadApp(config *Cfg, c *ini.File) {
	//应用的配置
	app, err := c.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	config.JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	config.PageSize = app.Key("PAGE_SIZE").MustInt(8)
	config.MessageSize = app.Key("MESSAGE_SIZE").MustInt(5)
	config.AppDomain = app.Key("APP_DOMAIN").MustString("")
	config.AppRefer = app.Key("APP_REFER").MustString("")
	config.AppHost = app.Key("APP_REFER").MustString("")
	config.AllowIE = app.Key("APP_ALLOW_IE").MustBool(true)
}

func loadMiddle(config *Cfg, c *ini.File)  {
	//中间件
	mid  ,_:= c.GetSection("middle")
	config.Uv_flag = mid.Key("UV").MustBool(true)
	config.SimpleAuth_flag = mid.Key("SIMPLEAUTH").MustBool(true)
	config.UseRedis = mid.Key("USEREDIS").MustBool(true)
	config.CORS_flag = mid.Key("CORS").MustBool(true)
	config.PostView_flag = mid.Key("POSTVIEW").MustBool(true)
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