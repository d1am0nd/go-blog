<template>
  <section class="admin-header">
    <div class="row">
      <div class="six columns">
        <ul>
          <li><router-link :to="{ name: 'images' }">Images</router-link></li>
          <li v-if="loggedIn"><router-link :to="{ name: 'newPost' }">New post</router-link></li>
          <li v-if="loggedIn"><router-link :to="{ name: 'newImage' }">New Image</router-link></li>
        </ul>
      </div>
      <div class="six columns">
        <ul>
          <li v-for="post in myPosts">
            <router-link :to="{ name: 'post', params: { slug: post.slug } }">{{ post.title }}</router-link>
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
import auth from '@/auth/auth'
import Errors from '@/errors'
import posts from '@/services/db/posts'

export default {
  name: 'Header',
  data () {
    return {
      myPosts: []
    }
  },
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    }
  },
  methods: {
    logout () {
      auth.logout()
    }
  },
  created () {
    if (this.loggedIn) {
      posts.getMine()
      .then((res) => {
        this.myPosts = res.body
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
