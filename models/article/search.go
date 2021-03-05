/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article

import (
	"cloudp/models"
)

//这里的Post全局变量在article里定义了

func Search(name string) (s []Posts){
	var pattern string = "%"+name+"%"
	models.THISdb.Where("title LIKE ?",pattern).Find(&s)

	return
}