<template>
    <div class="about">
        <div class="head">
            <h1 style="margin-bottom: 40px">关于</h1>
            <p>这是我的个人博客，记录了学习路上的点点滴滴</p>
            <p>使用Go作为后端api,前端由Vue和ElementUI开发</p>
        </div>
        <div style="margin-top: 40px">
            <ul>我的个人主页 <a :href=custom.site_url>{{custom.site_name}}</a></ul>
            <ul>我的Github <a :href="'https://github.com/' + custom.github">{{custom.github}}</a></ul>

        </div>
        <div class="thanks">
            <p>感谢</p>
            <ul>框架 <a href="https://cn.vuejs.org/">Vuejs</a></ul>
            <ul>引擎 <a href="https://github.com/gin-gonic/gin">Gin</a></ul>
            <ul>证书 <a href="https://freessl.cn/">FreeSSL</a></ul>
        </div>
        <div style="margin-bottom: 1.6rem;font-weight: bold">
            <p style="font-size: 1.2rem;">框架设计 - renj.io</p>
            <p style="padding: 10px"><a href="https://github.com/Landers1037/blog">Landers1037/blog</a></p>
        </div>
        <div style="margin-bottom: 1.6rem;font-weight: bold;">
            <p style="font-size: 1.2rem;">404页面设计 - renj.io</p>
            <p style="padding: 10px"><a href="/404">Error page not found</a></p>
        </div>

        <el-button type="primary" @click="back">返回主页</el-button>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    export default {
        name: "about",
        data(){
            return{
                custom: customData,
            }
        },
        methods:{
            back(){
                //this.$router.push("/")
                window.history.back();
            },
        }
    }
</script>

<style scoped>
    .about {
        padding: 20px 10px;
    }
    .thanks{
        padding: 20px 0 30px 0;
        text-align: left;
        width: 260px;
        margin: 0 auto;
    }
    .thanks p {
        text-align: center;
        font-size: 20px;
        font-weight: bold;
        margin-bottom: 15px;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>