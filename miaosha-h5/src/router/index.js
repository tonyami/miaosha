import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login/login'
import Home from '@/components/home/home'
import OrderList from '@/components/order-list/order-list'
import User from '@/components/user/user'
import Goods from '@/components/goods/goods'
import Order from '@/components/order/order'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      redirect: '/home'
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: {
        showTabbar: false,
        requireAuth: false
      }
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
      meta: {
        showTabbar: true,
        requireAuth: false
      }
    },
    {
      path: '/order/list',
      name: 'order-list',
      component: OrderList,
      meta: {
        showTabbar: true,
        requireAuth: true
      }
    },
    {
      path: '/user',
      name: 'user',
      component: User,
      meta: {
        showTabbar: true,
        requireAuth: true
      }
    },
    {
      path: '/goods',
      name: 'goods',
      component: Goods,
      meta: {
        showTabbar: false,
        requireAuth: false
      }
    },
    {
      path: '/order',
      name: 'order',
      component: Order,
      meta: {
        showTabbar: false,
        requireAuth: true
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.requireAuth && !localStorage.getItem('token')) {
    next({ path: '/login', query: { redirect: to.fullPath } })
  } else {
    next()
  }
})

const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

export default router
