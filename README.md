# miaosha

基于 Golang 实现高并发秒杀系统，技术栈：Golang、Gin、Redis、MySQL、NSQ等，包括用户注册登录、分布式 Session、秒杀、异步下单、限流等功能，采用 Jenkins+Docker 部署；项目中还有很多不足，欢迎大家批评指正。

> 基础思路参考慕课网若鱼1919老师 **[Java秒杀系统方案优化](https://coding.imooc.com/class/168.html)** 课程。

[在线演示](https://mtony.cn/miaosha/)

## 功能

#### 初级

- [x] 用户注册登录
- [x] 分布式 Session
- [x] 商品列表
- [x] 商品详情
- [x] 秒杀
- [x] 订单列表
- [x] 订单详情
- [x] 取消订单
- [x] 订单超时关闭
- [x] 我的（用户信息、订单统计、退出登录）

#### 进阶

- [x] 压测（JMeter）
- [x] 安全优化（限流）
- [x] 异步下单
- [ ] 支付
- [x] 自动化部署（Nginx、Docker、Jenkins）

#### 高级

- 微服务
- k8s

## 快照

![登录](./snapshot/login.png)

![商品列表](./snapshot/goods-list.png)

![商品详情](./snapshot/goods.png)

![秒杀](./snapshot/miaosha.png)

![订单详情](./snapshot/order.png)

![订单列表](./snapshot/order-list.png)

![用户信息](./snapshot/user-info.png)