<template>
  <div class="images">
    <h1>Hip Hop Blog Images</h1>

    <div class="row">
      <div class="col-md-12">
        <input
          type="text"
          @input="updateFilter"
          placeholder="Filter"
          v-model="filter">
        </input>
      </div>
    </div>

    <div class="row" v-for="image in filteredImages">
      <div class="col-md-12">
        {{ image.name }} <br>
        {{ image.path }} <br>
        <router-link :to="{ name: 'editImage', params: { id: image.id } }">
          <img :src="image.path">
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import Errors from '@/errors'
import images from '@/services/db/images'

export default {
  name: 'images',
  data () {
    return {
      filter: ''
    }
  },
  created () {
    this.fetchData()
  },
  computed: {
    filteredImages () {
      return this.$store.getters.filteredImages
    }
  },
  methods: {
    fetchData () {
      images.getImages()
      .then((res) => {
        this.$store.commit('setImages', res.body)
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    },
    updateFilter (e) {
      this.$store.commit('setImagesFilter', e.target.value)
    }
  }
}
</script>
