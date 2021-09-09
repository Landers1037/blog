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
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Heiti SC", 'Noto Sans SC', "Source Han Sans SC","Microsoft YaHei", "微软雅黑", sans-serif;
  text-align: center;
  color: var(--text-color);
  position: relative;
  background-color: var(--body-background)!important;
    overflow-y: auto;
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
