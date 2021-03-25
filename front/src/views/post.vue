<template>
    <div class="post">
    <top_banner></top_banner>
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
            <div class="markdown-body" v-html="post" :class="theme">

            </div>
        </div>
        <div class="bt-group">
            <el-button type="primary" icon="el-icon-back" size="small" id="prev" @click="toprev">上一篇</el-button>
            <el-button type="primary" id="next" size="small" @click="tonext">下一篇<i class="el-icon-arrow-right el-icon-right"></i></el-button>
        </div>
        <bottom_banner></bottom_banner>
    </div>
</template>

<script>
    import Top_banner from "../components/top_banner";
    import Bottom_banner from "../components/bottom_banner";
    import customData from "../custom/custom";
    import api_article from "../api/article";
    import {get_code_theme, set_code_theme} from "../store/store";
    import marked from 'marked';
    import hljs from 'highlight.js';
    export default {
        name: "post",
        components: {Top_banner, Bottom_banner},
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
                },{
                    value: "dracula",
                    label: "dracula"
                }
                ],
                theme: "github",
                theme_control: false
            }
        },
        created(){
            let _this = this;
            _this.init_theme();
            this.$http.get(api_article.api_article_more,{params:{name:this.url}}).then(res=>{
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
            this.$http.get(api_article.api_article_more,{params:{name:this.url}}).then(res=>{
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
                this.$http.get(api_article.api_article_more,{params:{name:url}}).then(res=>{
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
                this.$http.get(api_article.api_article_brother,{params:{name:this.url}}).then(res=>{
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
            init_theme(){
                // 存在主题配置时使用配置
                this.theme = get_code_theme();
                this.change_theme();
            },
            change_theme(){
                // 每次更换前都移除上一次的样式
                let linkTag = document.getElementById("dynamic-theme");
                let prefix = "/static/styles/"
                switch (this.theme) {
                    case "github":
                        console.log("使用默认主题 github");
                        set_code_theme("github");
                        linkTag.setAttribute("href", prefix + "github.css");
                        break;
                    case "monokai":
                        console.log("使用主题 monokai");
                        set_code_theme("monokai");
                        linkTag.setAttribute("href", prefix + "monokai.css");
                        break;
                    case "atom":
                        console.log("使用主题 atom one-dark");
                        set_code_theme("atom");
                        linkTag.setAttribute("href", prefix + "atom-one-dark.css");
                        break;
                    case "solarized":
                        console.log("使用主题 solarized dark");
                        set_code_theme("solarized");
                        linkTag.setAttribute("href", prefix + "solarized-dark.css");
                        break;
                    case "xcode":
                        console.log("使用主题 xcode");
                        set_code_theme("xcode");
                        linkTag.setAttribute("href", prefix + "xcode.css");
                        break;
                    case "xterm":
                        console.log("使用主题 xterm");
                        set_code_theme("xterm");
                        linkTag.setAttribute("href", prefix + "xt256.css");
                        break;
                    case "dracula":
                        console.log("使用主题 dracula");
                        set_code_theme("dracula");
                        linkTag.setAttribute("href", prefix + "dracula.css");
                        break;
                    default:
                        console.log("使用默认主题 github");
                        set_code_theme("github");
                        linkTag.setAttribute("href", prefix + "github.css");
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
<style>
    @import "../custom/custom.css";
</style>