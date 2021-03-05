/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package middleware

import (
	"cloudp/utils/settings"
	"encoding/json"
	"errors"
)
import "cloudp/rediscache"
import "cloudp/models/article"

var (
	useorno bool
)

func Cache(name string)  interface{}{
	getflag()
	if useorno{
		hit,e := getFromRedis(name)
		if e !=nil{
			//把内容缓存到cache
			hitdb := getFromDB(name)
			_, _ = rediscache.Set(name, hitdb, settings.Expires)
			return getFromDB(name)
		}
		var hitcache article.Post
		_ = json.Unmarshal(hit, &hitcache)
		//fmt.Println("hitcache")
		return hitcache
	}else {
		//fmt.Println("hitdb")
		return getFromDB(name)
	}
}

func PostCache(p int)  ([]article.Posts, int){
	getflag()
	if useorno{
		_ = rediscache.Setup()
		hit :=  rediscache.Exists("allposts")
		if hit{
			var s []article.Posts
			posts,_ := rediscache.Get("allposts")
			_ = json.Unmarshal(posts, &s)
			var length int = len(s)
			res := pagenation(p,s)
			//击中从缓存读取
			return res,length
		}else {
			posts := article.Getarticles()
			var length int = len(posts)
			res := pagenation(p,posts)
			_, _ = rediscache.Set("allposts", posts, settings.PostsTimeout)
			//未击中缓存更新
			return res,length
		}

	}else {
		//redis未开启
		posts := article.Getarticles()
		res := pagenation(p,posts)
		var length int = len(posts)
		return res,length
	}

}


func CheckCache(name string)  string{
	getflag()
	if useorno{
		_,e := getFromRedis(name)
		if e !=nil{
			return "miss"
		}

		return "hit"
	}else {
		return "fromdb"
	}
}

func getflag()  {
	useorno = settings.UseRedis
}

func getFromDB(name string) article.Post{
	dbFetch := article.Getarticle(name)
	return dbFetch
}

func getFromRedis(name string) ([]byte,error) {
	//init
	initerr := rediscache.Setup()
	if initerr != nil{
		return []byte(""),errors.New("initerr")
	}
	exit := rediscache.Exists(name)
	if exit{
		cacheFetch,err := rediscache.Get(name)
		if err != nil{
			return []byte(""),errors.New("nohit")
		}else {
			return cacheFetch,nil
		}
	}else {
		return []byte(""),errors.New("nohit")
	}

}

func pagenation(p int,data []article.Posts) []article.Posts {
	//分页数据从ini文件里读取
	page := settings.PageSize
	if p<=0{
		return data
	}else {
		data = data[(p-1)*page:(p-1)*page+page]
		return data
	}

}