import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/dbinstances',
      name: 'dbinstances',
      component: Home
    },
    {
      path: '/',
      name: 'dbinstances',
      component: Home
    },
    {
      path: '/sqlrequest',
      name: 'sqlrequest',
      component: () => import(/* webpackChunkName: "about" */ './views/Sql_request.vue')
    },
    {
      path: '/duty-1',
      name: 'dutylist',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/duty-2',
      name: 'createduty',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/create_duty.vue')
    }
  ]
})
