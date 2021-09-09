<template>
    <div class="zhuanlan">
        <Top_banner></Top_banner>
        <div class="zhuanlan_body">
            <div
                style="border: 4px solid #2f2f2f;width: fit-content;padding: 20px 16px;
                margin: 0 auto"
                class="animated fadeInDown"
            >
                <h3 id="title">专栏<br>
                    <span style="font-size: 1rem;font-weight: normal;font-family: Monaco, Consolas, 'DejaVu Sans Mono', monospace">renj.io</span>
                </h3>
            </div>
            <p style="font-weight: 200;padding: 20px 10px" class="animated fadeInDown">专栏 . 记录点点滴滴</p>
            <div class="zhuanlan-list">
                <el-row :gutter="10">
                    <el-col :xs="12" :sm="8" :md="6" :lg="6" :xl="6" v-for="z in zhuanlan_list" :key="z.link">
                        <div class="zhuanlan-tab animated fadeInDown" >
                            <h4><a :href="'/zhuanlan/' + z.link">{{z.title}}</a></h4>
                            <div style="height: 80px">
                                <p class="zhuanlan-content">{{z.content}}</p>
                            </div>
                            <div class="post-count">
                                <span>本专栏共包含{{z.posts.length}}篇文章</span>
                            </div>
                            <div class="zhuanlan-button">
                                <a :href="'/zhuanlan/' + z.link">进入专栏</a>
                            </div>
                        </div>
                    </el-col>
                </el-row>
            </div>
        </div>
        <Bottom_banner></Bottom_banner>
    </div>
</template>

<script>
    import api_zhuanlan from "../api/zhuanlan";
    import Top_banner from "../components/top_banner";
    import Bottom_banner from "../components/bottom_banner";
    export default {
        name: "zhuanlan",
        components: {Bottom_banner, Top_banner},
        data(){
            return{
                empty_zhuanlan: {link: 0, title: "No ZhuanLan", date: "2021-3-30", content: "Api Get failed when this shows."},
                zhuanlan_list: []
            }
        },
        mounted() {
            this.get_zhuanlan_list();
        },
        methods: {
            get_zhuanlan_list(){
                this.$http.get(api_zhuanlan.api_zhuanlan_list).then(res => {
                    if (res.data.data) {
                        this.zhuanlan_list = res.data.data;
                    }
                }).catch(e => {this.zhuanlan_list = [this.empty_zhuanlan]});
            }
        }
    }
</script>

<style>
html, body {
    background-color: var(--body-background)!important;
}
</style>
<style scoped>
    .zhuanlan {
        padding: 30px 10px;
        text-align: center;
        background-color: var(--body-background);
    }
    .zhuanlan .zhuanlan_body{
        padding: 10px 30px;
        text-align: center;
        margin-top: 120px;
    }
    @media (max-width: 460px) {
        .zhuanlan .zhuanlan_body {
            padding: 0;
        }
        .zhuanlan .zhuanlan-list {
            padding: 30px 4px;
        }
    }
    @media (max-width: 380px) {
        .zhuanlan-list /deep/ .el-row .el-col{
            width: 100%!important;
        }
    }
    .zhuanlan_body #title{
        font-size: 34px;
        font-weight: 500;
    }
    .zhuanlan-list{
        padding: 30px 1rem;
        max-width: 760px;
        margin: 0 auto;
        animation-delay: .1s;
    }
    .zhuanlan-list .zhuanlan-tab{
        font-size: 16px;
        margin-bottom: 10px;
        background-color: var(--post-background);
        box-shadow: -2px 5px 10px 2px var(--post-box);
        height: 200px;
        padding: 26px 8px;
    }
    .zhuanlan-tab h4{
        font-size: 1.2rem;
        margin-bottom: 5px;
    }
    .zhuanlan-tab .zhuanlan-content {
        margin: 10px auto 0;
        font-size: 14px;
        text-align: left;
        padding: 0 10px;
        overflow: hidden;
        text-overflow: ellipsis;
        color: gray;
        /*white-space: nowrap;*/
        line-height: 22px;
        height: 4rem;
    }
    .post-count{
        display: inline-block;
        font-weight: bold;
        padding: 4px 10px;
        background-color: var(--pagination-bg);
        color: #576b95;
        font-size: 14px;
        margin-top: 4px;
    }
    .zhuanlan-button {
        margin-top: 30px;
    }
    .zhuanlan-button a {
        border: 2px solid #42b983;
        border-radius: 4px;
        padding: 8px 12px;
        font-size: 14px;
    }
    .zhuanlan-button a:hover {
        background-color: rgba(17,166,104,.06);
    }
</style>
<style>
    @import "../custom/custom.css";
</style>