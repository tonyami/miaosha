<template>
  <div class="user">
    <div class="header">
      <div class="top">
        <img src="/static/go.png" alt="avatar" width="100%" height="100%" class="avatar">
        <div class="info">
          <div class="row"><span class="mobile">{{user.mobile}}</span></div>
          <div class="row"><span class="uid">UID:{{user.id}}</span></div>
        </div>
      </div>
      <div class="count">
        <div class="col">
          <div><span class="num">{{count.unfinished}}</span></div>
          <div><span class="desc">未完成</span></div>
        </div>
        <div class="col">
          <div><span class="num">{{count.finished}}</span></div>
          <div><span class="desc">已完成</span></div>
        </div>
        <div class="col">
          <div><span class="num">{{count.closed}}</span></div>
          <div><span class="desc">已关闭</span></div>
        </div>
      </div>
    </div>
    <van-button type="default" block plain style="color:#eb4d3c" @click="logout">退出登录</van-button>
  </div>
</template>

<script>
import {Button, Dialog} from 'vant'

export default {
  name: 'user',
  data () {
    return {
      user: {},
      count: {}
    }
  },
  created () {
    this.getUserInfo()
  },
  methods: {
    getUserInfo () {
      this.$api.getUserInfo().then(res => {
        this.user = res.user
        this.count = res.count
      }).catch(_ => {})
    },
    logout () {
      Dialog.confirm({
        title: '提示',
        message: '确认退出登录？'
      }).then(() => {
        localStorage.removeItem('token')
        this.$router.push('/')
      }).catch(() => {})
    }
  },
  components: {
    [Button.name]: Button,
    [Dialog.name]: Dialog
  }
}
</script>

<style scoped>
.header {
  height: 200px;
  background: linear-gradient(to top left,#E3A054,#E8535A);
}
.top {
  padding-top: 30px;
  padding-left: 25px;
  display: flex;
}
.avatar {
  width: 72px;
  height: 72px;
  background-color: #fff;
  border-radius: 50%;
  font-size: 0;
  border: 0;
}
.info {
  margin-left: 20px;
  color: #f3f3f3;
}
.info .row {
  margin: 10px 0;
}
.info .mobile {
  color: #f9f9f9;
  font-weight: bold;
}
.info .uid {
  font-size: 12px;
  color: #f2f2f2;
}
.count {
  margin: 20px 0;
  display: flex;
}
.count .col {
  height: 42px;
  flex: 1;
  border-right: 1px solid #f2f2f2;
  color: #f2f2f2;
  text-align: center;
}
.count .col:nth-child(3) {
  border-right: none;
}
.count .col .num {
  line-height: 1.5;
  color: #f9f9f9;
  font-weight: bold;
}
.count .col .desc {
  font-size: 12px;
  color: #f2f2f2;
}
</style>
