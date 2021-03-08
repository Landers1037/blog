<template>
    <div class="we-post">
        <div class="header">
            <span><i class="el-icon-arrow-left" style="margin-right: 10px" @click="back"></i>Landers</span>
        </div>
        <div class="we-body" v-show="load">
            <div class="title">
                <p style="font-size: 18px">{{title}}</p>
                <p style="color: #afafaf;font-size: 12px;font-weight: bold">Landers&emsp;&emsp;{{date}}</p>
            </div>

            <div class="markdown-body weui-article" v-html="post">

            </div>
        </div>
        <div class="weui-msg" v-show="error" style="padding-top: 80px">
            <div class="weui-msg__icon-area"><i class="weui-icon-warn weui-icon_msg"></i></div>
            <div class="weui-msg__text-area">
                <h2 class="weui-msg__title">操作失败</h2>
                <p class="weui-msg__desc">当前文章加载失败，如果刷新无效请前往<a href="/">博客主页</a></p>
            </div>
            <div class="weui-msg__tips-area">
                <p class="weui-msg__tips">文章API请求失败，如果你可以访问本站点其他页面，可能该错误由数据库请求错误导致，你可以前往<a href="https://github.com/landers1037/goblog">Github</a>提交该错误</p>
            </div>
            <div class="weui-msg__opr-area">
                <p class="weui-btn-area">
                    <a href="javascript:history.back();" class="weui-btn weui-btn_default">返回上页</a>
                </p>
            </div>
            <div class="weui-msg__extra-area">
                <div class="weui-footer">
                    <p class="weui-footer__links">
                        <a href="javascript:" class="weui-footer__link">Landers1037 Blog</a>
                    </p>
                    <p class="weui-footer__text">Copyright &copy; 2017-2020 renj.io</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import '@/assets/weui.min.css'
    export default {
        name: "post",
        data(){
            return{
                url: this.$route.params.url,
                title: "加载中",
                date: '2020',
                post: '',
                content: '',
                load: false,
                error: false
            }
        },
        mounted(){
            let _this = this;
            this.$http.get('/api/article/post',{params:{name:this.url}}).then(res=>{
                if(res.data.data.title === ''){
                    _this.error = true;
                }
                else{
                    let content = res.data.data["content"];
                    _this.title = res.data.data["title"];
                    _this.date = res.data.data["date"];
                    _this.mk(content);
                    _this.load = true;
                    _this.error = false;
                }
            }).catch(err=>{
                _this.$message.error('出现错误了，请求文章失败');
                _this.error = true;
                _this.load = false;
            })
        },
        methods:{
            back(){
                this.$router.push('/newui')
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
            }
        }
    }

</script>

<style scoped>
    .we-post{position: relative;background-color: #f0f0f0}
    .we-post .header{
        text-align: left;
        background-color: #06ae56;
        padding: 15px 15px;
        color: white;
        position: fixed;
        width: 100%;
        font-size: 20px;
    }
    .we-post .we-body{
        padding-top: 60px;
        height: fit-content;
    }
    .we-post .we-body .title{
        text-align: left;
        padding: 6px 10px 15px 10px;
    }
    .we-post .we-body .markdown-body{
        text-align: left;
        padding: 10px 10px;
    }
</style>
<style>
    #app{padding: 0;margin: 0}
    html{background-color: #f0f0f0}
</style>