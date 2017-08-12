import Vue from 'vue'
import Errors from '@/errors'

const LOGIN_URL = '/api/users/login'
const REGISTER_URL = '/api/users/register'
const LOGOUT_URL = '/api/users/logout'
const REFRESH_URL = '/api/users/current'

const USER_SESS = 'user'
const TOKEN_SESS = 'bearer'

export default {
  user: {},
  token: '',
  check: false,
  init () {
    if (sessionStorage.getItem(USER_SESS) !== null && sessionStorage.getItem(TOKEN_SESS) !== null) {
      this.token = sessionStorage.getItem(TOKEN_SESS)
      this.user = JSON.parse(sessionStorage.getItem(USER_SESS))
      this.check = true
    }
  },
  login (ctx, creds, redirect) {
    Vue.http.post(LOGIN_URL, creds, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      emulateJSON: true,
      emulateHTTP: true
    })
    .then((res) => {
      var header = res.headers.get('Authorization')
      this.setAuth(true, header, res.body)
    })
    .catch((err) => {
      Errors.newErrRes(err)
    })
  },

  register (ctx, creds, redirect) {
    Vue.http.post(REGISTER_URL, creds, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      emulateJSON: true,
      emulateHTTP: true
    })
    .then((res) => {
      this.auth(res.body, true)
    })
    .catch((err) => {
      Errors.newErrRes(err)
    })
  },

  logout () {
    Vue.http.get(LOGOUT_URL)
    .then((res) => {
      this.setAuth(false)
    })
    .catch((err) => {
      Errors.newErrRes(err)
      this.setAuth(false)
    })
  },

  logoutFront () {
    this.setAuth(false)
  },

  refresh (token = null) {
    var headers = {}
    if (token !== null) {
      this.token = token
      headers.Authorization = token
      sessionStorage.setItem(TOKEN_SESS, token)
    }

    Vue.http.get(REFRESH_URL, {
      headers: headers
    })
    .then((res) => {
      if (res.status === 200) {
        var header = res.headers.get('Authorization')
        this.setAuth(true, header, res.body)
      } else {
        this.setAuth(false)
      }
    })
    .catch((err) => {
      Errors.newErrRes(err)
      this.setAuth(false)
    })
  },

  setAuth (check, token = null, user = null) {
    if (check === true && this.check === false) {
      if (user !== null) {
        this.user = user
        this.token = token
        this.check = true
        sessionStorage.setItem(USER_SESS, JSON.stringify(user))
        sessionStorage.setItem(TOKEN_SESS, token)
      } else {
        this.refresh(token)
      }
    } else if (check === false) {
      this.user = null
      this.token = null
      this.check = false
      sessionStorage.removeItem(USER_SESS)
      sessionStorage.removeItem(TOKEN_SESS)
    }
  },

  getUser () {
    if (this.checkAuth()) {
      return JSON.parse(sessionStorage.getItem(USER_SESS))
    } else {
      return null
    }
  },

  getToken () {
    if (this.checkAuth()) {
      return sessionStorage.getItem(TOKEN_SESS)
    }
    return null
  },

  checkAuth () {
    return this.check === true
  }
}
