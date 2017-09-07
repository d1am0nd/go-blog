// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import store from '@/store'
import VueResource from 'vue-resource'
import App from '@/App'
import router from '@/router'
import auth from '@/auth/auth'
import 'skeleton-css/css/normalize.css'
import 'skeleton-css/css/skeleton.css'
import './styles/main.scss'

Vue.config.productionTip = false
Vue.use(VueResource)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  data () {
    return {
      auth: auth
    }
  },
  created () {
    auth.init()
  },
  components: { App }
})

Vue.http.interceptors.push((req, next) => {
  if (auth.hasToken()) {
    req.headers.set('Authorization', auth.getToken())
  }
  next((res) => {
    if (res.headers.has('Authorization')) {
      auth.setAuth(true, res.headers.get('Authorization'))
    }
  })
})
