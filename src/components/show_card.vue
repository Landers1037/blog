<template>
    <div class="charts">
        <h3>统计数据可视化</h3>
        <canvas id="chart-view-all" width="400" height="400"></canvas>
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
            }
        },
        mounted() {
            this.get_all_data();
        },
        methods: {
            get_all_data(){
                this.$http.get(api_dash.view).then(res => {
                    // 数组重排取最大值为all
                    let sort_data = res.data.data.sort(function (a, b) {
                        return b.view - a.view
                    });
                    this.view_all = sort_data[0];
                    console.log(sort_data);
                    this.view_all = res.data.data;
                    this.views = res.data.data;
                });
            },
            draw_all_view(){

            }
        }
    }
</script>

<style scoped>

</style>

<style scoped>

</style>