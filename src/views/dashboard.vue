<template>
    <div class="dashboard">
        <el-menu :default-active="activeIndex" class="el-menu" mode="horizontal" @select="handleSelect">
            <el-menu-item index="1">文章管理</el-menu-item>
            <el-submenu index="2">
                <template slot="title">统计数据管理</template>
                <el-menu-item index="2-1">访问量管理</el-menu-item>
                <el-menu-item index="2-2">分享量管理</el-menu-item>
                <el-menu-item index="2-3">点赞量管理</el-menu-item>
                <el-menu-item index="2-4">评论 管理</el-menu-item>
                <el-menu-item index="2-5">统计可视化</el-menu-item>
            </el-submenu>
            <el-menu-item index="3">专栏管理</el-menu-item>
            <el-menu-item index="4">留言管理</el-menu-item>
            <el-menu-item index="5" @click="logout">退出登录</el-menu-item>
            <el-menu-item index="6"><a href="/">返回主页</a></el-menu-item>
        </el-menu>
        <div class="dynamic-area">
<!--            动态组件-->
            <component :is="dynamic_com"></component>
        </div>
    </div>
</template>

<script>
    import Post_card from "../components/post_card";
    import Zhuanlan_card from "../components/zhuanlan_card";
    import Message_card from "../components/message_card";
    import View_card from "../components/view_card";
    import Share_card from "../components/share_card";
    import Like_card from "../components/like_card";
    import Comment_card from "../components/comment_card";
    import Show_card from "../components/show_card";
    export default {
        name: "dashboard",
        components: {Post_card, Zhuanlan_card, Message_card,
            View_card, Like_card, Share_card, Comment_card,
            Show_card},
        data(){
            return{
                activeIndex: 1,
                dynamic_com: '',
                components_map: {
                    "1": 'Post_card',
                    "2-1": 'View_card',
                    "2-2": "Share_card",
                    "2-3": "Like_card",
                    "2-4": "Comment_card",
                    "2-5": "Show_card",
                    "3": 'Zhuanlan_card',
                    "4": 'Message_card',
                }
            }
        },
        mounted() {
            this.dynamic_com = 'Post_card';
        },
        methods: {
            handleSelect(key, keyPath) {
                this.dynamic_com = this.components_map[key];
            },
            logout(){
                localStorage.removeItem("token");
                this.$message("你已经登出 即将返回主页");
                setTimeout(()=>{
                    this.$router.push("/");
                }, 1500);
            }
        }
    }
</script>

<style scoped>
    .dynamic-area {
        padding: 10px;
        text-align: left;
    }
</style>