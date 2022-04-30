import axios from 'axios'
import qs from 'qs'
import {Toast} from 'vant'
import router from '@/router'

const instance = axios.create({
  baseURL: process.env.BASE_API,
  timeout: 30000
})

instance.interceptors.request.use(config => {
  config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
  config.headers['Authorization'] = localStorage.getItem('token')
  return config
}, error => {
  return Promise.reject(error)
})

instance.interceptors.response.use(response => {
  console.log(response.data)
  return response.data
}, error => {
  console.log(error.response)
  if (error.response.status === 401) {
    localStorage.removeItem('token')
    router.replace({
      path: '/login',
      query: {redirect: router.currentRoute.fullPath}
    })
  } else {
    Toast.fail({
      message: error.response.data.msg,
      icon: 'cross'
    })
  }
  return Promise.reject(error)
})

export default {
  get (url, data) {
    return instance.get(url, {params: data})
  },
  post (url, data) {
    return instance.post(url, qs.stringify(data))
  }
}
