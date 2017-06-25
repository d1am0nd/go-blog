<template>
  <div id="app">
    <div class="container">
        <MainHeader></MainHeader>
        <AdminHeader v-if="auth.check"></AdminHeader>
        <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Header from './components/general/Header'
import AdminHeader from '@/components/general/AdminHeader'

export default {
  name: 'app',
  data () {
    return {
      auth: this.$root.auth
    }
  },
  components: {
    MainHeader: Header,
    AdminHeader: AdminHeader
  }
}
</script>
