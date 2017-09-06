<template>
  <section class="PostActions" v-if="can()">
    <div class="row">
        <div class="twelve columns">
        <ul>
          <li>
            <router-link :to="{ name: 'editPost', params: { slug: post.slug } }">Edit</router-link>
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PostActions',
  computed: {
    post () {
      return this.$store.getters.post
    }
  },
  methods: {
    can () {
      return this.$store.getters.loggedIn &&
        this.post.user_id === this.$store.getters.user.id
    }
  }
}
</script>
