import Vue from 'vue'

const POST_NEW_URL = '/api/posts/create'
const POST_EDIT_URL = '/api/posts/edit/'
const GET_MINE_URL = '/api/posts/my/all'
const GET_BY_SLUG_URL = '/api/posts/single/'
const GET_PUBLISHED_URL = '/api/posts/all'

export default {
  new (post) {
    return Vue.http.post(POST_NEW_URL, post, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'X-XSRF-TOKEN': Vue.cookie.get('XSRF-TOKEN')
      },
      emulateJSON: true,
      emulateHTTP: true
    })
  },

  update (post) {
    return Vue.http.post(POST_EDIT_URL + post.slug, post, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'X-XSRF-TOKEN': Vue.cookie.get('XSRF-TOKEN')
      },
      emulateJSON: true,
      emulateHTTP: true
    })
  },

  getPublished () {
    return Vue.http.get(GET_PUBLISHED_URL)
  },

  getBySlug (slug) {
    return Vue.http.get(GET_BY_SLUG_URL + slug)
  },

  getMine () {
    return Vue.http.get(GET_MINE_URL)
  }
}
