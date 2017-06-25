<template>
  <section class="Image">
    <div class="row">
      <div class="twelve columns">
        <ImageForm :image="image"></ImageForm>
      </div>
    </div>
  </section>
</template>

<script>
import Errors from '@/errors'
import images from '@/services/db/images'
import ImageForm from '@/components/partials/admin/forms/ImageForm'

export default {
  name: 'Image',
  components: {
    ImageForm: ImageForm
  },
  data () {
    return {
      image: {
        name: '',
        image: '',
        path: ''
      }
    }
  },
  created () {
    this.fetchData()
  },
  watch: {
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      images.getById(this.$route.params.id)
      .then((res) => {
        console.log(res.body)
        this.image = res.body
      })
      .catch((err) => {
        Errors.newErrRes(err)
      })
    }
  }
}
</script>
