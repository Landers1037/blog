import Vue from 'vue'
import './elem/element.js'

import App from './App.vue'
import router from './router'
import httprequest from "./axios";

//从cdn引入js文件

Vue.config.productionTip = false;
Vue.prototype.$http = httprequest;

Vue.use(httprequest);
new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
