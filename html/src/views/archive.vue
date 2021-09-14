<template>
    <div class="archive">
    <top_banner></top_banner>
        <div class="wrapper">
            <div class="archive-lists">
                <div class="archive-info animated slideInDown" v-for="a in archives">
                    <p><a class="archive-date" :href="'/archive/'+a.date">{{ a.date }}期</a></p>
                    <p class="archive-count">文章数 {{ a.count }}</p>
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
        name: "archive",
        components: {Top_banner, Bottom_banner},
        data(){
            return{
                custom: customData,
                archives: [],
                lazyloading: false
            }
        },
        created(){
            this.get_archive();
        },
        mounted() {
            this.loading(customData.loading_duration);
        },
        methods:{
            get_archive() {
                let _this = this;
                    this.$http.get(api_article.api_article_archive).then(res=>{
                        _this.archives = res.data.data;
                    }).catch(err=>{
                        this.$message.error('出现错误了，请求归档失败');
                    });
            },
            back(){
                this.$router.push("/")
            },
            loading(d) {
                const loading = this.$loading({
                    lock: true,
                    text: '归档信息加载中...',
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
    .wrapper{
        margin: 70px auto 0;
        padding: 10px 0;
        max-width: 980px;
    }
    @media (max-width: 460px) {
        .wrapper{
            padding: 0;
        }
    }
    .archive-lists {
        padding: 10px;
    }
    .archive-info {
        font-size: 1.2rem;
        text-align: left;
        background-color: var(--post-background);
        padding: 10px 20px;
        margin-bottom: 10px;
        border-radius: 2px;
        vertical-align: middle;
        box-shadow: -1px 1px 10px 2px var(--post-box);
        width: inherit;
    }
    .archive-info .archive-date {
        vertical-align: middle;
        font-weight: bold;
        display: inline-block;
    }
    .archive-info .archive-date a {
        vertical-align: middle;
        font-weight: bold;
    }
    .archive-info .archive-count {
        text-align: left;
        font-size: .9rem;
        vertical-align: middle;
        display: inline-block;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>