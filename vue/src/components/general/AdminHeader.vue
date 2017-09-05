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
          <li v-for="post in posts">
            <router-link :to="{ name: 'post', params: { slug: post.slug } }">{{ post.title }}</router-link>
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
import auth from '@/auth/auth'

export default {
  name: 'Header',
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    },
    posts () {
      return this.$store.getters.posts
    }
  },
  methods: {
    logout () {
      auth.logout()
    }
  }
}
</script>
