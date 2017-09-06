<template>
  <section class="Post">
    <div class="row">
      <div class="twelve columns">
        <PostForm></PostForm>
      </div>
    </div>
  </section>
</template>

<script>
import Errors from '@/errors'
import posts from '@/services/db/posts'
import PostForm from '@/components/partials/admin/forms/Postform'

export default {
  name: 'Post',
  components: {
    PostForm: PostForm
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
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
