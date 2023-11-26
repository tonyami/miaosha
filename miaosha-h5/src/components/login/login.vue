<template>
  <div class="login">
    <van-nav-bar title="登录/注册">
      <van-icon name="arrow-left" slot="left" class="back" @click="goBack"/>
    </van-nav-bar>
    <van-cell-group>
      <van-field
        v-model="mobile"
        type="number"
        maxlength="11"
        placeholder="请输入手机号码（新用户自动注册）"
        :error-message="mobileMsg"
      />
      <van-field
        v-model="smsCode"
        type="number"
        maxlength="6"
        center
        clearable
        placeholder="请输入验证码"
        :error-message="smsCodeMsg"
      >
        <template #button>
          <van-button size="small" plain type="primary" @click="getSmsCode" :loading="isSmsLoading" :disabled="isSmsDisabled">{{smsTxt}}</van-button>
        </template>
      </van-field>
    </van-cell-group>
    <div class="btn">
      <van-button type="primary" :loading="isLogining" :disabled="isLogining" block @click="login">登录</van-button>
    </div>
  </div>
</template>

<script>
import {NavBar, Icon, CellGroup, Field, Button, Dialog, Toast} from 'vant'

export default {
  name: 'login',
  data() {
    return {
      mobile: '',
      mobileMsg: '',
      smsCode: '',
      smsCodeMsg: '',
      isLogining: false,
      timer: 0,
      isSmsLoading: false,
      isSmsDisabled: false,
      duration: 30,
      smsTxt: '获取验证码'
    }
  },
  methods: {
    login() {
      if (!this.check()) {
        return
      }
      this.isLogining = true
      this.$api.login({
        'mobile': this.mobile,
        'smsCode': this.smsCode
      }).then(res => {
        this.isLogining = false
        localStorage.setItem('token', res.token)
        let path = this.$route.query.redirect ? this.$route.query.redirect : '/'
        this.$router.replace(path)
      }).catch(_ => {
        this.isLogining = false
      })
    },
    check() {
      if (this.mobile.length !== 11) {
        this.mobileMsg = '手机号码格式错误'
        return false
      }
      this.mobileMsg = ''
      if (this.smsCode.length !== 6) {
        this.smsCodeMsg = '验证码错误'
        return false
      }
      this.smsCodeMsg = ''
      return true
    },
    getSmsCode() {
      if (this.mobile.length !== 11) {
        Toast.fail({
          message: '手机号码格式错误',
          icon: 'cross'
        })
        return false
      }
      this.isSmsLoading = true
      this.isSmsDisabled = true
      this.$api.getSmsCode({
        'mobile': this.mobile
      }).then(res => {
        this.isSmsLoading = false
        this.smsTxt = this.duration-- + '秒后再试'
        this.timer = setInterval(() => {
          if (this.duration <= 0) {
            clearInterval(this.timer)
            this.isSmsDisabled = false
            this.duration = 30
            this.smsTxt = '获取验证码'
            return
          }
          this.smsTxt = this.duration-- + '秒后再试'
        }, 1000)
        setTimeout(() => {
          Dialog.alert({
            message: '【秒杀】验证码：' + res.code
          }).then(() => {
            // on close
          })
        }, 3000)
      }).catch(_ => {
        this.isSmsDisabled = false
      })
    },
    goBack() {
      this.$router.replace('/')
    }
  },
  components: {
    [NavBar.name]: NavBar,
    [Icon.name]: Icon,
    [CellGroup.name]: CellGroup,
    [Field.name]: Field,
    [Button.name]: Button
  }
}
</script>

<style scoped>
.login {
  background-color: #fff;
}

.back {
  color: #8b898b;
}

.btn {
  padding: 20px 16px;
}
</style>
