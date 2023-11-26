<template>
  <div class="home">
    <goods-list/>
  </div>
</template>

<script>
import GoodsList from '@/components/goods-list/goods-list'
export default {
  name: 'home',
  components: {
    GoodsList
  }
}
</script>

<style scoped>
.home {
  margin-bottom: 50px;
}
</style>
