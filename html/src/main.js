// import Vue from 'vue'
// import './elem/element.js'
//import axios from 'axios'

import App from './App.vue'
import router from './router'
import httprequest from "./axios";
import customData from "./custom/custom";
import { MessageBox,Message } from 'element-ui';
//从cdn引入js文件

Vue.config.productionTip = false;
// axios.defaults.baseURL = '/';
Vue.prototype.$http = httprequest;

Vue.use(httprequest);

// 覆盖messagebox
Vue.prototype.$message = function (msg) {
  Message(msg)
}
Vue.prototype.$message = function(msg){
  return Message({
    message: msg,
    duration: customData.message_duration
  })

}
Vue.prototype.$message.success = function (msg) {
  return Message.success({
    message: msg,
    duration: customData.message_duration
  })
}
Vue.prototype.$message.warning = function (msg) {
  return Message.warning({
    message: msg,
    duration: customData.message_duration
  })
}

Vue.prototype.$message.error = function (msg) {
  return Message.error({
    message: msg,
    duration: customData.message_duration
  })
}
Vue.prototype.$message.info = function (msg) {
  return Message.info({
    message: msg,
    duration: customData.message_duration
  })
}

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
