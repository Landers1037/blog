<template>
    <div class="zhuanlan-page">
        <div class="zhuanlan-top">
            <h2 style="font-size: 1.8rem"><a style="color: #fff" href="/zhuanlan">{{zhuanlan.title}}</a></h2>
            <div style="margin-top: 20px">
                <p class="title2">创建时间: {{zhuanlan.date}}</p>
            </div>
        </div>
        <div class="zhuanlan-body">
            <div :v-html="zhuanlan.content" class="content">
                {{zhuanlan.content}}
            </div>
            <div>
                <h3 style="margin: 10px;padding: 10px">专题文章</h3>
                <div v-for="p in zhuanlan.posts" :key="p.name" class="posts animated fadeInDown">
                    <div style="position:relative;">
                        <a class="post-a" :href="'/p/'+p.name">{{p.title}}</a>
                        <span class="post-date" v-if="p.date.indexOf('-')!==-1">{{p.date}}</span>
                    </div>
                    <div class="markdown-body abstract" v-html="mk(p.abstract)"></div>
                    <div class="post-tag" v-if="p.tags && p.tags !== '暂时没有标签'">
                        <el-tooltip v-for="t in tags_to_list(p.tags)"
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
    </div>
</template>

<script>
    import api_zhuanlan from "../api/zhuanlan";
    import Bottom_banner from "../components/bottom_banner";
    export default {
        name: "zhuanlan_page",
        components: {Bottom_banner},
        data(){
            return{
                link: this.$route.params.link,
                zhuanlan: {},
            }
        },
        mounted() {
            this.get_zhuanlan();
        },
        methods: {
            get_zhuanlan(){
                this.$http.get(api_zhuanlan.api_zhuanlan_detail + this.link).then(res => {
                    if (res.data.data === "failed") {
                        this.$message.error("获取专栏信息失败");
                    }else {
                        this.zhuanlan = res.data.data;
                    }
                }).catch();
            },
            tags_to_list(tags){
                return tags.split(" ");
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
    .zhuanlan-page{
        padding: 0 0 30px 0;
    }
    .zhuanlan-top{
        padding: 50px 10px;
        background-color: #409eff;
        color: #ffffff;
    }
    .title2{
        background-color: #61b1ff;
        display: inline;
        margin: 0 auto;
        padding: 4px 8px;
        border-radius: 4px;
        font-size: .9rem;
    }
    .zhuanlan-body {
        margin: 0 auto 12px auto;
        max-width: 780px;
    }
    .content{
        font-size: 18px;
        color: var(--text-color);
        padding: 20px 10px;
        text-align: left;
    }
    .posts{
        text-align: left;
        position: relative;
        padding: 16px;
        box-shadow: -1px 2px 8px 2px var(--post-box);
        border-radius: 3px;
        margin-bottom: 8px;
        background-color: var(--post-background);
        color: var(--post-color);
    }
    .post-a{
        font-size: 18px;
        font-weight: bold;
        color: var(--post-title);
        border-bottom: 1px solid var(--border-color);
        cursor: pointer;
        padding-bottom: 2px;
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
    .post-tag {
        margin-top: 20px;
    }
    .abstract{
        font-size: 15px;
        color: var(--post-color);
        margin-top: 8px;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>