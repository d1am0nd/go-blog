import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    user: {}
  },
  getters: {
    user (state) {
      return state.user
    },
    loggedIn (state) {
      return !(Object.keys(state.user).length === 0 && state.user.constructor === Object)
    },
    token (state) {
      return state.token
    }
  },
  mutations: {
    login (state, user) {
      state.user = user
    },
    logout (state) {
      state.user = {}
    }
  }
})
