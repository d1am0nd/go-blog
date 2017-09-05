import Vue from 'vue'
import store from '@/store'
import Errors from '@/errors'
import posts from '@/services/db/posts'

const LOGIN_URL = '/api/users/login'
// const REGISTER_URL = '/api/users/register'
const LOGOUT_URL = '/api/users/logout'
const REFRESH_URL = '/api/users/current'

const USER_SESS = 'user'
const TOKEN_SESS = 'bearer'

export default {
  token: '',
  init () {
    if (sessionStorage.getItem(USER_SESS) !== null && sessionStorage.getItem(TOKEN_SESS) !== null) {
      store.commit('login', JSON.parse(sessionStorage.getItem(USER_SESS)), sessionStorage.getItem(TOKEN_SESS))
      posts.getMine()
      .then((res) => {
        store.commit('setPosts', res.body)
      })
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

      posts.getMine()
      .then((res) => {
        store.commit('setPosts', res.body)
      })
    })
    .catch((err) => {
      Errors.newErrRes(err)
    })
  },
  /*
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
  */

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
    console.log('refresh')
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
        this.token = token
        store.commit('login', user)
        sessionStorage.setItem(USER_SESS, JSON.stringify(user))
        sessionStorage.setItem(TOKEN_SESS, token)
      } else {
        this.refresh(token)
      }
    } else if (check === false) {
      this.token = null
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
    return store.getters.loggedIn
  }
}
