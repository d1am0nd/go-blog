<template>
  <section class="footer">
    <small>{{ footer }}</small>
  </section>
</template>

<script>
import page from '@../../config/page.json'

export default {
  name: 'Footer',
  data () {
    return {
      footer: page.footer
    }
  }
}
</script>
