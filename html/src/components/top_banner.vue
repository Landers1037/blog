<template>
    <div class="header">
        <div class="animated slideInDown" style="padding: 10px">
            <label id="title" @click="back">{{custom.top_banner}}</label>
        </div>
        <el-divider>
            <span style="font-family: 'mo',monospace;font-size: 14px">
                {{custom.top_span}}
            </span>
        </el-divider>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    export default {
        name: "top_banner",
        data(){
            return{
                custom: customData,
            }
        },
        methods: {
            back(){
                // 如果在主页则重载页面 否则返回主页
                if (location.pathname === "/") {
                    location.reload();
                }else {
                    this.$router.push("/");
                }

            },
        }
    }
</script>

<style scoped>
    .header{
        position: fixed;
        z-index: 999;
        top: 0;
        left: 0;
        width: 100%;
        background-color: #fafafa;
        padding: 16px 0;
    }

    #title{
        background-color: #363636;
        color: white;
        padding: 8px 10px;
        cursor: pointer;
        font-family: mo,monospace;
    }
    .header /deep/ .el-divider--horizontal {
        margin: 16px 0 4px 0;
    }
    .header /deep/ .el-divider__text{
        width: fit-content;
        border-radius: 2px;
        background-color: #fafafa;
    }
    @media (max-width: 500px) and (min-width: 320px) {
        .el-divider__text {
            padding: 0 5px;
        }
    }

</style>
<style>
    @import "../custom/custom.css";
</style>