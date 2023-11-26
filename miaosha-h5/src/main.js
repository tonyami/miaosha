import Vue from 'vue'
import App from './App'
import router from './router'
import fastclick from 'fastclick'
import api from './utils/api'

import './assets/base.css'

fastclick.attach(document.body)
Vue.prototype.$api = api

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
