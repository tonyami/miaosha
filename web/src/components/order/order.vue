<template>
  <div class="order">
    <div class="header">
      <div class="back" @click="goBack">
        <van-icon size="20" name="arrow-left"/>
      </div>
      <div class="title">{{order.status|statusText}}</div>
    </div>
    <div class="status border" v-if="order.status === 0">
      <div class="desc">
        <van-icon name="clock-o"/>
        <span>{{order.status|statusText}}</span>
      </div>
      <div class="timer">
        <span>剩余：</span>
        <van-count-down :time="order.duration*1000" v-show="order.status === 0" format="mm分ss秒" @finish="timerOver"/>
      </div>
    </div>
    <van-card
      num="1"
      :price="(order.goodsPrice/100).toFixed(2)"
      :title="'【第' + order.goodsId + '期】' + order.goodsName"
      :thumb="order.goodsImg"
      class="border"/>
    <div class="info border">
      <div class="row">订单号：{{order.orderId}}</div>
      <div class="row">下单时间：{{order.createTime}}</div>
      <div class="row" v-show="order.status === -1">关闭时间：{{order.updateTime}}</div>
    </div>
    <van-submit-bar
      :tip="(order.timeout/60)+'分钟内未支付，订单自动关闭'"
      tip-icon="info-o"
      :price="order.goodsPrice"
      button-text="支付"
      :disabled="disabled"
      :loading="showPayLoading"
      @submit="pay"
      v-if="order.status === 0">
      <van-button plain round type="default" :disabled="disabled" :loading="showCancelLoading" @click="cancel">取消订单
      </van-button>
    </van-submit-bar>
  </div>
</template>

<script>
import {Icon, CountDown, Card, SubmitBar, Button, Dialog} from 'vant'

export default {
  name: 'order',
  data () {
    return {
      order: {},
      payment: 0,
      showList: false,
      disabled: false,
      showPayLoading: false,
      showCancelLoading: false,
      chosenCoupon: -1,
      coupons: [],
      disabledCoupons: []
    }
  },
  methods: {
    pay () {
      this.disabled = true
      this.showPayLoading = true
      Dialog.alert({
        message: 'Just a moment...'
      }).then(() => {
        this.disabled = false
        this.showPayLoading = false
      })
    },
    timerOver () {
      if (this.order.status === 0) {
        this.order.status = -1
      }
    },
    cancel () {
      Dialog.confirm({
        title: '提示',
        message: '确认取消订单？'
      }).then(() => {
        this.disabled = true
        this.showCancelLoading = true
        this.$api.cancelOrder({orderId: this.order.orderId}).then(res => {
          this.disabled = false
          this.showCancelLoading = false
          this.getOrder()
        }).catch(_ => {
          this.disabled = false
          this.showCancelLoading = false
        })
      }).catch(() => {})
    },
    goBack () {
      this.$router.push('/order/list')
    },
    getOrder () {
      this.$api.getOrder({orderId: this.order.orderId}).then(res => {
        this.order = res
      }).catch(_ => {
      })
    }
  },
  filters: {
    statusText (status) {
      if (status === 0) {
        return '待支付'
      } else if (status === 1) {
        return '支付中'
      } else if (status === 2) {
        return '已完成'
      } else if (status === -1) {
        return '已关闭'
      }
      return ''
    }
  },
  created () {
    this.order.orderId = this.$route.query.orderId
    if (!this.order.orderId) {
      return
    }
    this.getOrder()
  },
  components: {
    [Icon.name]: Icon,
    [CountDown.name]: CountDown,
    [Card.name]: Card,
    [SubmitBar.name]: SubmitBar,
    [Button.name]: Button,
    [Dialog.name]: Dialog
  }
}
</script>

<style scoped>
.order {
  height: 100vh;
  background-color: #fff;
}

.border {
  padding: 0 16px;
}

.header {
  height: 50px;
}

.back {
  position: absolute;
  left: 10px;
  top: 10px;
  padding: 5px;
}

.title {
  line-height: 50px;
  text-align: center;
}

.status {
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.status .timer {
  display: flex;
  align-items: center;
}

.info {
  margin: 20px 0;
}

.row {
  color: #666;
  margin: 5px 0;
}

.pay {
  margin-top: 20px;
}

.pay .method {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px 0;
  margin: 10px 0;
  border: 1px solid #fff;
}

.pay .method.active {
  border: 1px solid #cfcfcf;
}

.pay .desc {
  margin-left: 10px;
}
</style>
