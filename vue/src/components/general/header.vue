<template>
  <section class="header">
    <div class="title">{{ navTitle }}</div>
    <div class="row">
      <div class="twelve columns">
        <ul>
          <li><router-link :to="{ name: 'home' }">Home</router-link></li>
          <li v-if="loggedIn"><span>{{ user.name }}</span></li>
          <li v-if="loggedIn"><a href="javascript:;" @click="logout">Logout</a></li>
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
      navTitle: config.nav_title
    }
  },
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    },
    user () {
      return this.$store.getters.user
    }
  },
  methods: {
    logout () {
      auth.logout()
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
