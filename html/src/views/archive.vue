<template>
    <div class="archive">
    <top_banner></top_banner>
        <div class="wrapper">
                <div class="articlelists" v-infinite-scroll="load" infinite-scroll-disabled="lazyloading" infinite-scroll-distance="20" style="overflow:auto">
                    <div v-for="a in postSlice" :key="a.title" class="post animated slideInDown">
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
                    <p v-if="lazyloading" style="color: #6699ff;padding: 4px 0">( •̀ ω •́ )✧加载中...</p>
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
        name: "archive",
        components: {Top_banner, Bottom_banner},
        data(){
            return{
                custom: customData,
                posts: [],
                postSlice: [],
                count: 0,
                lazyloading: false
            }
        },
        created(){
            this.getarticles();
        },
        mounted() {
            this.loading(customData.loading_duration);
        },
        methods:{
            getarticles(){
                let _this = this;
                    this.$http.get(api_article.api_article_list).then(res=>{
                        _this.posts = res.data.data;
                        this.postSlice = this.posts.slice(0,5);
                        this.count = 1;
                    }).catch(err=>{
                        this.$message.error('出现错误了，请求文章失败');
                    });
            },
            tags_to_list(tags){
                return tags.split(" ");
            },
            back(){
                this.$router.push("/")
            },
            load(){
               let slice = this.posts.slice(this.count*5+1,(this.count+1)*5);
               this.count++;
               this.lazyloading = true;
               setTimeout(()=>{
                   for(var i=0;i<slice.length;i++){
                       this.postSlice.push(slice[i])
                   }
                   this.lazyloading = false;
               },1800);
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
            loading(d) {
                const loading = this.$loading({
                    lock: true,
                    text: '文章加载中...',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                setTimeout(() => {
                    loading.close();
                }, d);
            }
        }
    }
</script>

<style scoped>
    .archive {
        padding: 30px 10px;
    }
    #title{
        background-color: #363636;
        color: white;
        padding: 8px 10px;
        cursor: pointer;
        font-family: mo,monospace;
    }
    .wrapper{
        margin:  0 auto;
        padding: 0;
        max-width: 980px;
    }
    .articlelists{
        height: calc(100vh - 240px);
        padding: 0;
        margin: 0;
    }
    ::-webkit-scrollbar{
        width: 15px;
        height: 10px;
    }
    ::-webkit-scrollbar-thumb{
        background-color: #378de5;
    }
    .articlelists .post{
        text-align: left;
        position: relative;
        padding: 10px 10px;
        box-shadow: -1px 2px 8px 2px #e0e0e0;
        margin: 6px 8px 12px;
        border-radius: 2px;
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
    .post-date {
        font-size: .7rem;
        color: #909090;
        position: absolute;
        right: 4px;
        top: 4px;
    }
    .post .post-tag {
        margin-top: 20px;
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
    @media (max-width: 460px) {
        .wrapper{
            padding: 0;
        }
        .articlelists{
            height: 550px;
        }
        .bottom{
            margin-top: 40px;
        }
    }
    /deep/ .el-divider.el-divider--horizontal{
        margin: 36px 0 24px 0;
    }
    /deep/ .el-divider__text.is-center{
        display: block;
        width: fit-content;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>