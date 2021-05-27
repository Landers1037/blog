// import Vue from 'vue'
// import './elem/element.js'
//import axios from 'axios'

import App from './App.vue'
import router from './router'
import httprequest from "./axios";

//从cdn引入js文件

Vue.config.productionTip = false;
// axios.defaults.baseURL = '/';
Vue.prototype.$http = httprequest;

Vue.use(httprequest);
new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
