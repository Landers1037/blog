<template>
    <div>
    <div class="weui-index">
        <div class="header">
            <span>Landers</span>
        </div>
        <div class="chat">
            <p class="animated slideInLeft pp">
                😘你好,这是我的博客新版页面
            </p>
            <div v-show="dynamic_show" class="animated slideInLeft pp" v-html="dynamic_text">

            </div>
        </div>
        <div class="weui-tabbar" style="padding: 12px;align-self: flex-end;width: 100%">
            <div class="weui-tabbar__item weui-bar__item_on">
                <p class="weui-tabbar__label">欢迎</p>
            </div>
            <div class="weui-tabbar__item" @click="postslist">
                <p class="weui-tabbar__label">文章列表</p>
            </div>
            <div class="weui-tabbar__item" @click="explore">
                <p class="weui-tabbar__label">发现</p>
            </div>
            <div class="weui-tabbar__item" @click="about">
                <p class="weui-tabbar__label">关于</p>
            </div>
        </div>
    </div>
    <div class="mobile">
        <!--            宽度超出后的提示-->
        <p>请在移动设备查看此页面</p>
    </div>
    </div>
</template>

<script>
    import '@/assets/weui.min.css'
    export default {
        name: "index",
        data(){
            return{
                dynamic_text: "",
                dynamic_show: false,
                posts: [{title: "test1","url": 'test'},{title: "test2",url: "test"}]
            }
        },
        methods:{
            about(){
                this.dynamic_text = "<p>关于本页面：</p>" +
                    "<p>❤ with WeUI 使用WeUI开发，作为适配移动端页面测试</p>" +
                        "<p>试运行，如遇BUG欢迎反馈<a href='https://github.com/landers1037/goblog'>Github issue</a></p>";
                this.dynamic_show = false;
                setTimeout(()=>{
                    this.dynamic_show = true;
                },500);
            },
            explore(){
                // this.$router.push('/explore')
                this.dynamic_show = false;
                this.dynamic_text = "<p>探索页面仍在测试中，请等待上线\n</p>" + "本站点的错误提交页面<a href='https://github.com/landers1037/goblog'>Github</a></br>" +
                        "我的个人主页🍪<a href='http://renj.io'>renj.io</a><br>" +
                        "我的项目主页🚀<a href='http://mgek.cc'>Mgek</a>";

                setTimeout(()=>{
                    this.dynamic_show = true;
                },500);

            },
            postslist(){
                let _this = this;
                this.$http.get("/api/article/posts",{params:{"p":1}}).then(res=> {
                    this.posts = res.data.data;
                    //请求成功后显示文章列表
                    let posts = [];
                    for(var i=0;i<_this.posts.length;i++){
                        posts.push("<p>"+_this.posts[i].title+"<a href='/newui/p/"+_this.posts[i].url+"'> 链接</a>"+"</p>");
                    }
                    let html = "<p>最新文章列表</p></br>" + posts.join("") + "<p>全部文章 <a href='/newui/article'>链接</a></p>";
                    _this.dynamic_show = false;
                    _this.dynamic_text = html;
                    setTimeout(()=>{
                        this.dynamic_show = true;
                    },500);
                }).catch(()=>{
                    _this.$message.error("请求文章列表失败");
                })
            }
        }
    }
</script>

<style scoped>
    .weui-index{
        text-align: left;
        height: 100vh;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        align-items: flex-start;
        background-color: #f0f0f0;
    }
    .weui-index /deep/ a{
        color: #0000ff;
    }
    .mobile{
        display: none;
        margin-top: 60px;
        border-radius: 12px;
        background-color: #06ae56;
        padding: 12px;
        color: white;
        transition: all 0.4s ease;
    }
    .weui-index .header{
        align-self: flex-start;
        background-color: #06ae56;
        padding: 15px 15px;
        color: white;
        width: 100%;
        font-size: 20px;
    }
    .weui-index .chat{
        margin-left: 10px;
        display: inline-flex;
        width: 100%;
        flex-wrap: wrap;
    }
    .weui-index .chat .pp{
        background-color: #06ae56;
        border-radius: 12px;
        padding: 10px;
        width: 80%;
        flex-wrap: wrap;
        word-wrap: break-word;
        margin-bottom: 10px;
    }
    /*高宽度下隐藏*/
    @media (min-width: 640px) {
        .weui-index{display: none}
        .mobile{display: block}
    }
    .weui-index .weui-tabbar__label{
        font-size: 12px;
    }
</style>
<style>
    #app {margin-top: 0;padding: 0}
</style>