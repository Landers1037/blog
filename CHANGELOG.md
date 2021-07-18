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