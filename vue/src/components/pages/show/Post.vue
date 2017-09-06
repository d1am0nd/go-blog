<template>
  <section class="Post">
    <div class="row">
      <div class="twelve columns">
        <PostRender :post="post"></PostRender>
        <router-link
            :to="{ name: 'home' }"
            class="button">
            Back to list
        </router-link>
        <PostActions :post="post"></PostActions>
      </div>
    </div>
  </section>
</template>

<script>
import Errors from '@/errors'
import Meta from '@/config/head'
import posts from '@/services/db/posts'
import PostActions from '@/components/partials/admin/PostActions'
import PostRender from '@/components/partials/post/PostRender'

export default {
  name: 'Post',
  components: {
    PostRender: PostRender,
    PostActions: PostActions
  },
  computed: {
    post () {
      return this.$store.getters.post
    }
  },
  created () {
    this.fetchData()
  },
  watch: {
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      posts.getBySlug(this.$route.params.slug)
      .then((res) => {
        this.$store.commit('setPost', res.body)
        Meta.title(this.post.title)
        Meta.description(this.post.summary)
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
