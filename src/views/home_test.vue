<template>
    <div class="home">
        <div class="header">
            <div class="animated slideInDown">
                <label id="title" >Landers 1037</label>
            </div>
            <p style="margin-top: 30px"><span style="color: #c0c0c0">----</span>Never stop debuging<span style="color: #c0c0c0">----</span></p>

        </div>
        <div class="wrapper">
            <div class="contents">
                <div class="right">
                    <div class="img">
                        <img src="../assets/logo.png">
                        <div class="me">
                            <p>Me: Landers</p>
                            <p>My site: <a>lrenj.top</a></p>
                            <p>Github: <a href="https://github.com/landers1037">Landers1037</a></p>
                        </div>
                    </div>
                    <div class="tags">
                        <el-tag
                                v-for="item in tags"
                                :key="item"
                                type="success"
                                effect="dark">
                            {{ item}}
                        </el-tag>
                    </div>
                </div>
                <div class="left">
                    <div class="articlelists">
                        <div v-for="a in posts.list" :key="a.name" class="post">
                            <a class="post-a">{{a.name}}</a>
                            <div class="markdown-body abstract" v-html="mk(a.abstract)"></div>
                            <div class="post-tag">

                            </div>
                        </div>
                    </div>
                    <div class="pagenation">
                        <el-pagination
                                background
                                :current-page=1
                                :page-size=6
                                layout="prev, pager, next"
                                @current-change="reloadpage"
                                :total="posts.total">
                        </el-pagination>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
    import marked from 'marked';

    export default {
        name: "home",
        data(){
            return{
                tags:["sb","nmsl","python","jksnanskas","jishs","sasasa"],
                posts: {
                    total: 200,
                    list: [{name:"TEST1",data:"2020-1-1",abstract:"# hello world",tags:["1","2"]},{name:"测试",data:"2020-1-1",abstract:"## 标题2\n### 标题三\n**粗体**",tags:["1","2"]}]
                }
            }
        },
        mounted() {

        },
        methods:{
            gettags(){
                this.$http.get("/api/article/tags").then(res=>{
                    this.tags = res.data;
                })
            },
            getarticles(){
                this.$http.get("/api/article/posts").then(res=>{
                    this.posts = res.data;
                })
            },
            reloadpage(now){
                console.log(now)
            },
            mk(code) {
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
                return marked(code);
            }
        }
    }
</script>

<style scoped>
    .home{
        user-select: none;
    }
    #title{
        background-color: #363636;
        color: white;
        padding: 6px 10px;
    }
    .wrapper{
        margin:  0 auto;
        padding: 10px;
        max-width: 980px;
    }
    .contents{
        position: relative;
        width: 100%;
        margin-top: 30px;
    }
    .contents .right{
        width: 200px;
        padding: 10px;
        border-radius: 10px;
        box-shadow: -1px 2px 6px #5f5f5f;
        float: right;
    }
    .right .me{
        text-align: left;
        width: 80%;
        padding-left: 20px;
    }
    .right .tags{
        margin-top: 10px;
        text-align: left;
        padding: 10px;
    }
    .right .tags .el-tag{
        margin-right: 10px;
        margin-top: 5px;
    }
    .right .tags /deep/ .el-tag{
        height: 32px;
        line-height: 30px;
        font-size: 13px;
        cursor: pointer;
    }
    .contents .left{
        width: calc(100% - 240px);
        float: left;
    }
    .left .pagenation{
        padding: 20px 0 10px 0;
    }
    .left .articlelists .post{
        text-align: left;
        position: relative;
        padding: 10px 8px;
        box-shadow: -2px 1px 4px #a0a0a0;
        margin-bottom: 10px;
        border-radius: 2px;
    }
    .post .post-a{
        font-size: 18px;
        color: #409eff;
        border-bottom: 1px solid #9f9f9f;
        cursor: pointer;
        padding-bottom: 2px;
    }
    .post .post-a:hover{
        color: #2f343f;
    }
    .post .abstract{
        margin-top: 6px;
    }
    @media (max-width: 750px){
        .contents .left{
            width: 100%;
            float: unset;
        }
        .contents .right{
            display: none;
        }
        .wrapper{
            padding: 0;
        }
    }
</style>