import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/home.vue'
import BlogList from '@/list.vue'
import Snap from '@/Snap.vue'
import Blog from '@/components/blog.vue'

Vue.use(Router)

export default new Router({
  base: __dirname,
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/list',
      name: 'BlogList',
      component: BlogList
    },
    {
      path: '/snap',
      name: 'Snap',
      component: Snap
    },
    {
      path: '/blog/:id',
      name: 'blog',
      component: Blog
    }
  ]
})
