<template>
    <div class="comment-card">
        <p style="padding: 10px">文章评论统计</p>
        <el-table
            :data="comments"
            max-height="calc(100vh - 200px)"
            border
            style="width: 100%;overflow-y: auto">
            <el-table-column
                prop="name"
                label="文章URI"
                sortable
                width="200">
            </el-table-column>
            <el-table-column
                prop="user"
                sortable
                width="140"
                label="用户">
            </el-table-column>
            <el-table-column
                sortable
                label="评论">
                <template slot-scope="scope">
                    <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 6}" clearable show-word-limit v-model="scope.row.comment"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button
                        size="mini"
                        type="primary"
                        @click="update_comment(scope.row.name, scope.row.primary_id, scope.row.comment)">更新</el-button>
                    <el-button
                        size="mini"
                        type="danger"
                        @click="del_comment(scope.row.name, scope.row.primary_id)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";

    export default {
        name: "Comment_card",
        data(){
            return {
                comments: []
            }
        },
        mounted() {
            this.get_comments();
        },
        methods: {
            get_comments(){
                this.$http.get(api_dash.comment).then(res => {
                    if (res.data.data) {
                        this.comments = res.data.data;
                    }else {
                        this.$message.error("获取评论失败");
                    }
                });
            },
            del_comment(name, id){
                let data = {
                    "name": name,
                    "id": id,
                };
                this.$http.delete(api_dash.comment, {data: data}).then(res => {
                    if (res.data.data && res.data.data === "success") {
                        this.$message.success("评论删除成功");
                        this.get_comments();
                    }else {
                        this.$message.error("评论删除失败");
                    }
                });
            },
            update_comment(name, id, comment){
                let data = {
                    "name": name,
                    "id": id,
                    "comment": comment
                };
                this.$http.put(api_dash.comment, data).then(res => {
                    if (res.data.data && res.data.data === "success") {
                        this.$message.success("评论更新成功");
                        this.get_comments();
                    }else {
                        this.$message.error("评论更新失败");
                    }
                });
            }
        }
    }
</script>

<style scoped>

</style>