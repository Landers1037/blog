# Blog Front

博客服务的前端配置指导文档

博客前端服务由vue技术栈编写 支持界面的定制化

## 使用方式

需要定义请clone本项目的front分支 进行修改编译

```bash
git clone https://github.com/Landers1037/blog.git -b br_front
```

## 定制属于自己的服务

### 定制页面

**custom.js**

```javascript
// 可自由定制的项目
// 在打包的时候使用
const customData = {
    api_url: "http://127.0.0.1:5000",
    author: "Landers",
    top_banner: "Landers1037",
    top_span: "Never Stop Debugging",
    site_name: "renj.io",
    site_url: "http://renj.io",
    github: "landers1037",
    project: "mgek.cc",
    project_des: "Mgek项目记录生活中的灵感",
    project_url: "http://mgek.cc",
    bottom_tag: "By Landers",
    bottom_url: "http://renj.io",
    bottom_tag2 : "Github",
    bottom_url2: "https://landers1037.github.io",
    bottom_span: "Golang & Vue",
    email: "mail@renj.io",
    start_year: "2017",
    start_date: "2017/7/1",
    dashboard_count: 5
}
```

参数解析

`api_url` 需要请求的后端接口地址 正式发布时例如`https://api.xxx.com`

`author` 作者名称 显示在页面右侧栏的作者名 写上你的昵称

`top_banner` 顶部一级标题

`top_span` 顶部的自定义标语

`site_name` 右边栏可以自定义的你的网站名称

`site_url` 右边栏可以自定义的你的网站地址 点击后可以跳转

`github` 你的githb名称 只需要你的账户名称

`project` 右边栏可以定制的你的项目名称

`project_url` 右边栏可以定制的你的项目url 可以填你的github **不填写则不会跳转**

`project_des` 右边栏关于你的项目的描述

`bottom_tag`底部标签 左标签名称

`bottom_url` 底部标签 左标签的url为空时则忽略 否则可以点击跳转

`bottom_tag2` 底部标签 右标签的名称

`bottom_url2` 底部标签 右标签的url为空时则忽略 否则可以点击跳转

`bottom_span` 底部子标签 自定义内部语句支持emoji **为空时则不显示**

`email` 你的电子邮箱

`start_year`  显示在状态弹框中的开始年份 请在这里填上开始写博客的年份

`start_date` 用于计算具体的距离开始写博客的时间 请填写如下格式`YYYY/mm/dd`

`dashboard_count` 设置点击头像后经过多少次计数才会进入后台控制页面，因为本博客不提供显式的直接登陆页面，所以在经过正确点击次数后才会显示管理员的登陆页面

### 定制样式

**custom.css**

在此文件中定制你的样式 注意这里定义的样式时全局生效的

```css
// 博客文章以及摘要中出现的code代码块高亮颜色
.markdown-body p code{
    background-color: #808080;
    color: #ffffff;
}

// 主页的博客标题颜色
.post .post-a{
    color: #409eff;
}

// 主页的博客标题在鼠标放置时变化的颜色
.post .post-a:hover{
    color: #2f343f;
}

// 主页的博客摘要中摘要文字的颜色
.post .abstract{
    color: #3f3f3f;
}

// 右边栏中的标签主题颜色
.el-tag--dark.el-tag--success{
    background-color: #6aaaaa;
    border-color: #6aaaaa;
}
```

值得注意的是：

在此文件中默认预置的主题外 你还可以通过f12查看页面源码自由定制页面

在此文件中写的所有css都会覆盖之前的样式