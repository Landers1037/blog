<template>
    <div class="zhuanlan-card">
        <p style="padding: 10px 4px">操作专栏 你可以任意创建和修改专栏</p>
        <el-button @click="show = true" style="margin-bottom: 10px">新建专栏</el-button>
        <el-dialog
                title="新建专栏"
                :visible.sync="show"
                width="80%"
                :before-close="handleClose">
            <el-input placeholder="请输入内容" v-model="name" style="margin-bottom: 10px">
                <template slot="prepend">别名</template>
            </el-input>
            <el-input placeholder="请输入内容" v-model="title" style="margin-bottom: 10px">
                <template slot="prepend">标题</template>
            </el-input>
            <el-input placeholder="请输入内容" v-model="posts" style="margin-bottom: 10px">
                <template slot="prepend">列表</template>
            </el-input>
            <el-input placeholder="请输入内容" v-model="content" style="margin-bottom: 10px">
                <template slot="prepend">描述</template>
            </el-input>
            <span slot="footer" class="dialog-footer">
            <el-button @click="show = false">取 消</el-button>
            <el-button type="primary" @click="add_zhuanlan">新 建</el-button>
          </span>
        </el-dialog>
        <el-table
                :data="zlist"
                max-height="calc(100vh - 200px)"
                border
                style="width: 100%;overflow-y: auto">
            <el-table-column
                    prop="primary_id"
                    label="ID"
                    sortable
                    width="60"
            >
            </el-table-column>
            <el-table-column
                    label="别名"
                    sortable
                    width="80">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.name" placeholder="请输入内容"></el-input>
                </template>
            </el-table-column>
            <el-table-column
                    sortable
                    width="120px"
                    label="标题">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.title" placeholder="请输入内容"></el-input>
                </template>
            </el-table-column>
            <el-table-column
                    width="200px"
                    label="列表">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.posts" placeholder="请输入内容"></el-input>
                </template>
            </el-table-column>
            <el-table-column
                    label="描述">
                <template slot-scope="scope">
                    <el-input v-model="scope.row.content" placeholder="请输入内容"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            @click="handleUpdate(scope.row, scope.row.primary_id)">更新</el-button>
                    <el-button
                            size="mini"
                            type="danger"
                            @click="handleDelete(scope.row, scope.row.primary_id)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";
    export default {
        name: "zhuanlan_card",
        data(){
            return{
                show: false,
                zlist: [],
                id: 0,
                name: "",
                title: "",
                content: "",
                posts: ""
            }
        },
        mounted() {
            this.get_zhuanlan();
        },
        methods: {
            get_zhuanlan(){
                this.$http.get(api_dash.zhuanlan).then(res => {
                    if(res.data.data === "fail") {
                        this.$message.error("获取专栏失败");
                    }else {
                        this.zlist = res.data.data;
                    }
                })
            },
            handleClose(){
                // 清空全部缓存
                this.name = "";
                this.title = "";
                this.posts = "";
                this.content = "";
            },
            handleUpdate(row, id){
                if (row.posts.trim() === "") {
                    this.$message.error("指定文章列表为空");
                    this.show = false;
                }
                else if (row.title.trim() === "") {
                    this.$message.error("指定标题为空");
                    this.show = false;
                }
                else {
                    let data = {
                        "id": id,
                        "name": row.name,
                        "title": row.title,
                        "posts": this.format_posts(row.posts),
                        "content": row.content,
                        "date": row.date,
                    }
                    this.$http.put(api_dash.zhuanlan, data).then(res => {
                        if (res.data.data === "success"){
                            this.$message("专栏更新完毕");
                            this.get_zhuanlan();
                        } else {
                            this.$message.error("专栏更新失败");
                        }
                    });
                }
            },
            handleDelete(row, id){
                let data = {
                    "id": id,
                }
                this.$http.delete(api_dash.zhuanlan, {data: data}).then(res => {
                    if (res.data.data === "success"){
                        this.$message.success("专栏删除完毕");
                        this.get_zhuanlan();
                    } else {
                        this.$message.error("专栏删除失败");
                    }
                });
            },
            add_zhuanlan(){
                if (this.posts.trim() === "") {
                    this.$message("指定文章列表为空");
                    this.show = false;
                }
                else if (this.title.trim() === "") {
                    this.$message("指定标题为空");
                    this.show = false;
                }
                else {
                    let data = {
                        "id": 0,
                        "name": this.name,
                        "title": this.title,
                        "posts": this.format_posts(this.posts),
                        "content": this.content,
                        "date": ""
                    }
                    this.$http.put(api_dash.zhuanlan, data).then(res => {
                       if (res.data.data === "success"){
                           this.$message.success("专栏新建完毕");
                           this.show = false;
                           this.get_zhuanlan();
                       } else {
                           this.$message.error("专栏新建失败");
                           this.show = false;
                       }
                    });
                }
            },
            format_posts(posts) {
                // 中文逗号转换
                let res = "";
                res = posts.replace(/，/g, ",");
                return res;
            }
        }
    }
</script>

<style scoped>

</style>