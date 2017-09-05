<template>
  <div id="app">
    <div class="container">
        <MainHeader></MainHeader>
        <AdminHeader v-if="loggedIn"></AdminHeader>
        <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Header from '@/components/general/Header'
import AdminHeader from '@/components/general/AdminHeader'

export default {
  name: 'app',
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    }
  },
  components: {
    MainHeader: Header,
    AdminHeader: AdminHeader
  }
}
</script>
