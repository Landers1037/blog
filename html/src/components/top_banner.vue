<template>
    <div class="header">
        <div class="animated slideInDown">
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
        margin-bottom: 30px;
    }
    #title{
        background-color: #363636;
        color: white;
        padding: 8px 10px;
        cursor: pointer;
        font-family: mo,monospace;
    }
    .header /deep/ .el-divider__text {
        width: fit-content;
        border-radius: 2px;
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