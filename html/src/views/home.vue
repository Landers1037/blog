<template>
    <div class="home">
    <top_banner></top_banner>
        <div class="drawer">
            <el-button size="small" class="el-icon-s-unfold" @click="drawer=true"></el-button>
        </div>
        <el-drawer
                title="🔨"
                :visible.sync="drawer"
                direction="ltr"
                size="64%"
                :modal="false"
                :with-header="false">
            <div class="drawer-content">
                <div class="img">
                    <img src="../assets/avatar.jpg" @click="dashboard">
                    <div class="me">
                        <p style="font-weight: bold"><span class="label">Author</span><span class="link">{{custom.author}}</span></p>
                        <p style="font-weight: bold"><span class="label">My site</span><a class="link" :href=custom.site_url>{{custom.site_name}}</a></p>
                        <p style="font-weight: bold"><span class="label">Github</span><a class="link" :href="'https://github.com/' + custom.github">{{custom.github}}</a></p>
                    </div>
                    <div class="small-bt">
                        <el-tooltip content="搜索文章" placement="bottom">
                            <i class="el-icon-search" @click="open('se')"></i>
                        </el-tooltip>
                        <el-tooltip content="我的邮箱" placement="bottom">
                            <i class="el-icon-message" @click="open('mail')"></i>
                        </el-tooltip>
                        <el-tooltip content="赞助我" placement="bottom">
                            <i class="el-icon-present" @click="open('pay')"></i>
                        </el-tooltip>
                        <el-tooltip content="给我留言" placement="bottom">
                            <i class="el-icon-chat-dot-round" @click="toMessage"></i>
                        </el-tooltip>
                    </div>
                </div>
                <div style="margin-top: 15px;padding: 10px">
                    <el-collapse>
                        <el-collapse-item title="我的项目" name="1">
                            <div style="text-align: left">{{custom.project_des}}<a :href=custom.project_url>{{custom.project}}</a></div>
                        </el-collapse-item>
                    </el-collapse>
                </div>
                <div class="more">
                    <el-button plain class="bt" @click="$router.push('/archive')">归档</el-button>
                    <el-button plain class="bt" @click="$router.push('/about')">关于</el-button>
                    <el-button plain class="bt" @click="status">状态</el-button>
                    <el-button type="primary" class="bt" @click="$router.push('/zhuanlan')">专栏</el-button>
                    <el-tooltip content="查看站点总览" placement="bottom">
                        <i class="el-icon-info" style="margin-top: 15px;cursor: pointer;font-size: 14px;color: #9f9f9f;font-weight: bold" @click="overview">OVERVIEW</i>
                    </el-tooltip>
                    <p style="color: #9f9f9f;text-align: left;padding-left: 10px;margin-top: 15px;font-size: 14px">博客文章的版权归作者所有，转载时请注明来源</p>
                </div>
            </div>
        </el-drawer>
        <div class="wrapper">
            <div class="contents">
                <div class="right animated slideInRight">
                    <div class="img">
                        <img src="../assets/avatar.jpg" @click="dashboard">
                        <div class="me">
                            <p style="font-weight: bold"><span class="label">Author</span><span class="link">{{custom.author}}</span></p>
                            <p style="font-weight: bold"><span class="label">My site</span><a class="link" :href=custom.site_url>{{custom.site_name}}</a></p>
                            <p style="font-weight: bold"><span class="label">Github</span><a class="link" :href="'https://github.com/' + custom.github">{{custom.github}}</a></p>
                        </div>
                        <div class="small-bt">
                            <el-tooltip content="搜索文章" placement="bottom">
                                <i class="el-icon-search" @click="open('se')"></i>
                            </el-tooltip>
                            <el-tooltip content="我的邮箱" placement="bottom">
                                <i class="el-icon-message" @click="open('mail')"></i>
                            </el-tooltip>
                            <el-tooltip content="赞助我" placement="bottom">
                                <i class="el-icon-present" @click="open('pay')"></i>
                            </el-tooltip>
                            <el-tooltip content="给我留言" placement="bottom">
                                <i class="el-icon-chat-dot-round" @click="toMessage"></i>
                            </el-tooltip>
                        </div>
                    </div>
                    <div style="margin-top: 15px">
                        <el-collapse>
                            <el-collapse-item title="我的项目" name="1">
                                <div style="text-align: left">{{custom.project_des}}<a :href=custom.project_url>{{custom.project}}</a></div>
                            </el-collapse-item>
                            <el-collapse-item title="更多标签" name="2">
                                    <div class="tags">
                                        <el-tag
                                                v-for="item in tags_more"
                                                :key="item"
                                                type="success"
                                                effect="dark"
                                                @click="getbytag(item)">
                                            {{ item}}
                                        </el-tag>
                                    </div>
                            </el-collapse-item>
                        </el-collapse>
                    </div>
                    <div class="tags">
                        <el-tag
                                v-for="item in tags_less"
                                :key="item"
                                type="success"
                                effect="dark"
                                @click="getbytag(item)">
                            {{ item }}
                        </el-tag>
                    </div>
                    <div class="more">
                        <el-button plain class="bt" @click="$router.push('/archive')">归档</el-button>
                        <el-button plain class="bt" @click="$router.push('/about')">关于</el-button>
                        <el-button plain class="bt" @click="status">状态</el-button>
                        <el-button type="primary" class="bt" @click="$router.push('/zhuanlan')">专栏</el-button>
                        <el-tooltip content="查看站点总览" placement="bottom">
                            <i class="el-icon-info" style="margin-top: 15px;cursor: pointer;font-size: 14px;color: #9f9f9f;font-weight: bold" @click="overview">OVERVIEW</i>
                        </el-tooltip>
                        <p style="color: #9f9f9f;text-align: left;padding-left: 10px;margin-top: 10px;font-size: 14px">博客文章的版权归作者所有，转载时请注明来源</p>
                    </div>
                </div>
                <div class="left">
                    <div class="articlelists">
                        <el-skeleton :rows="6" animated v-if="posts.list.length === 0" />
                        <div v-for="a in posts.list" :key="a.title" class="post animated slideInDown">
                            <div style="position:relative;">
                                <a class="post-a" :href="'/p/'+a.name">{{a.title}}</a>
                                <span class="post-date" v-if="a.date.indexOf('-')!==-1">{{a.date}}</span>
                            </div>
                            <div class="markdown-body abstract" v-html="mk(a.abstract)"></div>
                            <div class="post-tag" v-if="a.tags && a.tags !== '暂时没有标签'">
                                <el-tooltip v-for="t in tags_to_list(a.tags)"
                                            :key="t"
                                            effect="dark"
                                            :content="'标签: ' + t"
                                            :enterable="false"
                                            placement="bottom-start">
                                    <el-tag
                                        type="info"
                                        size="small"
                                        style="cursor: pointer;margin-right: 8px"
                                        @click="$router.push('/t/' + t)"
                                    >{{t}}</el-tag>
                                </el-tooltip>
                            </div>
                        </div>
                    </div>
                    <div class="pagenation">
                        <el-pagination
                                background
                                :current-page=1
                                :pager-count="pagecount"
                                :page-size=8
                                layout="prev, pager, next"
                                @current-change="reloadpage"
                                :total="posts.total">
                        </el-pagination>
                    </div>
                </div>
            </div>
            <div style="clear: both"></div>
        </div>
        <el-dialog
                title="博客状态"
                :visible.sync="dialogVisible"
                class="sys">
            <div style="width: 100%;margin: 0 auto">
                <span style="width: 120px;display: inline-block;margin-right: 10px">当前的routine协程</span>
                <el-input
                        v-model="routine"
                        :disabled="true"
                         style="margin-bottom: 10px;max-width: 140px">
                </el-input><br>
                <span style="width: 120px;display: inline-block;margin-right: 10px">本站的访问量</span>
                <el-input
                        v-model="uv"
                        :disabled="true"
                        style="max-width: 140px;margin-bottom: 10px">
                </el-input><br>
                <span style="width: 120px;display: inline-block;margin-right: 10px">当天的请求次数</span>
                <el-input
                        v-model="count"
                        :disabled="true"
                        style="max-width: 140px;margin-bottom: 10px">
                </el-input><br>
                <span style="width: 120px;display: inline-block;margin-right: 10px">站点运行时长</span>
                <el-input
                        v-model="days"
                        :disabled="true"
                        style="max-width: 140px;margin-bottom: 10px">
                </el-input>
            </div>

            <span slot="footer" class="dialog-footer">
                <strong style="font-size: 14px;color: #9f9f9f">@{{custom.author}} {{custom.start_year}}-{{year}}</strong>
            </span>
        </el-dialog>
        <el-dialog
                id="user-login"
                title="登录到后台"
                :visible.sync="login"
                width="50%"
                >
            <el-input
                    placeholder="请输入用户名"
                    v-model="admin_name"
                    clearable>
            </el-input>
            <p style="margin-top: 10px"></p>
            <el-input placeholder="请输入密码" v-model="admin_passwd" show-password></el-input>
            <span slot="footer" class="dialog-footer">
                <el-button @click="login = false" style="margin-right: 20px">取 消</el-button>
                <el-button type="primary" @click="loginto">登 录</el-button>
            </span>
        </el-dialog>
        <bottom_banner></bottom_banner>
    </div>
</template>

<script>
    // import marked from 'marked';
    import customData from "../custom/custom";
    import api_article from "../api/article";
    import api_statistic from "../api/statistic";
    import api_dash from "../api/dashboard";
    import api_tags from "../api/tag";
    import pay from "../assets/pay.jpg";
    import Top_banner from "../components/top_banner";
    import Bottom_banner from "../components/bottom_banner";
    export default {
        name: "home",
        components: {Bottom_banner, Top_banner},
        data(){
            return{
                custom: customData,
                drawer: false,
                tags:[],
                tags_more: [],
                tags_less: [],
                posts: {
                    total: null,
                    list: []
                },
                pagecount: 9,
                swidth: 0,
                //sys
                routine: 0,
                uv: 0,
                count: 0,
                days: 0,
                year: 2017,
                dialogVisible: false,
                dashboard_count: customData.dashboard_count,
                login: false,
                admin_name: "",
                admin_passwd: ""
            }
        },
        watch:{
            swidth: function () {
                if(this.swidth<=540 && this.swidth>=460){
                    this.pagecount = 7
                }else if(this.swidth<460){
                    this.pagecount = 5
                }else if(this.swidth>540){
                    this.pagecount = 9
                }
            },
        },
        created(){
            let w = document.body.clientWidth;
            if(w<=540 && w>=460){
                this.pagecount = 7
            }else if(w<460){
                this.pagecount = 5
            }else if(w>540){
                this.pagecount = 9
            }
        },
        mounted() {
            this.loadpage(1);
            this.gettags();
            let _this = this;
            window.onresize = function () {
                _this.swidth = document.body.clientWidth;
            };
            this.re();
            _this.whatdays();
        },
        methods:{
            back(){
                //this.$router.push("/");//这个方法会有问题
                location.reload();
            },
            tags_to_list(tags){
                return tags.split(" ");
            },
            gettags(){
                let _this = this;
                _this.$http.get(api_tags.api_tags_all).then(res=>{
                    let tags = res.data.data;
                    //去重
                    if (tags) {
                      let ifcheck = {};
                      for(var i=0;i<tags.length;i++){
                        if(!ifcheck.hasOwnProperty(tags[i]["tag"])){
                          _this.tags.push(tags[i]["tag"]);
                          ifcheck[tags[i]["tag"]] = true;
                        }
                      }
                      _this.tags_less = _this.tags.slice(0,21);
                      _this.tags_more = _this.tags.slice(21,);
                    }
                }).catch(err=>{
                    this.$message.error('出现错误了，请求标签失败');
                })
            },
            getbytag(tag){
              this.$router.push("/t/"+tag);
            },
            getarticles(){
                this.loadpage(1)
            },
            reloadpage(now){
                this.loadpage(now);
            },
            mk(code) {
                marked.setOptions({
                    renderer: new marked.Renderer(),
                    highlight: function (c) {
                        return hljs.highlightAuto(c).value;
                    },
                    pendantic: false,
                    gfm: true,
                    tables: true,
                    breaks: true,
                    sanitize: false,
                    smartLists: true,
                    xhtml: false
                });
              this.$nextTick(()=>{
                  let markdown_bodys = document.getElementsByClassName("markdown-body");
                  for (let body of markdown_bodys) {
                      let pres = body.getElementsByTagName("pre");
                      for(let i=0;i<pres.length;i++){
                          pres[i].classList.add("hljs");
                      }
                  }
                this.reformat_images();
              })
                return marked(code);
            },
            loadpage(n){
                this.$http.get(api_article.api_article_list,{params:{"p":n}}).then(res=>{
                    this.posts.list = res.data.data;
                    this.posts.total = res.data.len;
                }).catch(err=>{
                    this.$message.error('出现错误了，请求文章失败');
                });
            },
            status(){
                this.dialogVisible = true;
                let _this= this;
                this.$http.get(api_statistic.api_statistic_views).then(res=>{
                   _this.uv = res.data?res.data:0;
                }).catch(err=>{
                    _this.$message.error('出现错误了，请求状态失败');
                });
                this.$http.get(api_statistic.api_statistic_routines).then(res=>{
                    _this.routine= res.data?res.data:0;
                });
                this.$http.get(api_statistic.api_statistic_daily).then(res=>{
                    _this.count = res.data?res.data:0;
                })
            },
            open(type) {
                if(type === "mail") {
                    let text = '<strong>你可以通过我的邮箱联系我 <i>' + this.custom.email + '</i></strong>';
                    this.$alert(text, '联系我', {
                        dangerouslyUseHTMLString: true
                    });
                }else if(type==="pay"){
                    let text = '<strong>觉得不错可以赞助我！</strong><br>' +
                        '<img style="max-width: 140px;margin-top: 10px;vertical-align: middle" src="' + pay +'">'
                        + '<span style="vertical-align: middle;padding-left: 20px;font-size: 18px"> 🥰 Thanks</span>';
                    this.$alert(text, '赞助我', {
                        dangerouslyUseHTMLString: true
                    });
                }else {
                    this.$router.push("/search");
                }
            },
            whatdays(){
                var dateSpan, iDays;
                let Date1 = new Date();
                this.year = Date1.getFullYear();
                Date1 = Date1.toLocaleDateString();
                let Date2 = this.custom.start_date;
                Date1 = Date.parse(Date1);
                Date2 = Date.parse(Date2);
                dateSpan = Date1 - Date2;
                dateSpan = Math.abs(dateSpan);
                iDays = Math.floor(dateSpan / (24 * 3600 * 1000));
                this.days = "🍀 " + iDays + "天";
            },
            toMessage(){
                this.$router.push("/message")
            },
            overview(){
                this.$router.push("/overview")
            },
            dashboard(){
                // 满足指定次数后才会进入隐藏的登陆页面
                if (this.dashboard_count > 0) {
                    this.dashboard_count--;
                }else {
                    if (localStorage.getItem("token")) {
                        this.goto_dashboard();
                    }else{
                        this.login = true;
                    }
                }
            },
            loginto(){
              // 为空时报错
              if (this.admin_name === "" || this.admin_passwd === "") {
                  this.login = false;
                  this.$message("输入的用户名或密码为空");
              }else {
                  this.$http.post(api_dash.login,
                      {"name": this.admin_name, "passwd": this.admin_passwd}).then(res => {
                       if (res.data.data !== "failed") {
                           localStorage.setItem("token", res.data.data);
                           this.goto_dashboard();
                       }else {
                           this.$message.error("登录失败");
                       }
                  }).catch(e => {
                      this.$message.error("登录失败");
                  });
              }
            },
            goto_dashboard(){
                this.$router.push("/dashboard");
            },
            re(){
              this.$nextTick(()=>{
                this.reformat_images();
              })
            },
            // 渲染图片资源 使用lightbox
            // 对于主页的多body情况 使用遍历方案
            // 针对元素位置丢失问题 在创建前先对
            reformat_images(){
              let bodys = document.getElementsByClassName("markdown-body");
              for (let i=0;i<bodys.length;i++) {
                let item = bodys[i];
                // 获取item中的images
                let img = item.getElementsByTagName("img");
                for (let i = 0;i< img.length;i++) {
                  // 重新构造img标签
                  // 是否构造看img是否存在lightbox属性
                  if (!img[i].getAttribute("lightbox")) {
                    // 获取img的父亲
                    let img_parent = img[i].parentNode;
                    let data_img_alt = "images" + i;
                    let href = img[i].src;
                    let light_box_attr = document.createElement("a");
                    img[i].setAttribute("lightbox", "true");
                    light_box_attr.href = href;
                    light_box_attr.classList.add("spotlight");
                    light_box_attr.setAttribute("data-image-alt", data_img_alt);
                    light_box_attr.append(img[i]);
                    img_parent.append(light_box_attr);
                  }
                }
              }
            }
        }
    }
</script>

<style scoped>
    .home{
        user-select: none;
        padding: 30px 10px;
    }
    .wrapper{
        padding: 10px;
        max-width: 980px;
        z-index: 10;
        margin: 70px auto 0;
    }
    .contents{
        position: relative;
        width: 100%;
        margin-top: 30px;
    }
    .contents .right{
        width: 200px;
        padding: 10px 16px;
        border-radius: 2px;
        box-shadow: 0 2px 8px 2px var(--post-box);
        background-color: var(--post-background);
        float: right;
    }
    .right .img img{
        border-radius: 50%;
        max-width: 200px;
        cursor: pointer;
        transition: .8s ease;
    }
    .drawer-content .img img {
        cursor: pointer;
        transition: .8s ease;
    }
    .img img:hover {
        transform: rotateY(180deg);
        transition: 1s ease;
        transition-delay: .1s;
    }
    .img .small-bt{
        padding-top: 15px;
    }
    .small-bt i{
        font-size: 22px;
        padding: 0 10px;
        cursor: pointer;
        transition: .3s ease;
    }
    .small-bt i:hover{
        transition: .3s ease;
        color: red;
    }
    .right .me{
        text-align: left;
        width: 100%;
        margin-top: 24px;
        font-size: 14px;
    }
    .right .tags{
        margin-top: 10px;
        text-align: left;
        padding: 10px;
    }
    .right .tags .el-tag{
        margin-right: 6px;
        margin-top: 5px;
    }
    .right .tags /deep/ .el-tag{
        height: 28px;
        line-height: 26px;
        font-size: 12px;
        cursor: pointer;
    }
    .right .more{
        padding: 10px 0;
    }
    .right .more a{
        cursor: pointer;
    }
    .right .more .bt{
        display: block;
        margin: 6px auto 0;
        padding: 12px 30px;
    }
    .more .bt.el-button {
        background-color: var(--button-background);
        border-color: var(--button-border);
    }
    .more .bt.el-button--primary {
        background-color: var(--button-primary-background);
        border: none;
    }
    /deep/ .el-button+.el-button{
        margin-left: 0;
    }
    .contents .left{
        width: calc(100% - 250px);
        float: left;
    }
    .left .pagenation{
        padding: 20px 0 10px 0;
    }
    @media (max-width: 750px){
        .contents .left{
            width: 100%;
            float: unset;
        }
        .contents .right{
            display: none;
        }
        .wrapper{
            padding: 0;
        }
    }
    .drawer{
        display: none;
    }
    .drawer .el-button:focus, .el-button:hover {
        background-color: var(--button-background);
        border-color: var(--button-border);
    }
    .drawer-content .bt{
        display: block;
        margin: 6px auto 0;
        padding: 12px 30px;
    }
    .drawer-content .img img{
        border-radius: 10px;
        max-width: 85%;
        margin-top: 20px;
        margin-bottom: 10px;
    }
    .drawer-content .me{
        text-align: left;
        width: 85%;
        margin: 0 auto;
        font-size: 14px;
    }
    .el-drawer__wrapper /deep/ .el-drawer__body::-webkit-scrollbar {
        display: none;
    }
    .home /deep/ .el-drawer.ltr {
        background-color: var(--post-background);
        padding: 4px;
    }
    .me p {
        margin-bottom: .6rem;
    }
    .me .label{
        background-color: var(--link-label);
        padding: 8px;
        display: inline-block;
        width: 60px;
        text-align: center;
    }
    .me .link {
        padding: 8px 10px;
        display: inline-block;
        width: calc(100% - 100px);
        background-color: var(--link-background);
    }
    /deep/ .el-input.is-disabled .el-input__inner{
        color: #ff5d65;
    }
    .sys /deep/ .el-dialog{
        width: 60%;
        max-width: 420px;
    }
    @media (max-width: 460px) {
        .drawer{
            display: inline-block;
            position: fixed;
            left: 10px;
            top: 16px;
            z-index: 999;
        }
        .sys /deep/ .el-dialog{
            width: 85%;
        }
    }
    @media (max-width: 360px) {
        .pagenation /deep/ .el-pagination .btn-prev,.pagenation /deep/ .el-pagination .btn-next,.pagenation /deep/ .el-pagination .el-pager li{
            min-width: 24px;
        }
        .sys /deep/ .el-dialog{
            width: 95%;
        }
    }
    .more /deep/ .el-button:focus, .more /deep/ .el-button:hover{
        border-color: var(--button-border);
    }
    @media (max-width: 640px) {
        #user-login /deep/ .el-dialog {
            width: 80%!important;
        }
    }
    @media (max-width: 500px) {
        #user-login /deep/ .el-dialog  {
            width: 92%!important;
        }
    }
    .home /deep/ .el-collapse-item__header {
        background-color: var(--post-background);
        border-color: var(--border-color);
        color: var(--text-color);
    }
    .home /deep/ .el-collapse-item__content, .home /deep/ .el-collapse-item__wrap {
        background-color: var(--post-background);
        border-color: var(--border-color);
        color: var(--text-color);
    }
    .home /deep/ .el-collapse {
        border-color: var(--border-color);
    }
</style>
<style>
    @import "../custom/custom.css";
</style>