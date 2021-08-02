# CHANGELOG 更新记录

## v1.1

初始版本 基于gin-example的blog示例编写

## v3.2

- 改用`marked.js`作为前端md的渲染器， 抛弃使用数据库直接存取渲染html的做法
- 支持了linux的signal特性来实现服务的守护进程
- 支持多实例的集群部署

## v3.3

- 新增了数据库表 支持blog的订阅 评论 点赞 分享
- 修改了原先的统计量刷新逻辑，改为后台全局定时器协程方式实现
- 修复了因为可能的空refer或者host导致的空指针异常

- **新增特性**： 静态文件托管

在`app.ini`配置中启动此特性

```ini
[server]

STATIC_ROUTER = 1

[middle]

TRY_FILE = 1
TRY_FILE_INDEX = dist/index.html
```

开启后访问/会自动托管到`TRY_FILE_INDEX`中定义的前端文件上

适用于不使用前后端分离而是直接使用本服务渲染前端文件的场景

## v4.2 offline version

新增了分支`br_offline`包含了只使用离线构建版本的js依赖

适用于无网络情况下，资源不通过cdn引入而是直接编译在静态文件中

## v4.2 upx

- 使用upx打包本服务
- 基于go1.16.2编译

## v5.1

- 全新设计 基于go1.16.6编译 使用go mod方式管理项目

- 新增编译脚本和docker构建脚本， 为适配环境全部使用-static静态链方式编译。基于gin的编译加入优化参数使用go_json作为json序列化库

- 修复了前端空tag时仍然显示的问题

- 修复了tag数组为空时使用length判断的报错问题

- 新增了控制面板的**新增文章**按钮， **导出文章**按钮

- docker镜像v5发布 

  ```bash
  docker pull landers1037/blog:v5
  ```

  

## v5.4

- 修正评论表数据库字段属性错误

- 正式支持评论功能，前端新增评论区

- 新增两个前端自定义项

  ```json
      message_duration: 1500,
      loading_duration: 1000
  ```

- 新增控制面板评论管理页面

- 优化前端显示样式，更换avatar资源

## v5.6

- 增加文章的评论 点赞 分享功能
- 修复了评论偶现的无法显示问题
- 修复了点赞无法刷新的问题
- 添加了额外的标题锚点样式 更加耐看
- 增加了评论，点赞和分享的路由处理逻辑

## V6.0

- 修正引入错误

- 完整的命令行支持

    现在使用`./blog`即可查看命令行参数

    ```bash
    NAME:
       blog - powerful markdown-based blog
    
    USAGE:
       blog [global options] command [command options] [arguments...]
    
    VERSION:
       v6.0
    
    DESCRIPTION:
       一个基于markdown文档的动态博客部署工具
    
    AUTHORS:
       Landers <liaorenj@gmail.com>
       wxk <xk_wang@qq.com>
    
    COMMANDS:
       help, h  Shows a list of commands or help for one command
       App configs:
         config, conf, c  应用配置
       Run a web service:
         web, w, serve, server  启动blog的web服务
       Service manager:
         service, s  服务管理
       Tools of blog:
         tool, t  博客配套工具
    
    GLOBAL OPTIONS:
       --help, -h     show help (default: false)
       --version, -v  print the version (default: false)
    
    COPYRIGHT:
       renj.io 2021.
    ```

    

## v6.1

- 支持docker compose部署
- 添加了404页面
- 增加了`.env`的方式自定义head

## v6.2

- 支持动态的标题title适配
- 更新关于页面

## v6.3

- 后台管理新增了点赞和分享接口

- 添加了makefile构建方式，可以直接使用`make build`构建了

  ```bash
  # make help
  this Makefile help you build blog binary
  Usage: 
  
     build   build the blog application
     run     run go run app.go with -race
     clean   clean the go binary
     setup   setup go modules
     env     setup envs
     help    show help usages
  ```

  

## v6.4

- 优化了首页加载逻辑，添加加载骨架
- 优化了全局圆角样式和阴影样式
- 优化了留言页面
- 优化了首页按钮点击体验
- 修复了部分页面的代码高亮样式无法显示问题
- 新增了文章列表的标签显示
- 重新设计了专栏
- 关于页面添加感谢列表 iconfont

## V6.5

- 后台完整的支持评论 点赞的管理
- 留言支持更新
- 对于所有长文本支持以`textarea`的方式显示内容
- 去除了无效的调试代码
- 调整数据表格宽度，现后台页面不再支持移动端

## V6.6

- 修复专栏的分栏问题
- 文章列表支持显示日期
- 更新了总览图
- 后台程序添加路由日志打印

## v6.7

- 修复评论渲染高亮问题
- 全局增加可显示标签，增加文章时间

## v6.8

- 修复同标签页面的跳转问题
- 优化了文章加载的loading页面显示逻辑
- 配置文件路径和数据库路径支持多层级创建

## v6.9

- 模板保存位置修改为用户根目录例如`linux: ~`

## v6.10

- 添加文章浏览量统计显示
- 评论区提示语优化