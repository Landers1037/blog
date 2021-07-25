<template>
    <div class="message_card">
        <p style="padding: 10px 4px">管理留言</p>
        <el-table
                :data="message_list"
                max-height="calc(100vh - 200px)"
                border
                style="width: 100%;overflow-y: auto">
            <el-table-column
                    prop="primary_id"
                    label="ID"
                    sortable
                    width="80"
            >
            </el-table-column>
            <el-table-column
                    prop="user"
                    label="用户"
                    sortable
                    width="120">
            </el-table-column>
            <el-table-column
                    prop="date"
                    sortable
                    width="160"
                    label="留言日期">
            </el-table-column>
            <el-table-column
                    sortable
                    label="留言">
                <template slot-scope="scope">
                    <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 6}" clearable v-model="scope.row.message"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button
                        size="mini"
                        type="primary"
                        @click="handleUpdate(scope.$index, scope.row.primary_id, scope.row.message)">更新</el-button>
                    <el-button
                            size="mini"
                            type="danger"
                            @click="handleDelete(scope.$index, scope.row.primary_id)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";
    export default {
        name: "message_card",
        data(){
            return{
                message_list: []
            }
        },
        mounted() {
            this.get_message_list();
        },
        methods: {
            get_message_list(){
                this.$http.get(api_dash.message).then(res => {
                    if (res.data.data === "fail") {
                        this.$message.error("获取留言失败");
                    }else{
                        this.message_list = res.data.data;
                    }
                })
            },
            handleDelete(index, id){
                let data = {"id" : id};
                this.$http.delete(api_dash.message, {data: data}).then(res => {
                    if(res.data.data === "fail"){
                        this.$message.error("留言" + id + "删除失败");
                    }else {
                        this.$message.success("留言" + id + "删除成功");
                        this.get_message_list();
                    }
                })
            },
            handleUpdate(index, id, message){
                let data = {
                    "id": id,
                    "message": message
                }
                this.$http.put(api_dash.message, data).then(res => {
                    if(res.data.data === "fail"){
                        this.$message.error("留言" + id + "更新失败");
                    }else {
                        this.$message.success("留言" + id + "更新成功");
                        this.get_message_list();
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>