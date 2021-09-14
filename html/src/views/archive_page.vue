<template>
    <div class="archive-page">
    <top_banner></top_banner>
        <div class="wrapper">
            <h3 class="title">{{title}}期</h3>
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
        name: "archive-page",
        components: {Top_banner, Bottom_banner},
        data(){
            return{
                title: this.$route.params.date,
                custom: customData,
                posts: [],
                postSlice: [],
                count: 0,
                lazyloading: false
            }
        },
        mounted() {
            this.getarticles();
            this.loading(customData.loading_duration);
        },
        methods:{
            getarticles(){
                let _this = this;
                    this.$http.get(api_article.api_article_archives + "?date=" + _this.title).then(res=>{
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
                if (this.posts.length <= 5) {
                    this.lazyloading = false;
                }else {
                    let slice = this.posts.slice(this.count*5+1,(this.count+1)*5);
                    this.count++;
                    this.lazyloading = true;
                    setTimeout(()=>{
                        for(let i=0;i<slice.length;i++){
                            this.postSlice.push(slice[i])
                        }
                        this.lazyloading = false;
                    },1500);
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
    .archive-page {
        padding: 30px 10px;
    }
    .wrapper{
        margin: 70px auto 0;
        padding: 10px 0;
        max-width: 980px;
    }
    .wrapper .title {
        margin-bottom: 10px;
    }
    .articlelists{
        height: calc(100vh - 260px);
        padding: 0;
        margin: 0;
    }
    @media (max-width: 460px) {
        .wrapper{
            padding: 0;
        }
    }
    .articlelists .post {
        margin: 10px;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>