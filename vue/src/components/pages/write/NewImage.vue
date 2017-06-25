<template>
  <section class="newImage">
    <div class="row">
      <div class="twelve columns">
        <ImageForm :image="image"></ImageForm>
      </div>
    </div>
  </section>
</template>

<script>
import ImageForm from '@/components/partials/admin/forms/ImageForm'

export default {
  name: 'NewImage',
  components: {
    ImageForm: ImageForm
  },
  data () {
    return {
      image: {
        name: '',
        image: ''
      }
    }
  }
}
</script>
<style type="text/css">
.post-input {
  max-height: 100%;
  height: 500px;
}
</style>
