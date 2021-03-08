<template>
    <div class="message">
        <div class="header">
            <div class="animated slideInDown">
                <label id="title" @click="back">Landers 1037</label>
            </div>
            <el-divider><span style="font-family: 'DejaVu Sans Mono','Source Code Pro','Liberation Mono',monospace;font-size: 14px">Never stop debugging</span></el-divider>
        </div>
        <div class="box">
            <p>在这里写下你的留言</p>
            <div class="messageBox">
                <ul v-for="m in messages"id="mes" class="animated fadeInDown"><i style="font-size: 20px;padding-right: 6px" class="el-icon-chat-round"></i>{{m}}</ul>
            </div>
            <el-input
                    type="textarea"
                    :rows="2"
                    class="yourmes"
                    placeholder="请输入留言"
                    autosize
                    v-model="textarea">
            </el-input>
        </div>
        <el-button class="btnSend" icon="el-icon-chat-square" @click="sendmes">发表留言</el-button>
    </div>
</template>

<script>
    export default {
        name: "message",
        data(){
            return{
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
                this.$http.get("/api/message").then(res=>{
                    this.messages = res.data;
                })
            },
            sendmes(){
                let data = {};
                let _this = this;
                if(this.textarea.trim().length>=4 && !_this.saved){
                    data = {"mes":this.textarea};
                    this.$http.post("/api/message",data).then(res=>{
                        if(res.data === "saved"){
                            _this.tishi();
                            _this.saved = true;
                            setTimeout(function () {
                                _this.getmes();
                            },1000);
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
                    this.$message.error('你已经留言过了');
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
        padding: 4px;
    }
    #title{
        background-color: #363636;
        color: white;
        padding: 8px 10px;
        cursor: pointer;
        font-family: mo,monospace;
    }
    /deep/ .el-divider.el-divider--horizontal{
        margin: 36px 0 24px 0;
    }
    /deep/ .el-divider__text.is-center{
        display: block;
        width: fit-content;
    }
    .messageBox{
        max-width: 960px;
        height: 400px;
        overflow-y: auto;
        margin: 15px auto 0;
        text-align: left;
        padding-top: 6px;
        box-shadow: 1px 1px 4px #9f9f9f;
    }
    .messageBox #mes{
        padding: 6px 10px;
        word-break: break-all;
        word-wrap: break-spaces;
    }
    .yourmes{
        max-width: 640px;
        margin: 20px auto 0;
    }
    .btnSend{
        margin-top: 20px;
    }
    .message /deep/ .el-textarea__inner:focus{
        border-color: #808080;
    }
</style>