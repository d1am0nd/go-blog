<template>
  <section class="QuickPostRender">
    <h2>{{ post.title }}</h2>
    <span>{{ post.published_at.String }}</span>
    <div
      class="post-summary-render">
      {{ post.summary }}
    </div>
    <router-link
        :to="{ name: 'post', params: { slug: post.slug } }"
        class="button">
        Check it out
    </router-link>
  </section>
</template>

<script>
export default {
  name: 'QuickPostRender',
  props: ['post']
}
</script>
