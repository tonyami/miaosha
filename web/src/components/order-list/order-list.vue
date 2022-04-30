<template>
  <div class="order-list">
    <van-tabs v-model="active" @click="changeTab">
      <van-tab title="全部" name="all"></van-tab>
      <van-tab title="未完成" name="unfinished"></van-tab>
      <van-tab title="已完成" name="finished"></van-tab>
      <van-tab title="已关闭" name="closed"></van-tab>
    </van-tabs>
    <van-list
      v-model="loading"
      :finished="finished"
      finished-text="~~~ 我也是有底线的 ~~~"
      @load="onLoad">
      <van-card
        :price="(item.goodsPrice/100).toFixed(2)"
        :title="'【第' + item.goodsId + '期】' + item.goodsName"
        :thumb="item.goodsImg"
        v-for="item in list" :key="item.id"
        @click="getOrder(item)">
        <template #tags>
          <div class="row">订单号：{{item.orderId}}</div>
          <div class="row">下单时间：{{item.createTime}}</div>
        </template>
        <template #num>
          <span v-show="item.status === 0" style="color: #ff976a;">待支付</span>
          <span v-show="item.status === 1" style="color: #ff976a;">支付中</span>
          <span v-show="item.status === 2">已完成</span>
          <span v-show="item.status === -1">已关闭</span>
        </template>
      </van-card>
    </van-list>
  </div>
</template>

<script>
import {Tab, Tabs, List, Card} from 'vant'

export default {
  name: 'order',
  data () {
    return {
      active: '',
      page: 1,
      list: [],
      loading: false,
      finished: false
    }
  },
  methods: {
    getOrder (item) {
      this.$router.push({
        name: 'order',
        query: {
          orderId: item.orderId
        }
      })
    },
    changeTab (name, title) {
      this.list = []
      this.page = 1
      this.finished = false
      this.onLoad()
    },
    onLoad () {
      this.loading = true
      this.$api.getOrderList({page: this.page, status: this.active}).then(res => {
        this.list = this.list.concat(res)
        // 加载状态结束
        this.loading = false
        this.page++
        // 数据全部加载完成
        if (res.length < 10) {
          this.finished = true
        }
      }).catch(_ => {
      })
    }
  },
  components: {
    [Tabs.name]: Tabs,
    [Tab.name]: Tab,
    [List.name]: List,
    [Card.name]: Card
  }
}
</script>

<style scoped>
.order-list {
  margin-bottom: 50px;
}

.row {
  margin: 5px 0;
}
</style>
