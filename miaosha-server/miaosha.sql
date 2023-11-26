/*
Date: 2021-11-22 21:47:59
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for miaosha_goods
-- ----------------------------
DROP TABLE IF EXISTS `miaosha_goods`;
CREATE TABLE `miaosha_goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `name` varchar(255) NOT NULL COMMENT '商品名称',
  `img` varchar(255) NOT NULL COMMENT '商品图片',
  `origin_price` bigint(20) NOT NULL COMMENT '商品价格',
  `price` bigint(20) NOT NULL COMMENT '秒杀价格',
  `stock` int(11) unsigned NOT NULL COMMENT '库存',
  `start_time` datetime NOT NULL COMMENT '秒杀开始时间',
  `end_time` datetime NOT NULL COMMENT '秒杀结束时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- ----------------------------
-- Records of miaosha_goods
-- ----------------------------
INSERT INTO `miaosha_goods` VALUES ('1', '棒棒糖（苹果味）', '/static/bbt.jpg', '200', '50', '0', '2020-12-25 15:00:00', '2020-12-31 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('2', '棒棒糖（橘子味）', '/static/bbt.jpg', '200', '10', '0', '2020-12-25 00:00:00', '2021-01-31 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('3', '棒棒糖（芒果味）', '/static/bbt1.jpg', '200', '10', '0', '2020-12-25 00:00:00', '2021-01-31 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('4', '棒棒糖（荔枝味）', '/static/bbt.jpg', '200', '20', '0', '2021-04-01 00:00:00', '2021-11-30 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('5', '棒棒糖（葡萄味）', '/static/bbt1.jpg', '200', '20', '0', '2021-04-01 00:00:00', '2021-11-30 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('6', '棒棒糖（原味）', '/static/bbt.jpg', '200', '20', '10', '2021-11-01 00:00:00', '2021-11-30 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '200', '30', '10', '2021-11-01 00:00:00', '2021-11-30 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '200', '10', '10', '2021-11-02 00:00:00', '2021-11-30 23:59:59');
INSERT INTO `miaosha_goods` VALUES ('9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '200', '15', '10', '2021-11-03 00:00:00', '2021-11-30 23:59:59');

-- ----------------------------
-- Table structure for miaosha_order
-- ----------------------------
DROP TABLE IF EXISTS `miaosha_order`;
CREATE TABLE `miaosha_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` varchar(32) NOT NULL COMMENT '订单id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `goods_id` bigint(20) NOT NULL COMMENT '商品id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uid_gid` (`user_id`,`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=313 DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- ----------------------------
-- Records of miaosha_order
-- ----------------------------

-- ----------------------------
-- Table structure for miaosha_order_info
-- ----------------------------
DROP TABLE IF EXISTS `miaosha_order_info`;
CREATE TABLE `miaosha_order_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(32) NOT NULL COMMENT '订单号',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `goods_id` bigint(20) unsigned NOT NULL COMMENT '商品id',
  `goods_name` varchar(128) NOT NULL COMMENT '商品名称',
  `goods_img` varchar(128) NOT NULL COMMENT '商品图片',
  `goods_price` bigint(20) unsigned NOT NULL COMMENT '商品价格',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=474 DEFAULT CHARSET=utf8mb4 COMMENT='订单信息表';

-- ----------------------------
-- Table structure for miaosha_user
-- ----------------------------
DROP TABLE IF EXISTS `miaosha_user`;
CREATE TABLE `miaosha_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `mobile` varchar(16) NOT NULL COMMENT '手机号码',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_mobile` (`mobile`) USING BTREE COMMENT '手机号码索引'
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
