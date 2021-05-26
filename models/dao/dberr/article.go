/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dberr

import (
	"errors"
)

var (
	ERR_ARTICLE_EXIST = errors.New("article already exist")
	ERR_EMPTY_POSTS = errors.New("zhuanlan posts empty")
)

