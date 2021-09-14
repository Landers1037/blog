/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package crdb

import (
	"blog/models/admin"
	"blog/models/article"
	"blog/models/message"
	"blog/models/subscribe"
)

// Tables 用于驱动的结构体字典
var Tables =  map[string]interface{}{
	"table_admin": admin.DB_BLOG_ADMIN{},
	"table_post": article.DB_BLOG_POST{},
	"table_cate": article.DB_BLOG_CATES{},
	"table_tag": article.DB_BLOG_TAGS{},
	"table_share": article.DB_BLOG_SHARE{},
	"table_like": article.DB_BLOG_LIKES{},
	"table_view": article.DB_BLOG_VIEWS{},
	"table_comment": article.DB_BLOG_COMMENTS{},
	"table_zhuanlan": article.DB_BLOG_ZHUANLAN{},
	"table_message": message.DB_BLOG_MESSAGES{},
	"table_sub": subscribe.DB_BLOG_SUBSCRIBE{},
}
