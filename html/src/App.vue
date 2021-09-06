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
  *{
    margin: 0;
    padding: 0;
  }
  html, body {
      height: 100%!important;
      background-color: #fafafa;
  }
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Heiti SC", 'Noto Sans SC', "Source Han Sans SC","Microsoft YaHei", "微软雅黑", sans-serif;
  text-align: center;
  color: #2c3e50;
  position: relative;
  background-color: #fafafa!important;
}
  a{
    text-decoration: none;
    color: #42b983;
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
    background-color: #378de5;
    border-radius: 16px;
  }
  pre::-webkit-scrollbar{
      width: 4px;
      height: 4px;
  }
  pre::-webkit-scrollbar-thumb{
      background-color: #378de5;
      border-radius: 10px;
  }
</style>
