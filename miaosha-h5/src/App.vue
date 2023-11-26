<template>
  <div id="app">
    <router-view />
    <tabbar v-show="$route.meta.showTabbar"></tabbar>
  </div>
</template>

<script>
import Tabbar from '@/components/tabbar/tabbar'
export default {
  name: 'App',
  components: {
    Tabbar
  }
}
</script>

<style>
</style>
