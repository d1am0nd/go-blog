<template>
  <section class="PostRender">
    <h1>{{ post.title }}</h1>
    <div
      class="post-summary-render">
      {{ post.summary }}
    </div>
    <div
      class="post-content-render"
      v-html="compiledMarkdown">
    </div>
  </section>
</template>

<script>
import marked from 'marked'

export default {
  name: 'PostRender',
  props: ['post'],
  computed: {
    compiledMarkdown () {
      return marked(this.post.content, { renderer: false, sanitize: true })
    }
  }
}
</script>
<style type="text/css" scoped>
/* Responsive images */


/* Post design */
.post-summary-render {
  font-size: 24px;
  letter-spacing: 2px;
  margin-bottom: 20px;
}
.post-content-render p {
  font-size: 21px;
  letter-spacing: 1px;
  font-weight: 400;
}
</style>
