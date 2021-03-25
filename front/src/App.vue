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
  @import "assets/animate.min.css";
  *{
    margin: 0;
    padding: 0;
  }

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Heiti SC","Hiragino Sans GB", "Source Han Sans SC Medium","Droid Sans fallback","Microsoft YaHei", "微软雅黑",Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  padding: 50px 20px 20px 20px;
}
  a{
    text-decoration: none;
    color: #42b983;
  }
  @media (max-width: 480px) {
    #app{
      padding: 40px 6px 20px 6px;
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
    width: 15px;
    height: 10px;
  }
  ::-webkit-scrollbar-thumb{
    background-color: transparent;
  }
  .github {
    @import 'highlight.js/styles/github.css';
  }
  .monokai {
    @import 'highlight.js/styles/monokai.css';
  }
  .solarized {
  @import 'highlight.js/styles/solarized-dark.css';
  }
  .atom {
  @import 'highlight.js/styles/atom-one-dark.css';
  }
  .xcode {
  @import 'highlight.js/styles/xcode.css';
  }
  .xterm {
  @import 'highlight.js/styles/xt256.css';
  }
</style>
