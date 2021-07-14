// 可自由定制的项目
// 在打包的时候使用
let prefix = 'https:' === document.location.protocol?'https':'http';
const customData = {
    // 前端分离时使用
    // api_url: prefix + "://127.0.0.1:5000",
    api_url: "",
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
    dashboard_count: 1,
    // default theme support github|monokai|atom|xterm|solarized|xcode
    default_theme: "xt256"
}

export default customData;