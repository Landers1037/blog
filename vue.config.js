module.exports = {
  publicPath: "/",
  outputDir: "dist",
  productionSourceMap: false,
  filenameHashing: false,
  devServer: {
    port: 8080,
    proxy: {
      '/api': {
        //target: 'https://blog.renj.io',
        target: '',
        changeOrigin: true,
        ws: false,
        pathRewrite: {
          '^/api': '/api'
        },
      },
    },
  },
  configureWebpack: config =>{
    // config.entry.app = ["@babel/polyfill","./src/main.js"];
    config.externals = {
      'vue': 'Vue',
      'element-ui': 'ELEMENT',
      'axios': 'axios',
      'vue-router': 'VueRouter',
      'marked': 'marked',
    }
  },
};