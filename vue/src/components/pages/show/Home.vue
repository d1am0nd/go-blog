<template>
  <div class="hello">
    <h1>{{ title }}</h1>

    <div class="row" v-for="post in posts">
      <div class="col-md-12">
        <QuickPost :post="post"></QuickPost>
      </div>
    </div>
  </div>
</template>

<script>
import QuickPost from '@/components/partials/post/QuickPostRender'
import Errors from '@/errors'
import posts from '@/services/db/posts'
import Meta from '@/config/head'
import config from '@/../../config/page.json'

export default {
  name: 'hello',
  components: {
    QuickPost: QuickPost
  },
  data () {
    return {
      title: config.home_title
    }
  },
  computed: {
    posts () {
      return this.$store.getters.posts
    }
  },
  created () {
    this.fetchData()
    Meta.title('Home')
    Meta.description(config.meta_description)
  },
  methods: {
    fetchData () {
      posts.getPublished()
      .then((res) => {
        this.$store.commit('setPosts', res.body)
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
