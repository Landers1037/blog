<template>
<!--    作为图表可视化只展示文章总数 访问量-->
    <div class="charts">
        <h2 style="margin-bottom: 10px">统计数据可视化</h2>
        <div id="article-part">
            <h3>文章统计</h3>
            <p class="data-item">总数: {{post_count}}</p>
            <h3>最受欢迎的文章</h3>
            <p class="data-item"><a style="text-decoration: none;color: white" :href="'/p/'+post_most_like">{{post_most_like_name}}</a></p>
        </div>
        <canvas id="chart-view-all" width="540" height="420"></canvas>
        <div id="view-part">
            <h3>访问统计</h3>
            <p class="data-item">总数: {{view_count}}</p>
        </div>

        <div id="comment-part">
            <h3>评论统计</h3>
            <p class="data-item">总数: {{comment_count}}</p>
        </div>

        <div id="like-part">
            <h3>点赞统计</h3>
            <p class="data-item">总数: {{like_count}}</p>
        </div>

        <div id="share-part">
            <h3>分享统计</h3>
            <p class="data-item">总数: {{share_count}}</p>
        </div>
    </div>
</template>

<script>
    import api_dash from "../api/dashboard";

    export default {
        name: "Show_card",
        data(){
            return{
                view_all: "",
                views: [],
                post_count: 0,
                post_most_like: "",
                post_most_like_name: "",
                view_count: 0,
                comment_count: 0,
                like_count: 0,
                share_count: 0
            }
        },
        mounted() {
            this.get_post_count();
            this.get_all_data();
            this.get_view_count();
            this.get_comment_count();
            this.get_like_count();
            this.get_share_count();
        },
        methods: {
            get_all_data(){
                this.$http.get(api_dash.view).then(res => {
                    // 数组重排取最大值为all
                    let sort_data = res.data.data.sort(function (a, b) {
                        return b.view - a.view
                    });
                    this.view_all = sort_data[0];
                    this.view_all = res.data.data;
                    this.views = res.data.data;
                    // max post
                    this.post_most_like_name = this.view_all[1].title;
                    this.post_most_like = this.view_all[1].name;
                    this.draw_all_view();
                });
            },
            draw_all_view(){
                let node = document.getElementById("chart-view-all");
                // gen
                let _labels = [];
                let _datas = [];
                for (let d of this.views) {
                    if (d.name !== "all"){
                        _datas.push(d.view);
                        _labels.push(d.name);
                    }
                }
                if (_datas.length >= 10) {
                    _datas = _datas.slice(0,10);
                    _labels = _labels.slice(0,10);
                }
                let data = {
                    labels: _labels,
                    datasets: [
                        {
                            label: "访问量",
                            fillColor: "rgba(151,187,205,0.5)",
                            strokeColor: "rgba(220,220,220,0.8)",
                            highlightFill: "rgba(151,187,205,0.75)",
                            highlightStroke: "rgba(220,220,220,1)",
                            data: _datas
                        }
                    ]
                };
                new Chart(node, {type: "bar", data: data, options: {
                    responsive: false
                }});
            },
            get_post_count(){
                this.$http.get(api_dash.post).then(res => {
                    this.post_count = res.data.data.length;
                });
            },
            get_comment_count(){
                this.$http.get(api_dash.comment + "?name=all").then(res => {
                    this.comment_count = res.data.data.length;
                });
            },
            get_view_count(){
                let url = api_dash.view + "?name=all";
                this.$http.get(url).then(res => {
                    if (res.data.data){
                        this.view_count = res.data.data;
                    }
                })
            },
            get_like_count(){
                let url = api_dash.like + "?name=all";
                this.$http.get(url).then(res => {
                    if (res.data.data){
                        this.like_count = res.data.data;
                    }
                })
            },
            get_share_count(){
                let url = api_dash.share + "?name=all";
                this.$http.get(url).then(res => {
                    if (res.data.data){
                        this.share_count = res.data.data;
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .charts{
        padding: 10px;
    }
    .data-item{
        background-color: #378de5;
        color: white;
        font-weight: bold;
        width: fit-content;
        padding: 10px 15px;
        font-size: 18px;
        border-radius: 10px;
        margin: 10px 0;
    }
</style>
