import http from './http'

export default {
  getSmsCode (data) {
    return http.get('/code/sms', data)
  },
  login (data) {
    return http.post('/user/login', data)
  },
  getGoodsList (data) {
    return http.get('/goods/list', data)
  },
  getGoods (data) {
    return http.get('/goods', data)
  },
  miaosha (data) {
    return http.post('/miaosha', data)
  },
  getMiaoshaResult (data) {
    return http.get('/miaosha/result', data)
  },
  getOrderList (data) {
    return http.get('/order/list', data)
  },
  getOrder (data) {
    return http.get('/order', data)
  },
  cancelOrder (data) {
    return http.post('/order/close', data)
  },
  getUserInfo (data) {
    return http.get('/user/info', data)
  }
}
