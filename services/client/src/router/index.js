import Vue from 'vue'
import Router from 'vue-router'
import ThePage from '@/components/ThePage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'ThePage',
      component: ThePage
    }
  ]
})
