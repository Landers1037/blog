/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package models

//import (
//	"encoding/json"
//	"io/ioutil"
//	"os"
//)
//
//type Post struct {
//	Title string `json:"title"`
//	Url string `json:"url"`
//	Content string `json:"content"`
//}
//func Getjson() []Post {
//	path,_ := os.Getwd()
//	f,_ := os.Open(path+"/conf/test.json")
//	jf,_ := ioutil.ReadAll(f)
//	var node []Post
//	err := json.Unmarshal([]byte(jf), &node)
//	if err != nil{
//		return []Post{}
//		}
//	return node
//}