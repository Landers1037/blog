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
                    <div style="position:relative;">
                        <a class="post-a" :href="'/p/'+a.name">{{a.title}}</a>
                        <span class="post-date" v-if="a.date.indexOf('-')!==-1">{{a.date}}</span>
                    </div>
                    <div class="markdown-body abstract" v-html="mk(a.abstract)"></div>
                    <div class="post-tag" v-if="a.tags && a.tags !== '暂时没有标签'">
                        <el-tooltip v-for="t in tags_to_list(a.tags)"
                                    :key="t"
                                    effect="dark"
                                    :content="'标签: ' + t"
                                    :enterable="false"
                                    placement="bottom-start">
                            <el-tag
                                type="info"
                                size="small"
                                style="cursor: pointer;margin-right: 8px"
                                @click="$router.push('/t/' + t)"
                            >{{t}}</el-tag>
                        </el-tooltip>
                    </div>
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
            tags_to_list(tags){
                return tags.split(" ");
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
                }, customData.loading_duration);
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
                this.$nextTick(()=>{
                    let pres = document.getElementsByTagName("pre");
                    for(let i=0;i<pres.length;i++){
                        pres[i].classList.add("hljs");
                    }
                });
                return marked(code);
            },
        }
    }
</script>

<style scoped>
    .search{
        position: relative;
        padding: 30px 10px;
    }
    .sebar{
        max-width: 520px;
        margin: 80px auto 20px;
    }
    .wrapper{
        margin:  0 auto;
        padding: 10px;
        max-width: 980px;
    }
    .noresult{
        padding: 10px;
        color: #ff7420;
    }
    .search /deep/ .el-input-group__append{
        background-color: var(--main-color);
        color: var(--text-color);
        font-size: 20px;
    }
    .search /deep/ .el-icon-search {
        color: #ffffff;
    }
    @media (max-width: 460px) {
        .wrapper{
            padding: 0;
        }
        .search /deep/ .el-input-group__append{
            font-size: 16px;
        }
    }
    .search /deep/ .el-input__inner:focus{
        border-color:  #DCDFE6;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>