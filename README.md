# miaosha

从零开始基于 Golang 实现秒杀系统，技术栈：Golang、Gin、Redis、MySQL、NSQ等，包括用户注册登录、分布式 Session、秒杀、异步下单、防刷限流等功能，采用 Jenkins+Docker 部署；文档将逐渐完善，项目中还有很多不足，欢迎大家批评指正。

> 基础思路参考慕课网若鱼1919老师 **[Java秒杀系统方案优化](https://coding.imooc.com/class/168.html)** 课程。

在线演示：Please wait

前端项目：https://github.com/mmtony/miaosha-web

## 功能

#### 初级

- [x] 用户注册登录
- [x] 分布式 Session
- [x] 商品列表
- [x] 商品详情
- [x] 秒杀（超卖问题）
- [x] 订单列表
- [x] 订单详情
- [x] 取消订单
- [x] 订单超时关闭
- [x] 我的（用户信息、订单统计、退出登录）

#### 进阶

- [x] 压测（JMeter）
- [ ] 安全优化（隐藏秒杀接口地址、请求限流）
- [x] 异步下单
- [ ] 支付
- [ ] 自动化部署（Nginx、Docker、Jenkins）

#### 高级

- 微服务
- k8s

......

