# miaosha

从零开始基于 Golang 实现秒杀系统，技术栈：Golang、Gin、RabbitMQ、Redis、MySQL等，包括用户注册登录、分布式 Session、秒杀、异步下单、防刷限流等功能，采用Jenkins+Docker部署；前期要同步完成配套前端项目，时间比较紧，文档将逐渐完善，项目中还有很多不足，欢迎大家批评指正，一起学习，共同进步。

> 项目基础思路参考慕课网若鱼1919老师 **[Java秒杀系统方案优化](https://coding.imooc.com/class/168.html)** 课程，收获颇丰。

在线演示：Please wait

前端项目：https://github.com/mmtony/miaosha-web

## 功能

#### 初级

- [x] 用户注册登录
- [x] 分布式 Session（Redis）
- [x] 商品列表
- [x] 商品详情
- [x] 秒杀（解决超卖问题）
- [x] 订单列表
- [x] 订单详情
- [ ] 我的（用户信息、订单统计、退出登录）

#### 进阶

- [ ] 防刷限流
- [ ] 异步下单（消息队列）
- [ ] 订单超时关闭（定时任务、延迟队列）
- [ ] 优惠券
- [ ] 支付
- [ ] 部署（Nginx、Docker、Jenkins）

#### 高级

- 微服务
- k8s

......

