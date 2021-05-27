<template>
    <div class="view-card">
        <p style="padding: 10px">访问量统计</p>
<!--        最顶部一定是all-->
        <el-table
                :data="view_list"
                max-height="480"
                border
                style="width: 100%">
            <el-table-column
                    prop="name"
                    label="文章URI"
                    sortable
                    width="100">
            </el-table-column>
            <el-table-column
                    prop="title"
                    sortable
                    label="标题">
            </el-table-column>
            <el-table-column
                    sortable
                    width="100"
                    label="访问量">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.view"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="80">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            type="primary"
                            @click="handleUpdate(scope.$index, scope.row.name, scope.row.view)">更新</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";
    export default {
        name: "View_card",
        data(){
            return{
                view_list: []
            }
        },
        mounted() {
            this.get_views();
        },
        methods: {
            get_views(){
                this.$http.get(api_dash.view).then(res => {
                    if (res.data.data){
                        this.view_list = res.data.data;
                    }else {
                        this.$message.error("获取访问量列表失败");
                    }
                });
            },
            handleUpdate(index, name, count){

                let data = {
                    "name": name,
                    "count": parseInt(count, 10),
                }
                this.$http.put(api_dash.view, data).then(res => {
                   if (res.data.data === "success"){
                       this.$message("更新访问量成功");
                       this.get_views();
                   } else {
                       this.$message.error("更新访问量失败");
                   }
                });
            }
        }
    }
</script>

<style scoped>

</style>