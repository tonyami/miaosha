<template>
  <div class="goods-list">
    <van-list
      v-model="loading"
      :finished="finished"
      finished-text="~~~ 我也是有底线的 ~~~"
      @load="onLoad">
      <van-card
        :tag="'第'+item.id+'期'"
        :price="(item.price/100).toFixed(2)"
        :title="item.name"
        :thumb="item.img"
        :origin-price="item.originPrice/100"
        v-for="item in list"
        :key="item.id"
        @click="getGoods(item)">
        <template #num>
          <van-tag type="success" size="medium" v-show="item.status === 1">进行中</van-tag>
          <van-tag type="warning" size="medium" v-show="item.status === 2">售罄</van-tag>
          <van-tag type="danger" size="medium" v-show="item.status === -1">已结束</van-tag>
          <van-count-down :time="item.duration*1000" v-show="item.status === 0" @finish="timerOver(item)"/>
        </template>
      </van-card>
    </van-list>
  </div>
</template>

<script>
import {List, Card, Tag, Button, CountDown} from 'vant'

export default {
  name: 'goods-list',
  data() {
    return {
      page: 1,
      list: [],
      loading: false,
      finished: false
    }
  },
  methods: {
    timerOver(item) {
      if (item.status === 0) {
        item.status = 1
      }
    },
    getGoods(item) {
      this.$router.push({
        name: 'goods',
        query: {
          id: item.id
        }
      })
    },
    onLoad() {
      this.$api.getGoodsList({page: this.page}).then(res => {
        this.loading = false
        this.list = this.list.concat(res)
        this.page++
        if (res && res.length < 10) {
          this.finished = true
        }
      }).catch(_ => {
        this.finished = true
      })
    }
  },
  components: {
    [List.name]: List,
    [Card.name]: Card,
    [Tag.name]: Tag,
    [Button.name]: Button,
    [CountDown.name]: CountDown
  }
}
</script>

<style scoped>
</style>
