<template>
    <div class="bottom">
        <p>
            <el-icon class="el-icon-lollipop"></el-icon><a style="color: var(--bottom-icon-color);font-weight: bold;margin-right: 8px" :href=custom.bottom_url>{{custom.bottom_tag}}</a>
            <el-icon class="el-icon-coffee-cup"></el-icon><a style="color: var(--bottom-icon-color);font-weight: bold" :href=custom.bottom_url2>{{custom.bottom_tag2}}</a>
            <br><span id="bottom-span" style="font-size: 12px;">{{custom.bottom_span}}</span>
        </p>
    </div>
</template>

<script>
    import customData from "../custom/custom";
    export default {
        name: "bottom_banner",
        data(){
            return{
                custom: customData,
            }
        }
    }
</script>

<style scoped>
    .bottom{
        margin-top: 6px;
        font-family: "DejaVu Sans Mono","Segoe UI",Monaco,monospace;
    }
    #bottom-span {
        color: var(--bottom-color);
    }
</style>
<style>
    @import "../custom/custom.css";
</style>