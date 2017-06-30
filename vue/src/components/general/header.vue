<template>
  <section class="header">
    <div class="title">{{ navTitle }}</div>
    <div class="row">
      <div class="twelve columns">
        <ul>
          <li><router-link :to="{ name: 'home' }">Home</router-link></li>
          <li v-if="auth.check"><span>{{ auth.user.name }}</span></li>
          <li v-if="auth.check"><a href="javascript:;" @click="logout">Logout</a></li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
import auth from '@/auth/auth'
import config from '@/../../config/page.json'

export default {
  name: 'Header',
  data () {
    return {
      auth: this.$root.auth,
      navTitle: config.nav_title
    }
  },
  methods: {
    logout () {
      auth.logout()
    },
    refreshAuth () {
      auth.refresh()
    },
    checkAuth () {
      return auth.checkAuth()
    }
  }
}
</script>
<style type="text/css" scoped>
ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    overflow: hidden;
}

li {
    float: left;
}

li a, span {
    display: block;
    text-align: center;
    padding-right: 32px;
    text-decoration: none;
}

li a:hover {
    color: grey;
}
</style>
