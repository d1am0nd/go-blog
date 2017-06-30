<template>
  <section class="newImage">
    <div class="row">
      <div class="six columns">
        <label for="title">Name</label>
        <input
          class="u-full-width"
          type="text"
          v-model="image.name"
          autocomplete="off"
          id="name">
      </div>
      <div class="six columns">
      </div>
    </div>
    <div class="row" v-if="image.name != ''">
      <div class="twelve columns">
        <input
          id="image-upload-input"
          type="file"
          name="image"
          accept="image/*"
          v-on:change="imageChange">
      </div>
    </div>
    <div class="row" v-if="image.name != ''">
      <div class="twelve columns">
        <img v-if="image.image !== ''" :src="image.image" />
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
        <button
          class="button button-primary"
          @click="post(image)">
          Post
        </button>
        <button
          v-if="image.id"
          class="button button-danger"
          @click="deleteById(image.id)">
          Delete
        </button>
      </div>
    </div>
  </section>
</template>

<script>
import Errors from '@/errors'
import images from '@/services/db/images'

export default {
  name: 'ImageForm',
  props: ['image'],
  methods: {
    post (image) {
      var formData = new FormData()
      var fileDom = document.getElementById('image-upload-input')
      formData.set('image', fileDom.files[0])
      formData.set('name', image.name)

      if (image.id) {
        images.update(formData, image.id)
        .then((res) => {
          this.image + res.body
        })
        .catch((res) => {
          Errors.newErrRes(res)
        })
      } else {
        images.new(formData)
        .then((res) => {
          this.$router.push({ name: 'images' })
        })
        .catch((res) => {
          Errors.newErrRes(res)
        })
      }
    },
    deleteById (id) {
      images.deleteById(id)
      .then((res) => {
        this.$router.push({ name: 'images' })
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    },
    imageChange (e) {
      var files = e.target.files || e.dataTransfer.files
      if (!files.length) {
        return
      }
      this.createImage(files[0])
    },
    createImage (file) {
      var reader = new FileReader()
      var vm = this

      reader.onload = (e) => {
        vm.image.image = e.target.result
      }
      reader.readAsDataURL(file)
    },
    removeImage: function (e) {
      this.image.image = ''
    }
  }
}
</script>
<style>
  .button-danger {
    color: #FFF;
    background-color: #ff5252;
    border-color: #ff5252;
  }
</style>
