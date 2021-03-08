<template>
    <div class="about">
        <div class="head">
            <h1 style="margin-bottom: 40px">关于</h1>
            <p>这是我的个人博客，由我的个人站点拆分所得。对文章获取和标签归类进行了优化</p>
            <p>使用Go作为后端api,前端由Vue和ElementUI开发</p>
        </div>
        <div style="margin-top: 40px">
            <ul>我的个人主页 <a href="http://renj.io">renj.io</a></ul>
            <ul>我的Github <a href="https://github.com/landers1037">Landers1037</a></ul>

        </div>
        <div class="thanks">
            <p>感谢</p>
            <ul>框架 <a href="https://cn.vuejs.org/">Vuejs</a></ul>
            <ul>引擎 <a href="https://github.com/gin-gonic/gin">Gin</a></ul>
            <ul>证书 <a href="https://freessl.cn/">FreeSSL</a></ul>
        </div>
        <el-button type="primary" @click="back">返回主页</el-button>
    </div>
</template>

<script>
    export default {
        name: "about",
        methods:{
            back(){
                //this.$router.push("/")
                window.history.back();
            },
        }
    }
</script>

<style scoped>
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