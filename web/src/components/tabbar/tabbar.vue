<template>
  <div class="tabbar">
    <van-tabbar v-model="active" active-color="#eb4d3c" inactive-color="#8b898b" safe-area-inset-bottom>
      <van-tabbar-item :name="item.name" :to="item.path" v-for="(item, index) in tabbars" :key="index">
        <van-icon slot="icon" slot-scope="props" :name="props.active ? item.active : item.inactive"></van-icon>
        <span>{{ item.title }}</span>
      </van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
import {Tabbar, TabbarItem, Icon} from 'vant'

export default {
  name: 'tabbar',
  data () {
    return {
      active: 'home',
      tabbars: [
        {title: '首页', name: 'home', path: '/home', active: 'wap-home', inactive: 'home-o'},
        {title: '订单', name: 'order-list', path: '/order/list', active: 'label', inactive: 'orders-o'},
        {title: '我的', name: 'user', path: '/user', active: 'friends', inactive: 'friends-o'}
      ]
    }
  },
  watch: {
    $route: 'changeActive'
  },
  created () {
    const name = this.$route.name
    this.setActive(name)
  },
  methods: {
    changeActive ({name}) {
      this.setActive(name)
    },
    setActive (name) {
      this.tabbars.forEach((item, index) => {
        if (item.name === name) {
          this.active = name
          return false
        }
      })
    }
  },
  components: {
    [Tabbar.name]: Tabbar,
    [TabbarItem.name]: TabbarItem,
    [Icon.name]: Icon
  }
}
</script>

<style scoped>
</style>
