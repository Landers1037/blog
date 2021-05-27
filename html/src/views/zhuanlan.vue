<template>
    <div class="zhuanlan">
        <Top_banner></Top_banner>
        <div class="zhuanlan_body">
            <h3 id="title">专栏</h3>
            <div class="zhuanlan-list">
                <div class="zhuanlan-tab animated fadeInDown" v-for="z in zhuanlan_list" :key="z.link">
                    <h4><a :href="'/zhuanlan/' + z.link">{{z.title}}</a></h4>
                    <p>{{z.content}}</p>
                    <div class="post-count">
                        <span>本专栏共包含{{z.posts.length}}篇文章</span>
                    </div>
                </div>
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

<style scoped>
    .zhuanlan {
        padding: 30px 10px;
    }
    .zhuanlan .zhuanlan_body{
        text-align: left;
        padding: 10px;
    }
    .zhuanlan_body #title{
        font-size: 36px;
    }
    .zhuanlan-list{
        padding: 15px 0;
    }
    .zhuanlan-list .zhuanlan-tab{
        font-size: 16px;
        margin-bottom: 10px;
    }
    .zhuanlan-tab h4{
        font-size: 26px;
        margin-bottom: 5px;
    }
    .post-count{
        display: inline-block;
        padding: 4px;
        background-color: #f0f0f0;
        color: #576b95;
        font-size: 14px;
        margin-top: 4px;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>