<template>
    <div class="tag">
    <top_banner></top_banner>
        <div class="title">
            <p>{{tag}}</p>
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
        </div>
        </div>
        <bottom_banner></bottom_banner>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    import api_tags from "../api/tag";
    import Top_banner from "../components/top_banner";
    import Bottom_banner from "../components/bottom_banner";
    export default {
        name: "tag",
        components: {Bottom_banner, Top_banner},
        data(){
            return{
                custom: customData,
                tag: this.$route.params.tag,
                posts:[]
            }
        },
        beforeRouteUpdate(to, from, next) {
            // 路由变化时获取路由中的name
            this.tag = to.params.tag;
            this.get();
            next();
        },
        mounted() {
            this.get();
        },
        methods:{
            get(){
                let _this = this;
                if(this.tag !== ""){
                    this.$http.get(api_tags.api_tags_list,{params:{"tag":this.tag}}).then(res=>{
                       _this.posts = res.data.data;
                    })
                }else {
                    this.$message.error('出现错误了，请求文章失败');
                }
            },
            tags_to_list(tags){
                return tags.split(" ");
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
    .tag {
        padding: 30px 10px;
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
        padding: 16px;
        box-shadow: -1px 2px 8px 2px var(--post-box);
        margin-bottom: 12px;
        border-radius: 2px;
        background-color: var(--post-background);
    }
    .post-a{
        font-size: 18px;
        color: var(--post-title);
        border-bottom: 1px solid var(--border-color);
        cursor: pointer;
        padding-bottom: 2px;
        font-weight: bold;
    }
    .post-a:hover{
        color: var(--post-title-hover);
    }
    .post-date {
        font-size: .7rem;
        color: var(--post-date);
        position: absolute;
        right: 4px;
        top: 4px;
    }
    .post .post-tag {
        margin-top: 20px;
    }
    .abstract{
        font-size: 15px;
        margin-top: 8px;
        color: var(--post-color);
    }
    .bottom{
        margin-top: 20px;
        font-family: "DejaVu Sans Mono","Segoe UI",Monaco,monospace;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>