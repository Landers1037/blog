/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/

package err

//定义错误的响应码
//默认以233开头
const (
	SUCCESS        = 233200
	ERROR          = 233500
	INVALID_PARAMS = 233400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)
