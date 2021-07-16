<template>
    <div class="post-card">
        <h3 style="padding: 10px 4px">注意在此处上传文章包含创建和更新方式</h3>
        <el-row>
            <el-col :span="12">
                <el-upload
                        class="upload"
                        name="uploadmd"
                        drag
                        :headers="insert_token"
                        :action="upload_url"
                        :on-success="success"
                        :on-error="fail">
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                    <div class="el-upload__tip" slot="tip">只能上传md文件</div>
                </el-upload>
            </el-col>
            <el-col :span="12">
                <el-button type="primary" @click="export_all_post">导出全部文章</el-button>
                <el-button type="success">导出数据库</el-button>
                <el-button type="warning" @click="back_up_db">备份数据库</el-button>
            </el-col>
        </el-row>

        <div id="post-list">
            <el-table
                    id="post-list-inner"
                    :data="postData"
                    border
                    style="width: 100%">
                <el-table-column
                        prop="id"
                        label="ID"
                        sortable
                        width="80"
                >
                </el-table-column>
                <el-table-column
                        prop="name"
                        label="URI"
                        sortable
                        width="160">
                </el-table-column>
                <el-table-column
                        prop="title"
                        sortable
                        label="标题">
                </el-table-column>
                <el-table-column label="操作" width="220">
                    <template slot-scope="scope">
                        <el-button
                                size="mini"
                                type="success"
                                @click="handleOpen(scope.$index, scope.row.name)">编辑</el-button>
                        <el-button
                                size="mini"
                                @click="handleEdit(scope.$index, scope.row.name)">打开</el-button>
                        <el-button
                                size="mini"
                                type="danger"
                                @click="handleDelete(scope.$index, scope.row.name)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-dialog
                title="文章更新"
                :visible.sync="post_edit"
                width="80%">
            <el-input style="margin-bottom: 10px" placeholder="请输入内容" v-model="name">
                <template slot="prepend">URI</template>
            </el-input>
            <el-input style="margin-bottom: 10px" placeholder="请输入内容" v-model="newname">
                <template slot="prepend">新URI</template>
            </el-input>
            <el-input style="margin-bottom: 10px" placeholder="请输入内容" v-model="title">
                <template slot="prepend">标题</template>
            </el-input>
            <el-input style="margin-bottom: 10px" placeholder="请输入内容" v-model="date_plus">
                <template slot="prepend">日期</template>
            </el-input>
            <el-input style="margin-bottom: 10px" placeholder="请输入内容" v-model="tags">
                <template slot="prepend">标签</template>
            </el-input>
            <el-switch
                    style="margin-bottom: 20px"
                    v-model="pin"
                    active-text="置顶"
                    inactive-text="取消置顶">
            </el-switch>
            <el-upload
                    name="uploadmd"
                    :action="api_post_update"
                    :limit="1"
                    ref="upload"
                    :http-request="upload_file"
                    :on-error="update_post_err"
                    :on-success="update_post_success">
                <el-button size="small" type="primary">上传更新文件</el-button>
                <div slot="tip" class="el-upload__tip">本方式默认读取md文件内容作为正文内容</div>
            </el-upload>
            <br>
            <span slot="footer" class="dialog-footer">
                <el-button type="info" @click="export_post">导 出</el-button>
                <el-button @click="post_edit = false">取 消</el-button>
                <el-button type="primary" @click="update_post">更 新</el-button>
            </span>
        </el-dialog>
        <el-dialog
                title="文章编辑"
                :visible.sync="post_open"
                width="94%">
            <el-row :gutter="20">
                <el-col :span="8"><el-input
                        style="color: #404040;font-weight: bold;"
                        v-model="post_open_title"
                >
                    <template slot="prepend">标题</template>
                </el-input></el-col>
                <el-col :span="16"><el-input
                        v-model="post_open_tags"
                        style="">
                    <template slot="prepend">标签</template>
                </el-input></el-col>
            </el-row>
            <el-input
                    style="margin-top: 10px"
                    type="textarea"
                    :rows="16"
                    placeholder="正文内容"
                    v-model="post_open_content">
            </el-input>
            <span slot="footer" class="dialog-footer">
                <el-button size="mini" @click="post_open = false">取 消</el-button>
                <el-button size="mini" type="primary" @click="update_post_all">更 新</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";
    import customData from "../custom/custom";
    import api_article from "../api/article";
    export default {
        name: "post_card",
        data(){
            return{
                upload_url: customData.api_url + api_dash.upload_file,
                postData: [],
                post_edit: false,
                post_open: false,
                post_data: {},
                name: "",
                newname: "",
                title: "",
                date_plus: "",
                tags: "",
                pin: false,
                // post open
                post_open_name: "",
                post_open_title: "",
                post_open_tags: "",
                post_open_abstract: "",
                post_open_content: "",
                api_post_update: api_dash.post + "?type=file&name=" + this.name
            }
        },
        mounted() {
            this.getPostsData();
        },
        computed: {
            insert_token(){
                let t = localStorage.getItem("token");
                return {"admin_token": t}
            },
        },
        methods: {
            success(res){
                if(res.data === "fail"){
                    this.$message.error("文章上传失败\n" + res.msg)
                }else {
                    this.$message.success("文章上传成功");
                    this.getPostsData();
                }
            },
            fail(err){
                this.$message.error("文章上传失败");
            },
            getPostsData(){
                this.$http.get(api_dash.post).then(res => {
                    if(res.data.data === "" || res.data.data.length === 0) {
                        this.$message.error("获取文章列表失败");
                    }else {
                        this.postData = res.data.data;
                    }
                });
            },
            handleDelete(index, name){
                let data = {"name" : name};
                this.$http.delete(api_dash.post, {data: data}).then(res => {
                    if(res.data.data === "fail"){
                        this.$message.error("文章" + name + "删除失败");
                    }else {
                        this.$message.success("文章" + name + "删除成功");
                        this.getPostsData();
                    }
                })
            },
            handleOpen(index, name){
                if (name){
                    this.$http(api_dash.post + "?name=" + name).then(res => {
                        if (res.data.data !== "fail") {
                            this.post_open_name = name;
                            this.post_open_title = res.data.data.title;
                            this.post_open_tags = res.data.data.tags;
                            this.post_open_abstract = res.data.data.abstract;
                            this.post_open_content = res.data.data.content;
                            this.post_open = true;
                        }else {
                            this.$message.error("获取文章内容失败");
                        }
                    });
                }
            },
            handleEdit(index, name){
                // 根据name获取文章内容
                this.$http.get(api_dash.post + "?name=" + name).then(res => {
                    if (res.data.data !== ""){
                        let d = res.data.data;
                        this.name = d.name;
                        this.newname = d.name;
                        this.title = d.title;
                        this.date_plus = d.date_plus;
                        this.tags = d.tags;
                        this.pin = d.pin === 1;
                        this.post_edit = true;
                    }else {
                        this.$message.error("获取文章内容失败");
                    }
                });
            },
            clear_data(){
                this.name = this.newname = this.title = this.date_plus = "";
                this.tags = "";
                this.pin  = false;
            },
            upload_file(params){
                let form = new FormData;
                form.append("uploadmd", params.file);
                this.$http.post(api_dash.post + "?type=file&name=" + this.name, form, {headers: {
                    'Content-Type': 'multipart/form-data'
                    }}).then(res => {
                        if (res.data.data === "success"){
                            params.onSuccess("上传成功");
                        }else {
                            params.onError("上传失败");
                        }
                }).catch(() => {
                    params.onError("上传失败");
                });
            },
            update_post(){
                // 文章内容为单独更新
                let type = "args";
                let data = {
                  "name": this.name,
                  "newname": this.newname,
                  "title": this.title,
                  "date": this.date_plus,
                  "tags": this.tags,
                  "pin": this.pin ?1: 0,
                };
                this.$http.put(api_dash.post + "?type=" + type, data).then(res => {
                   if (res.data.data === "success"){
                       this.$message.success("文章更新成功");
                       this.post_edit = false;
                       this.clear_data();
                   } else {
                       this.$message.error("文章更新失败");
                   }
                });
            },
            update_post_all(){
              // 更新带正文的文章 type=editor
              let url = api_dash.post + "?type=editor";
              let data = {
                  "name": this.post_open_name,
                  "title": this.post_open_title,
                  "tags": this.post_open_tags,
                  "content": this.post_open_content,
              };
              this.$http.put(url, data).then(res => {
                 if (res.data.data !== "fail") {
                     this.$message.success("更新文章成功");
                     this.post_open = false;
                     this.post_open_name = "";
                     this.post_open_title = "";
                     this.post_open_tags = "";
                     this.post_open_abstract =  "";
                     this.post_open_content = "";
                 }else {
                     this.$message.error("更新文章失败");
                 }
              });
            },
            update_post_err(){
                this.$message.error("更新正文内容失败");
            },
            update_post_success(){
                this.$message.success("更新正文内容成功");
            },
            export_all_post(){
                this.$http.post(api_dash.export_file).then(res => {
                    if (res.data.data !== "fail") {
                        let zip = new JSZip();
                        for (let p of res.data.data) {
                            let tags = p.tags !== ""?p.tags.split(" ").join(","):"";
                            let cates = p.categories !== ""?p.categories.split(" ").join(","):"";
                            let data = `---
title: ${p.title}
name: ${p.name}
date: ${p.date_plus}
id: 0
tags: [${tags}]
categories: [${cates}]
abstract: ""
---
${p.abstract}
<!--more-->
${p.content}`;
                            zip.file(p.name + ".md", data);
                        }
                        zip.generateAsync({type:"blob"})
                            .then(function(content) {
                                // Force down of the Zip file
                                saveAs(content, "posts.zip");
                            });
                    }else {
                        this.$message.error("获取文章失败");
                    }
                });
            },
            export_post(){
                let query = "?name=" + this.name;
                let file_name = this.name + ".md";
                let file_data = "";
                this.$http.post(api_dash.export_file + query).then(res => {
                   if (res.data.data !== "fail"){
                       let tags = res.data.data.tags !== ""?res.data.data.tags.split(" ").join(","):"";
                       let cates = res.data.data.categories !== ""?res.data.data.categories.split(" ").join(","):"";
                       file_data = `---
title: ${res.data.data.title}
name: ${res.data.data.name}
date: ${res.data.data.date_plus}
id: 0
tags: [${tags}]
categories: [${cates}]
abstract: ""
---
${res.data.data.abstract}
<!--more-->
${res.data.data.content}`;
                       let file = new File([file_data], file_name, {type: "text/plain;charset=utf-8"});
                       saveAs(file);
                   }else {
                       this.$message.error("获取文件内容失败");
                   }
                });
            },
            back_up_db(){
                // 数据库的备份都会在backup目录下
                this.$http.post(api_dash.db_back).then(res => {
                    if (res.data.data === "success") {
                        this.$message.success("数据库备份完毕");
                    }else {
                        this.$message.error("数据库备份失败");
                    }
                })
            }
        }
    }
</script>

<style scoped>
  #post-list {
    margin-top: 20px;
  }
  #post-list #post-list-inner {
    height: calc(100vh - 400px);
    overflow-y: auto;
  }
</style>