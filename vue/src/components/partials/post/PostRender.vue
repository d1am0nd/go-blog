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
var renderer = new marked.Renderer()
console.log(renderer)
renderer.link = (href, title, text) => {
  if (renderer.options.sanitize) {
    try {
      var prot = decodeURIComponent(unescape(href))
        .replace(/[^\w:]/g, '')
        .toLowerCase()
    } catch (e) {
      return ''
    }
    if (prot.indexOf('javascript:') === 0 || prot.indexOf('vbscript:') === 0) {
      return ''
    }
  }
  var out = '<a href="' + href + '"'
  out += ' target="_blank"'
  if (title) {
    out += ' title="' + title + '"'
  }
  out += '>' + text + '</a>'
  return out
}

export default {
  name: 'PostRender',
  props: ['post'],
  computed: {
    compiledMarkdown () {
      return marked(this.post.content, { renderer: renderer, sanitize: true })
    }
  }
}
</script>
<style type="text/css">
/* Responsive images and centered */
.post-content-render img {
  max-width: 100%;
  box-sizing: border-box;
  position: relative;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, 0%);
}

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
