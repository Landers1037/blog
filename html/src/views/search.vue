<template>
    <div class="search">
    <top_banner></top_banner>
        <div class="sebar">
            <el-input placeholder="请输入查找内容" v-model="word">
                <el-button slot="append" icon="el-icon-search" @click="search"></el-button>
            </el-input>
        </div>
        <div class="wrapper">
            <div class="articlelists">
                <div v-for="a in posts" :key="a.title" class="post animated slideInDown">
                    <a class="post-a" :href="'/p/'+a.name">{{a.title}}</a>
                    <div class="markdown-body abstract" v-html="mk(a.abstract)"></div>
                </div>
                <div v-if="posts==null">
                    <p class="noresult">没有找到你想要的结果😭</p>
                </div>
            </div>
        </div>
        <bottom_banner></bottom_banner>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    import api_article from "../api/article";
    import Top_banner from "../components/top_banner";
    import Bottom_banner from "../components/bottom_banner";
    export default {
        name: "search",
        components: {Bottom_banner, Top_banner},
        data(){
            return{
                custom: customData,
                posts: [],
                word: ""
            }
        },
        watch:{
          word:function () {
              let _this = this;
              setTimeout(function () {
                  _this.search();
              },1400)
          }
        },
        mounted(){

        },
        methods:{
            back(){
                this.$router.push("/")
            },
            loading(){
                const loading = this.$loading({
                    lock: true,
                    text: '加载中...',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.6)'
                });
                setTimeout(() => {
                    loading.close();
                }, 1200);
            },
            search(){
              let s = this.word;
              let _this = this;
              if(s.length!==0){
                  this.$http.get(api_article.api_article_search,{params:{"key": s}}).then(res=>{
                      if(res.data.length>0){
                          _this.posts = res.data;
                          _this.loading();
                      }else {
                          _this.posts = null;
                          _this.loading();
                      }
                  }).catch(err=>{
                      _this.posts = [];
                  })
              }
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
    .search{
        position: relative;
        padding: 30px 10px;
    }
    .sebar{
        max-width: 520px;
        margin: 30px auto 20px;
    }
    .wrapper{
        margin:  0 auto;
        padding: 10px;
        max-width: 980px;
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
        margin-top: 6px;
        font-family: "DejaVu Sans Mono","Segoe UI",Monaco,monospace;
    }
    .noresult{
        padding: 10px;
        color: #ff7420;
    }
    .search /deep/ .el-input-group__append{
        background-color: #6699ff;
        color: #f7f7f7;
        font-size: 20px;
    }
    @media (max-width: 460px) {
        .wrapper{
            padding: 0;
        }
        .search /deep/ .el-input-group__append{
            font-size: 16px;
        }
    }
    /deep/ .el-divider.el-divider--horizontal{
        margin: 36px 0 24px 0;
    }
    /deep/ .el-divider__text.is-center{
        display: block;
        width: fit-content;
    }
    .search /deep/ .el-input__inner:focus{
        border-color:  #DCDFE6;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>