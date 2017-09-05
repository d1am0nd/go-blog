import Vue from 'vue'
import store from '@/store'
import Router from 'vue-router'
import Home from '@/components/pages/show/Home'
import Post from '@/components/pages/show/Post'
import Login from '@/components/pages/auth/Login'
import NewPost from '@/components/pages/write/NewPost'
import Images from '@/components/pages/show/Images'
import NewImage from '@/components/pages/write/NewImage'
import EditPost from '@/components/pages/write/EditPost'
import EditImage from '@/components/pages/write/EditImage'
// import Register from '@/components/pages/auth/Register'

Vue.use(Router)

var router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    /*
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    */
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/admin/posts/write',
      name: 'newPost',
      component: NewPost
    },
    {
      path: '/admin/posts/edit/:slug',
      name: 'editPost',
      component: EditPost
    },
    {
      path: '/admin/images/new',
      name: 'newImage',
      component: NewImage
    },
    {
      path: '/admin/images',
      name: 'images',
      component: Images
    },
    {
      path: '/admin/images/edit/:id',
      name: 'editImage',
      component: EditImage
    },
    {
      path: '/posts/:slug',
      name: 'post',
      component: Post
    }
  ]
})

export default router

var authRoutes = [
  'newPost',
  'editPost',
  'newImage',
  'images',
  'editImage'
]
var guestRoutes = [
  'register',
  'login'
]

router.beforeEach((to, from, next) => {
  if (guestRoutes.indexOf(to.name) !== -1 && to.name !== 'login' && store.getters.loggedIn) {
    router.push({name: 'home'})
  } else if (authRoutes.indexOf(to.name) !== -1 && store.getters.loggedIn === false) {
    router.push({name: 'login'})
  } else {
    next()
  }
})
