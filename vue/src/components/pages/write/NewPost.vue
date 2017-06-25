<template>
  <section class="newPost">
    <div class="row">
      <div class="twelve columns">
        <PostForm :post="post"></PostForm>
      </div>
    </div>
  </section>
</template>

<script>
import PostForm from '@/components/partials/admin/forms/Postform'

export default {
  name: 'NewPost',
  components: {
    PostForm: PostForm
  },
  data () {
    return {
      post: {
        title: 'Title',
        content: '### hello',
        summary: 'Intro',
        active: 0
      }
    }
  }
}
</script>
<style type="text/css">
.post-input {
  max-height: 100%;
  height: 500px;
}
</style>
