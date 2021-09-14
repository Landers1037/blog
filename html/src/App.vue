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
  @import "./assets/common.css";
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
  font-family: Lato, Helvetica, "PingFang SC", "Heiti SC", "Source Han Sans SC", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  text-align: center;
  color: var(--text-color);
  position: relative;
  background-color: var(--body-background)!important;
  overflow-y: auto;
  /*对于所有图片背景去除  */
}

#app img {
    background-color: transparent!important;
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
</style>
