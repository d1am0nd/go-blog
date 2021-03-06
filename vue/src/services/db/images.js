import Vue from 'vue'

const POST_NEW_URL = '/api/images/create'
const POST_EDIT_URL = '/api/images/edit/'
const GET_IMAGES_URL = '/api/images/all'
const GET_BY_ID_URL = '/api/images/single/'
const GET_DELETE_BY_ID = '/api/images/delete/'
/*
const GET_MINE_URL = '/api/images/my/all'
const GET_PUBLISHED_URL = '/api/images/all'
*/
export default {
  new (image) {
    return Vue.http.post(POST_NEW_URL, image, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  },
  update (image, id) {
    return Vue.http.post(POST_EDIT_URL + id, image, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  },
  deleteById (id) {
    return Vue.http.get(GET_DELETE_BY_ID + id)
  },
  getById (id) {
    return Vue.http.get(GET_BY_ID_URL + id)
  },
  getImages () {
    return Vue.http.get(GET_IMAGES_URL)
  }
}
