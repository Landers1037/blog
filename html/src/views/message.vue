<template>
    <div class="message">
    <top_banner></top_banner>
        <div class="box">
            <p>在这里写下你的留言</p>
            <div class="messageBox">
                <ul v-for="m in messages" id="mes" class="animated fadeInDown" style="animation-delay: .1s"><i style="font-size: 20px;padding-right: 6px" class="el-icon-chat-round"></i>{{m.message}}</ul>
            </div>
            <el-input
                    type="textarea"
                    class="yourmes"
                    placeholder="请输入留言"
                    maxlength="300"
                    show-word-limit
                    :autosize="{ minRows: 6, maxRows: 8}"
                    v-model="textarea">
            </el-input>
        </div>
        <el-button class="btnSend" icon="el-icon-chat-square" @click="sendmes">发表留言</el-button>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    import api_message from "../api/message";
    import Top_banner from "../components/top_banner";
    export default {
        name: "message",
        components: {Top_banner},
        data(){
            return{
                custom: customData,
                messages:[],
                textarea: "",
                saved : false
            }
        },
        mounted(){
            this.getmes();
        },
        methods:{
            back(){
                this.$router.push("/")
            },
            getmes(){
                this.$http.get(api_message.api_message_all).then(res=>{
                    this.messages = res.data;
                })
            },
            sendmes(){
                let data = {};
                let _this = this;
                if(this.textarea.trim().length>=4 && !_this.saved){
                    data = {"mes":this.textarea};
                    this.$http.post(api_message.api_message_add, data).then(res=>{
                        if(res.data === "saved"){
                            _this.tishi();
                            _this.saved = true;
                            setTimeout(function () {
                                _this.getmes();
                            },500);
                        }else {
                            _this.$message.error('留言失败，请重试');
                        }
                    }).catch(err=>{
                        this.$message({
                            message: '你的留言没有成功保存，请稍后重试',
                            type: 'warning'
                        });
                    })
                }else if(this.saved){
                    this.$message.error('你已经留言过了,请稍后再试');
                    setTimeout(()=>{
                        this.saved = false;
                    }, 1000);
                }else{
                    this.$message.error('你的留言至少应超过4个字符');
                }
            },
            tishi(){
                this.$message({
                    message: '注意哦，你发表的留言会在缓存中保留，稍后会写入数据库',
                    type: 'warning'
                });
            }
        }
    }
</script>

<style scoped>
    .message{
        padding: 30px 10px;
    }
    .message .box {
        margin-top: 70px;
    }
    .messageBox{
        max-width: 960px;
        height: auto;
        max-height: 480px;
        overflow-y: auto;
        margin: 15px auto 0;
        text-align: left;
        padding: 16px 10px;
        box-shadow: 0 0 10px 2px var(--post-box);
        background-color: var(--post-background);
    }
    .messageBox #mes{
        padding: 6px 10px;
        word-break: break-all;
        word-wrap: break-word;
    }
    .yourmes{
        max-width: 640px;
        margin: 20px auto 0;
    }
    .btnSend{
        margin-top: 20px;
    }
    .message /deep/ .el-textarea__inner:focus{
        border-color: var(--border-color);
    }
    .message /deep/ .el-textarea__inner {
        background-color: var(--post-background);
        border-color: var(--border-color);
    }
    .message /deep/ .el-textarea .el-input__count {
        background: transparent;
    }
</style>
<style>
    @import "../custom/custom.css";
</style>