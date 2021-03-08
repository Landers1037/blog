<template>
    <div class="we-article">
        <div class="header">
            <span>Landers</span>
        </div>
        <div class="we-article-body">
            <div class="animated sildeInDown card" v-for="p in posts">
                <p style="background-color: #6699ff;padding: 6px 10px;color: white"><a :href="'/newui/p/'+p.url">{{p.title}}</a>
                    <br><span style="font-size: 10px">Landers1037</span>
                </p>
                <p style="padding: 15px 10px;word-wrap: break-word;word-break: break-all">{{p.content}}</p>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "article",
        data(){
            return{
                posts: [{title: 's',content: 's'}]
            }
        },
        mounted() {
            this.getarticles();
        },
        methods:{
            getarticles(){
                let _this = this;
                this.$http.get("/api/article/posts").then(res=>{
                    _this.posts = res.data.data;

                }).catch(err=>{
                    this.$message.error('出现错误了，请求文章失败');
                });
            },
        }
    }
</script>

<style scoped>
    .we-article{
        position: relative;
        width: 100%;
    }
    .we-article .header{
        text-align: left;
        background-color: #06ae56;
        padding: 15px 15px;
        color: white;
        position: fixed;
        width: 100%;
        font-size: 20px;
    }
    .we-article .we-article-body{
        padding-top: 80px;
        text-align: left;
        padding-left: 6px;
        padding-right: 6px;
    }
    .we-article-body .card{
        background-color: white;
        margin-bottom: 10px;
    }
    .we-article-body .card a{
        color: white;
    }
    .we-article-body .card a:hover{
        color: #2c3e50;
    }
</style>
<style>
    #app{padding: 0;margin: 0}
    html{background-color: #f0f0f0}
</style>