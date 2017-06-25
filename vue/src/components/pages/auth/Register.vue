<template>
  <section class="register">
    <div class="row">
      <div class="offset-by-four four columns">
        <h2 class="docs-header">Register</h2>
        <form v-on:submit.prevent="register">
          <label for="register-email">Email</label>
          <input
            class="u-full-width"
            type="email"
            v-model="creds.email"
            id="register-email">

          <label for="register-username">Username</label>
          <input
            class="u-full-width"
            type="text"
            v-model="creds.name"
            id="register-username">

          <label for="register-password">Your password</label>
          <input
            class="u-full-width"
            type="password"
            v-model="creds.password"
            id="register-password">

          <label for="register-password-repeat">Repeat password</label>
          <input
            class="u-full-width"
            type="password"
            v-model="creds.password_confirmation"
            id="register-password-repeat">
          <input
            class="button-primary"
            type="submit"
            value="Submit">
        </form>
          <button
            class="button-primary"
            @click="test"
            type="button">test</button>
      </div>
    </div>
  </section>
</template>

<script>
import auth from '@/auth/auth.js'

export default {
  name: 'Register',
  data () {
    return {
      creds: {
        name: '',
        email: '',
        password: '',
        password_confirmation: ''
      }
    }
  },
  methods: {
    register () {
      auth.register(this, this.creds)
    },
    test () {
      this.$http.get('/test')
      .then((res) => {
        console.log(res)
      })
    }
  }
}
</script>
