<template>
    <div class="tag">
        <div class="header">
            <div class="animated slideInDown">
                <label id="title" @click="back">Landers 1037</label>
            </div>
            <el-divider><span style="font-family: 'DejaVu Sans Mono','Source Code Pro','Liberation Mono',monospace;font-size: 14px">Never stop debugging</span></el-divider>
<!--            <p style="margin-top: 30px"><span style="color: #c0c0c0">&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;</span>Never stop debuging<span style="color: #c0c0c0">&#45;&#45;&#45;&#45;&#45;&#45;&#45;&#45;</span></p>-->
        </div>
        <div class="title">
            <p>{{tag}}</p>
        </div>
        <div class="wrapper">
        <div class="articlelists">
            <div v-for="a in posts" :key="a.title" class="post animated slideInDown">
                <a class="post-a" :href="'/p/'+a.posturl">{{a.title}}</a>
                <div class="markdown-body abstract" v-html="mk(a.content)"></div>
            </div>
        </div>
        </div>
        <div class="bottom">
            <p><el-icon class="el-icon-lollipop"></el-icon><a style="color: #5f5f5f;font-weight: bold;margin-right: 8px" href="http://renj.io">By Landers</a>
                <el-icon class="el-icon-coffee-cup"></el-icon><a style="color: #5f5f5f;font-weight: bold" href="https://landers1037.github.io">Github</a>
                <br><span style="font-size: 12px;color: #2c3e50">Golang & Vue</span></p>

        </div>
    </div>
</template>

<script>
    export default {
        name: "tag",
        data(){
            return{
                tag: this.$route.params.tag,
                posts:[]
            }
        },
        mounted() {
            this.get();
        },
        methods:{
            get(){
                let _this = this;
                if(this.tag !== ""){
                    this.$http.get("/api/article/tag",{params:{"tag":this.tag}}).then(res=>{
                       _this.posts = res.data.data;
                    })
                }else {
                    this.$message.error('出现错误了，请求文章失败');
                }
            },
            back(){
                this.$router.push("/")
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
            },
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
    .wrapper{
        max-width: 980px;
        margin: 25px auto 0;
    }
    .articlelists .post{
        text-align: left;
        position: relative;
        padding: 10px 10px;
        box-shadow: -1px 1px 4px #a0a0a0;
        margin-bottom: 12px;
        border-radius: 6px;
    }
    .post-a{
        font-size: 18px;
        color: #409eff;
        border-bottom: 1px solid #cfcfcf;
        cursor: pointer;
        padding-bottom: 2px;
        font-weight: bold;
    }
    .post-a:hover{
        color: #2f343f;
    }
    .abstract{
        font-size: 15px;
        color: #3f3f3f;
        margin-top: 8px;
    }
    .bottom{
        margin-top: 20px;
        font-family: "DejaVu Sans Mono","Segoe UI",Monaco,monospace;
    }
    /deep/ .el-divider.el-divider--horizontal{
        margin: 36px 0 24px 0;
    }
    /deep/ .el-divider__text.is-center{
        display: block;
        width: fit-content;
    }
</style>