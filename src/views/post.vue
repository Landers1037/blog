<template>
    <div class="post">
        <div class="header">
            <div class="animated slideInDown">
                <label id="title" @click="back">{{custom.top_banner}}</label>
            </div>
            <el-divider><span style="font-family: 'DejaVu Sans Mono','Source Code Pro','Liberation Mono',monospace;font-size: 14px">{{custom.top_span}}</span></el-divider>
        </div>
        <div class="title">
            <p>{{title}}</p>
        </div>
        <div id="theme" v-show="theme_control">
            <el-select v-model="theme" placeholder="代码主题" @change="change_theme">
                <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value">
                </el-option>
            </el-select>
        </div>
        <p style="color: #afafaf;font-size: 12px;font-weight: bold">{{date}}</p>
        <div class="wrapper animated fadeIn">
            <div class="markdown-body" v-html="post">

            </div>
        </div>
        <div class="bt-group">
            <el-button type="primary" icon="el-icon-back" size="small" id="prev" @click="toprev">上一篇</el-button>
            <el-button type="primary" id="next" size="small" @click="tonext">下一篇<i class="el-icon-arrow-right el-icon-right"></i></el-button>
        </div>

        <div class="bottom">
            <p><el-icon class="el-icon-lollipop"></el-icon><a style="color: #5f5f5f;font-weight: bold;margin-right: 8px" :href=custom.bottom_url>{{custom.bottom_tag}}</a>
                <el-icon class="el-icon-coffee-cup"></el-icon><a style="color: #5f5f5f;font-weight: bold" href="https://landers1037.github.io">{{custom.bottom_tag2}}</a>
                <br><span style="font-size: 12px;color: #2c3e50">{{custom.bottom_span}}</span></p>
        </div>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    export default {
        name: "post",
        data(){
            return{
                custom: customData,
                url: this.$route.params.url,
                post: null,
                title: null,
                date: null,
                //文章
                prev: "",
                next: "",
                options: [{
                    value: 'github',
                    label: 'github'
                }, {
                    value: 'monokai',
                    label: 'monokai',
                }, {
                    value: 'atom',
                    label: 'atom dark'
                },{
                    value: "solarized",
                    label: "solarized"
                },{
                    value: "xcode",
                    label: "xcode"
                },{
                    value: "xterm",
                    label: "xterm"
                    }
                ],
                theme: "github",
                theme_control: false
            }
        },
        created(){
            let _this = this;
            this.$http.get('/api/article/post',{params:{name:this.url}}).then(res=>{
                let content = res.data.data["content"];
                _this.title = res.data.data["title"];
                _this.date = res.data.data["date"];
                _this.mk(content);
                _this.$nextTick(()=>{
                    let pres = document.getElementsByTagName("pre");
                    for(let i=0;i<pres.length;i++){
                        pres[i].classList.add("hljs");
                    }
                });
            }).catch(err=>{
                _this.$message.error('出现错误了，请求文章失败');
            })
        },
        mounted() {
            let _this = this;
            this.$http.get('/api/article/post',{params:{name:this.url}}).then(res=>{
                let content = res.data.data["content"];
                _this.title = res.data.data["title"];
                _this.date = res.data.data["date"];
                _this.mk(content);
                _this.$nextTick(()=>{
                    this.theme_control = true;
                    let pres = document.getElementsByTagName("pre");
                    for(let i=0;i<pres.length;i++){
                        pres[i].classList.add("hljs");
                    }
                });
            }).catch(err=>{
                this.theme_control = true;
                _this.$message.error('出现错误了，请求文章失败');
            });
            this.loading(1400);
            this.brother();
        },
        methods:{
            back(){
                this.$router.push("/")
            },
            init(url){
                let _this = this;
                this.$http.get('/api/article/post',{params:{name:url}}).then(res=>{
                    let content = res.data.data["content"];
                    _this.title = res.data.data["title"];
                    _this.date = res.data.data["date"];
                    _this.mk(content);
                    _this.$nextTick(()=>{
                        let pres = document.getElementsByTagName("pre");
                        for(let i=0;i<pres.length;i++){
                            pres[i].classList.add("hljs");
                        }
                    });
                }).catch(err=>{
                    _this.$message.error('出现错误了，请求文章失败');
                })
            },
            brother(){
                let _this = this;
                this.$http.get("/api/article/brother",{params:{name:this.url}}).then(res=>{
                    let d = res.data;
                    if(d){
                        _this.prev = d[0];
                        _this.next = d[1];
                    }
                })
            },
            mk(code){
                marked.setOptions({
                    renderer: new marked.Renderer(),
                    highlight: function (c) {
                        return hljs.highlightAuto(c).value;
                    },
                    pendantic: false,
                    gfm: true,
                    tables: true,
                    breaks: true,
                    sanitize: false,
                    smartLists: true,
                    xhtml: false
                });
                this.post = marked(code);
            },
            loading(sec) {
                const loading = this.$loading({
                    lock: true,
                    text: '文章加载中...',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                setTimeout(() => {
                    loading.close();
                }, sec);
            },
            handleScrollTop(){
                window.scrollTo(0, 0);
            },
            toprev(){
                if(this.prev!==""){
                    this.$router.push("/p/"+this.prev);
                    this.init(this.prev);
                    this.handleScrollTop();
                    this.loading(1000);
                    this.url = this.prev;
                    this.brother();
                }
            },
            tonext(){
                if(this.next!==""){
                    this.$router.push("/p/"+this.next);
                    this.init(this.next);
                    this.handleScrollTop();
                    this.loading(1000);
                    this.url = this.next;
                    this.brother();
                }
            },
            change_theme(){
                // 每次更换前都移除上一次的样式
                let head = document.getElementsByTagName('head')[0];
                let linkTag = document.createElement('link');
                switch (this.theme) {
                    case "github":
                        console.log("使用默认主题 github");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/github.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    case "monokai":
                        console.log("使用主题 monokai");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/monokai-sublime.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    case "atom":
                        console.log("使用主题 atom one-dark");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/atom-one-dark.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    case "solarized":
                        console.log("使用主题 solarized dark");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/solarized-dark.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    case "xcode":
                        console.log("使用主题 xcode");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/xcode.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    case "xterm":
                        console.log("使用主题 xterm");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/xt256.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                    default:
                        console.log("使用默认主题 github");
                        document.getElementsByTagName('head').item(0).removeChild(document.getElementById('dynamic-theme'));
                        linkTag.id = 'dynamic-theme';
                        linkTag.href = "https://cdn.staticfile.org/highlight.js/10.0.0/styles/github.min.css";
                        linkTag.setAttribute('rel','stylesheet');
                        linkTag.setAttribute('media','all');
                        linkTag.setAttribute('type','text/css');

                        head.appendChild(linkTag);
                        break;
                }
            }
        }

    }
</script>

<style scoped>
    #title{
        background-color: #363636;
        color: white;
        padding: 8px 10px;
        cursor: pointer;
        font-family: mo,monospace;
    }
    .title{
        margin: 20px auto 10px;
        width: fit-content;
        padding: 4px 8px;
        background-color: white;
        border-radius: 4px;
        border-style: dashed;
    }
    #theme{
        position: absolute;
        z-index: 999;
        right: 20px;
        top: 20px;
        max-width: 120px;
        transition: all .3s ease;
    }
    #theme /deep/ .el-select .el-input.is-focus .el-input__inner{
        border-color: #DCDFE6;
    }
    #theme /deep/ .el-input.is-active .el-input__inner, .el-input__inner:focus{
        border-color: #DCDFE6;
    }
    #theme /deep/ .el-select .el-input__inner:focus{
        border-color: #DCDFE6;
    }
    @media (max-width: 640px) {
        #theme{
            right: 8px;
            top: 34px;
            max-width: 80px;
            transition: all .3s ease;
        }
        #theme /deep/ .el-input {
            font-size: 12px;
        }
        #theme /deep/ .el-input--suffix .el-input__inner{
            padding-right: 16px;
        }
        #theme /deep/ .el-input__inner{
            height: 32px;
            line-height: 32px;
            padding: 0 0 0 8px;
        }
        #theme /deep/ .el-input__icon{
            line-height: 32px;
            width: 15px;
        }
    }
    .wrapper{
        text-align: left;
        box-shadow: 1px 1px 10px #afafaf;
        border-radius: 10px;
        max-width: 980px;
        margin: 15px auto 0;
    }
    .markdown-body{
        box-sizing: border-box;
        min-width: 200px;
        max-width: 980px;
        margin: 0 auto;
        padding: 25px;
    }
    .bottom{
        margin-top: 20px;
        font-family: "DejaVu Sans Mono","Segoe UI",Monaco,monospace;
    }
    .bt-group{
        position: relative;
        max-width: 980px;
        margin: 15px auto 0;
        padding-bottom: 30px;
    }
    .bt-group #prev{
        position: absolute;
        left: 10px;
    }
    .bt-group #next{
        position: absolute;
        right: 10px;
    }
    @media (max-width: 767px) {
        .markdown-body {
            padding: 15px;
        }
    }
</style>
<style>
    .markdown-body p code{
        background-color: #8d8cff;
        color: #ffffff;
    }
    .markdown-body .highlight pre, .markdown-body pre{
        border-radius: 8px;
    }
</style>