/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package routers

import (
	"blog/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type staticRouter struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Alias string `json:"alias"`
}

// 代理静态文件
// 只有在标志启用的时候生效
func addStaticRouter(r *gin.Engine) {
	if config.Cfg.StaticRouter {
		var router []staticRouter
		data, e := ioutil.ReadFile("conf/router.json")
		if e != nil {
			return
		}else {
			_ = json.Unmarshal(data, &router)
			var fileGroup [][]string
			var DirGroup [][]string
			var TryGroup [][]string
			for _, sp := range router {
				if sp.Type == "file" {
					fileGroup = append(fileGroup, []string{sp.Path, sp.Alias})
				}else if sp.Type == "dir" {
					DirGroup = append(DirGroup, []string{sp.Path, sp.Alias})
				}else if sp.Type == "try" {
					TryGroup = append(TryGroup, []string{sp.Path, sp.Alias})
				}
				// try_file机制应该写在json定义内部
			}
			for _, v := range fileGroup {
				r.StaticFile(v[0], v[1])
			}
			for _ ,v := range DirGroup {
				r.StaticFS(v[0], http.Dir(v[1]))
			}
			for _, v := range TryGroup {
				r.GET(v[0], func(c *gin.Context) {
					c.File(v[1])
				})
			}

		}
	}
}
