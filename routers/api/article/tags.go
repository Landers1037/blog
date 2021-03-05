/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article
/*
和文章相关的api接口
主要实现
文章标签的编辑

 */

import (
	"cloudp/models/article"
	"cloudp/utils/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Gettags(c *gin.Context)  {
	//标签获取
	//因为标签是从markdown文件里提取的所以只是需要把标签提取出来
	//前期并没有数据库的支持所以只是测试
	name := c.Query("name")
	maps := make(map[string]interface{})
	if name != ""{
		maps["name"] = name
	}
	//其他的query参数
	code := err.SUCCESS
	data := article.Gettags()
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": err.GetMsg(code),
		"data": data,
	})
}


