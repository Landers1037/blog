<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'app',
}
</script>

<style>
  @import "./assets/github-markdown.min.css";
  @import "./assets/theme.css";
  *{
    margin: 0;
    padding: 0;
  }
  html, body {
      height: 100%!important;
      background-color: var(--body-background);
  }
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Heiti SC", 'Noto Sans SC', "Source Han Sans SC","Microsoft YaHei", "微软雅黑", sans-serif;
  text-align: center;
  color: var(--text-color);
  position: relative;
  background-color: var(--body-background)!important;
}
  a{
    text-decoration: none;
    color: var(--a-color);
  }
  @media (max-width: 480px) {
    #app{
      -webkit-tap-highlight-color: transparent;
    }
  }
  @font-face {
    font-face-name: mo;
    font-family: mo;
    src: url("assets/Monaco.woff2") format("woff");
    font-display: fallback;
    font-style: normal;
  }
  ::-webkit-scrollbar{
    width: 10px;
    height: 10px;
  }
  ::-webkit-scrollbar-thumb{
    background-color: var(--main-color);
    border-radius: 16px;
  }
  pre::-webkit-scrollbar{
      width: 4px;
      height: 4px;
  }
  pre::-webkit-scrollbar-thumb{
      background-color: var(--main-color);
      border-radius: 10px;
  }
  [class*=" el-icon-"], [class^=el-icon-] {
      color: var(--bottom-icon-color);
  }
  /*全局按钮*/
  .post .el-button {
      background-color: var(--button-background);
      border-color: var(--button-border);
  }
  .post .el-button--primary {
      background-color: var(--button-primary-background);
      border-color: var(--button-border);
  }
  /*全局标签*/
  .el-tag.el-tag--info {
      background-color: var(--tag-background);
      color: var(--tag-color);
      border-color: var(--tag-border);
  }
  /*分页*/
  .el-pagination.is-background .btn-next, .el-pagination.is-background .btn-prev, .el-pagination.is-background .el-pager li {
      background-color: var(--pagination-bg);
      color: var(--pagination-color);
  }
  .el-select-dropdown {
      background-color: var(--post-background);
      border: none;
  }
  .el-select-dropdown__item.hover, .el-select-dropdown__item:hover {
      background-color: var(--select-hover);
  }
  .el-divider {
      background-color: var(--banner-divider-background);
  }
  .el-input__inner {
      background-color: var(--body-background);
  }
  /*全局加载*/
  .el-skeleton.is-animated .el-skeleton__item {
      background-image: linear-gradient(90deg, var(--loading-from) 25%, var(--loading-to) 37%, var(--loading-from) 63%);

  }
  /*全局登录菜单*/
  .home .el-dialog,.home .el-pager li {
      background-color: var(--post-background);
  }
  .home .el-dialog__title {
      color: var(--text-color);
  }
  .home .el-input--suffix .el-input__inner {
      background-color: var(--body-background)!important;
  }
  .dashboard .el-menu {
      background-color: var(--body-background);
  }
  .dashboard .el-menu.el-menu--horizontal {
      border-color: var(--border-color);
  }
  .dashboard .el-table th, .el-table tr {
      background-color: var(--post-background);
  }
</style>
