<template>
  <section class="newPost">
    <div class="row">
      <div class="six columns">
        <label for="title">Title</label>
        <input
          class="u-full-width"
          type="text"
          v-model="post.title"
          id="title"
          placeholder="That shit real">
      </div>
      <div class="six columns">
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
        <textarea
          class="u-full-width post-summary"
          rows="4"
          v-model="post.summary"></textarea>
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
        <textarea
          class="u-full-width post-content"
          rows="10"
          v-model="post.content"></textarea>
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
        <PostRender :post="post"></PostRender>
      </div>
    </div>
    <div class="row">
      <div class="six columns">
        <label class="example-send-yourself-copy">
          <input
            type="radio"
            id="ch-post-draft"
            v-model="post.active"
            value="0"
            name="publish">
          <span class="label-body">Draft</span>
        </label>
        <label class="example-send-yourself-copy">
          <input
            type="radio"
            id="ch-post-publish"
            v-model="post.active"
            value="1"
            name="publish">
          <span class="label-body">Publish</span>
        </label>
      </div>
      <div class="six columns">
        <br>
        <DatePicker
          v-if="post.active == 1"
          :date="startTime"
          :option="dpconfig.option"
          :limit="dpconfig.limit">
        </DatePicker>
      </div>
    </div>
    <div class="row">
      <div class="twelve columns">
        <button
          class="button button-primary"
          @click="publish(post)">
          Post
        </button>
      </div>
    </div>
  </section>
</template>

<script>
import Errors from '@/errors'
import datepicker from 'vue-datepicker/vue-datepicker-es6.vue'
import dpconfig from '@/config/datepicker'
import posts from '@/services/db/posts'
import PostRender from '@/components/partials/post/PostRender'

export default {
  name: 'PostForm',
  props: ['post'],
  components: {
    PostRender: PostRender,
    DatePicker: datepicker
  },
  data () {
    return {
      dpconfig: dpconfig,
      startTime: {
        time: ''
      }
    }
  },
  created () {
    this.startTime.time = this.post.published_at.String
  },
  watch: {
    'startTime.time': 'updatePublishedAt',
    'post.published_at.String': 'updateStartTime'
  },
  methods: {
    publish (post) {
      if (post.id) {
        posts.update(post)
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          Errors.newErrRes(err)
        })
      } else {
        posts.new(post)
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          Errors.newErrRes(err)
        })
      }
    },
    updatePublishedAt (val) {
      this.post.published_at.String = val
      this.post.published_at.Valid = 1
    },
    updateStartTime (val) {
      this.startTime.time = val
    }
  }
}
</script>
<style type="text/css">
.post-content {
  max-height: 100%;
  height: 500px;
}
.post-summary {
  height: 120px;
}
</style>
