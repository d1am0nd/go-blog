<template>
  <section class="register">
    <div class="row">
      <div class="offset-by-four four columns">
        <h2 class="docs-header">Login</h2>
        <form v-on:submit.prevent="login">
          <label for="register-email">Email</label>
          <input
            class="u-full-width"
            type="email"
            v-model="creds.email"
            id="register-email">

          <label for="register-password">Your password</label>
          <input
            class="u-full-width"
            type="password"
            v-model="creds.password"
            id="register-password">

          <input
            class="button-primary"
            type="submit"
            value="Submit">
        </form>
      </div>
    </div>
  </section>
</template>

<script>
import auth from '@/auth/auth.js'

export default {
  name: 'Header',
  data () {
    return {
      creds: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    login () {
      auth.login(this, this.creds)
    }
  }
}
</script>
