<template>
  <section class="Post">
    <div class="row">
      <div class="twelve columns">
        <PostForm :post="post"></PostForm>
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
  data () {
    return {
      post: {
        title: '',
        slug: '',
        summary: '',
        content: '',
        active: 0,
        published_at: ''
      }
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
        this.post = res.body
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
