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
  props: ['post'],
  data () {
    return {
      auth: this.$root.auth
    }
  },
  methods: {
    can () {
      return this.auth.check && this.post.user_id === this.auth.user.id
    }
  }
}
</script>
