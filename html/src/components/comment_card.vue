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
                width="100">
            </el-table-column>
            <el-table-column
                prop="user"
                sortable
                width="160"
                label="用户">
            </el-table-column>
            <el-table-column
                sortable
                label="评论">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.comment"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="80">
                <template slot-scope="scope">
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
            }
        }
    }
</script>

<style scoped>

</style>