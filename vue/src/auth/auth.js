import Vue from 'vue'
import Errors from '@/errors'
import router from '@/router'

const LOGIN_URL = '/api/users/login'
const REGISTER_URL = '/api/users/register'
const LOGOUT_URL = '/api/users/logout'
const REFRESH_URL = '/api/users/current'

export default {
  user: {},
  check: true,
  login (ctx, creds, redirect) {
    Vue.http.post(LOGIN_URL, creds, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'X-XSRF-TOKEN': Vue.cookie.get('XSRF-TOKEN')
      },
      emulateJSON: true,
      emulateHTTP: true
    })
    .then((res) => {
      this.auth(res.body, true)
      router.push({ name: 'home' })
    })
    .catch((err) => {
      Errors.newErrRes(err)
      sessionStorage.removeItem('user')
      this.auth({}, false)
      router.push({ name: 'home' })
    })
  },

  register (ctx, creds, redirect) {
    Vue.http.post(REGISTER_URL, creds, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'X-XSRF-TOKEN': Vue.cookie.get('XSRF-TOKEN')
      },
      emulateJSON: true,
      emulateHTTP: true
    })
    .then((res) => {
      this.auth(res.body, true)
      router.push({ name: 'home' })
    })
    .catch((err) => {
      Errors.newErrRes(err)
      sessionStorage.removeItem('user')
      this.auth({}, false)
      router.push({ name: 'home' })
    })
  },

  logout () {
    Vue.http.get(LOGOUT_URL)
    .then((res) => {
      sessionStorage.removeItem('user')
      this.auth({}, false)
      router.push({ name: 'home' })
    })
    .catch((err) => {
      Errors.newErrRes(err)
      sessionStorage.removeItem('user')
      this.auth({}, false)
      router.push({ name: 'home' })
    })
  },

  logoutFront () {
    sessionStorage.removeItem('user')
    this.user = null
    router.push({ name: 'home' })
  },

  refresh () {
    Vue.http.get(REFRESH_URL)
    .then((res) => {
      this.auth(res.body, true)
    })
    .catch((err) => {
      sessionStorage.removeItem('user')
      this.auth({}, false)
      Errors.newErrRes(err)
    })
  },

  getUser () {
    if (this.checkAuth()) {
      return JSON.parse(sessionStorage.getItem('user'))
    } else {
      return null
    }
  },

  checkAuth () {
    return sessionStorage.getItem('user') !== null
  },

  auth (user, check) {
    if (check) {
      sessionStorage.setItem('user', user)
      this.user = JSON.parse(user)
    } else {
      sessionStorage.removeItem('user')
      this.user = {}
    }
    this.check = check
  }
}
