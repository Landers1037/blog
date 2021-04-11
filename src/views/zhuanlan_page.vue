<template>
    <div class="zhuanlan-page">
        <div class="zhuanlan-top">
            <h2><a style="color: #fff" href="/">{{zhuanlan.title}}</a></h2>
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
                    <a class="post-a" :href="'/p/'+p.name">{{p.title}}</a>
                    <div class="markdown-body abstract" v-html="mk(p.abstract)"></div>
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
    .zhuanlan-page{
        padding: 0;
    }
    .zhuanlan-top{
        padding: 40px 10px;
        background-color: #409eff;
        color: #ffffff;
    }
    .title2{
        background-color: #61b1ff;
        display: inline;
        margin: 0 auto;
        padding: 4px 8px;
        border-radius: 4px;
    }
    .zhuanlan-body {
        margin: 0 auto 12px auto;
        max-width: 780px;
    }
    .content{
        font-size: 18px;
        color: #222222;
        padding: 20px 10px;
        text-align: left;
    }
    .posts{
        text-align: left;
        position: relative;
        padding: 10px 10px;
        box-shadow: 0 2px 6px #a0a0a0;
        border-radius: 3px;
    }
    .post-a{
        font-size: 18px;
        font-weight: bold;
        color: #409eff;
        border-bottom: 1px solid #cfcfcf;
        cursor: pointer;
        padding-bottom: 2px;
    }
    .post-a:hover{
        color: #2f343f;
    }
    .abstract{
        font-size: 15px;
        color: #3f3f3f;
        margin-top: 8px;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>