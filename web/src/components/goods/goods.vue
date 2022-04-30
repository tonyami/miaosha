<template>
  <div class="goods">
    <div class="back" @click="goBack">
      <van-icon size="20" name="arrow-left"/>
    </div>
    <div class="img">
      <img :src="goods.img" :alt="goods.name" width="100%" height="100%"/>
    </div>
    <div class="title">【第{{goods.id}}期】{{goods.name}}
      <span>
        <van-tag type="primary" v-show="goods.status === 0">未开始</van-tag>
        <van-tag type="success" v-show="goods.status === 1">进行中</van-tag>
        <van-tag type="warning" v-show="goods.status === 2">售罄</van-tag>
        <van-tag type="danger" v-show="goods.status === -1">已结束</van-tag>
      </span>
    </div>
    <div class="desc">
      <p>
        棒棒糖，这是一种深受广大人民喜爱的糖果类食品，最初是一颗硬糖插在一根小棒上，后来研发出很多更好吃好玩的品种，不仅小孩子深爱棒棒糖，一些童心未泯的成年人也会吃。棒棒糖的类型有凝胶型糖果，硬糖型，牛奶型糖果，巧克力型还有牛奶加水果型。</p>
      <p>1958年，闻名世界的棒棒糖（lollipop）的发明人恩里克·伯纳特·丰利亚多萨，首次推出这种带棍的糖果，结果使一家几乎经营不下去的糖果公司扭亏为盈。</p>
      <p>对于一些人来说，在嘴里含着一颗糖果，糖果的棍从嘴唇间露出来，已经成为一种时髦而有趣的标志。</p>
      <p>生产棒棒糖（lollipop）的这家西班牙家族公司每年生产40亿只棒棒糖。在全世界拥有许多分公司和工厂，雇佣了将近2000人。他们生产的棒棒糖超过50多个品种，其中包括一种专门针对墨西哥市场的辣味棒棒糖。</p>
    </div>
    <van-submit-bar
      :loading="showLoading"
      :price="goods.price"
      label="秒杀价："
      :disabled="goods.status !== 1 || actionDisabled"
      button-text="立即秒杀" @submit="action">
      <van-count-down :time="goods.duration*1000" v-show="goods.status === 0" @finish="timerOver"/>
    </van-submit-bar>
    <van-overlay :show="showOverLay" :z-index="9999" :lock-scroll="true">
      <van-loading size="36px" vertical color="#fff" class="wrapper">排队中...</van-loading>
    </van-overlay>
  </div>
</template>

<script>
import {Icon, SubmitBar, Tag, CountDown, Overlay, Loading} from 'vant'

export default {
  name: 'goods-detail',
  data () {
    return {
      showLoading: false,
      actionDisabled: false,
      showOverLay: false,
      goods: {}
    }
  },
  methods: {
    timerOver () {
      if (this.goods.status === 0) {
        this.goods.status = 1
      }
    },
    action () {
      this.actionDisabled = true
      this.showLoading = true
      this.$api.miaosha({goodsId: this.goods.id}).then(_ => {
        this.showLoading = false
        this.showOverLay = true
        this.getMiaoshaResult()
      }).catch(_ => {
        this.showLoading = false
        this.actionDisabled = false
      })
    },
    getMiaoshaResult () {
      this.$api.getMiaoshaResult({goodsId: this.goods.id}).then(res => {
        if (res.status === 0) {
          setTimeout(this.getMiaoshaResult, 2000)
        } else if (res.status === 1) {
          this.showOverLay = false
          this.actionDisabled = false
          this.$router.push({
            name: 'order',
            query: {
              orderId: res.orderId
            }
          })
        }
      })
    },
    goBack () {
      this.$router.push('/')
    },
    getGoods () {
      this.$api.getGoods({id: this.goods.id}).then(res => {
        this.goods = res
      }).catch(_ => {
      })
    }
  },
  created () {
    this.goods.id = this.$route.query.id
    if (!this.goods.id) {
      return
    }
    this.getGoods()
  },
  components: {
    [Icon.name]: Icon,
    [SubmitBar.name]: SubmitBar,
    [Tag.name]: Tag,
    [CountDown.name]: CountDown,
    [Overlay.name]: Overlay,
    [Loading.name]: Loading
  }
}
</script>

<style scoped>
.goods {
  position: relative;
  background-color: #fff;
  margin-bottom: 50px;
}

.back {
  position: fixed;
  top: 10px;
  left: 10px;
  padding: 6px;
  color: #999;
}

.img {
  width: 100%;
  height: 360px;
  border: 0;
}

.title, .desc {
  margin: 16px;
}

.wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>
