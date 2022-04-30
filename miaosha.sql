/*
Navicat MySQL Data Transfer

Source Server         : 192.168.100.1
Source Server Version : 50651
Source Host           : 192.168.100.1:3306
Source Database       : miaosha

Target Server Type    : MYSQL
Target Server Version : 50651
File Encoding         : 65001

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
-- Records of miaosha_order_info
-- ----------------------------
INSERT INTO `miaosha_order_info` VALUES ('422', '20210331105337823348', '1', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-03-31 18:53:37', '2021-03-31 18:57:04');
INSERT INTO `miaosha_order_info` VALUES ('423', '20210331105715598261', '1', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-03-31 18:57:15', '2021-03-31 19:27:15');
INSERT INTO `miaosha_order_info` VALUES ('424', '20210331190236658081', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-03-31 19:02:36', '2021-03-31 19:32:36');
INSERT INTO `miaosha_order_info` VALUES ('425', '20210331191536969430', '2', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-03-31 19:15:36', '2021-03-31 19:16:36');
INSERT INTO `miaosha_order_info` VALUES ('426', '20210401100527697341', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-01 10:05:27', '2021-04-01 10:06:27');
INSERT INTO `miaosha_order_info` VALUES ('427', '20210401101058383444', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-01 10:10:58', '2021-04-01 10:11:58');
INSERT INTO `miaosha_order_info` VALUES ('428', '20210401140645822741', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-01 14:06:45', '2021-04-01 14:07:45');
INSERT INTO `miaosha_order_info` VALUES ('429', '20210404140333540847', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-04 14:03:33', '2021-04-04 14:04:33');
INSERT INTO `miaosha_order_info` VALUES ('430', '20210405110254976044', '10', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-05 11:02:54', '2021-04-05 11:03:09');
INSERT INTO `miaosha_order_info` VALUES ('431', '20210405110327830840', '10', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-05 11:03:27', '2021-04-05 11:04:27');
INSERT INTO `miaosha_order_info` VALUES ('432', '20210405113632295430', '10', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-05 11:36:33', '2021-04-05 11:37:33');
INSERT INTO `miaosha_order_info` VALUES ('433', '20210406141427301520', '11', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-06 14:14:27', '2021-04-06 14:15:10');
INSERT INTO `miaosha_order_info` VALUES ('434', '20210406141530063370', '11', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-06 14:15:30', '2021-04-06 14:16:30');
INSERT INTO `miaosha_order_info` VALUES ('435', '20210406141755897667', '11', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-06 14:17:55', '2021-04-06 14:18:55');
INSERT INTO `miaosha_order_info` VALUES ('436', '20210407213942786919', '12', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-07 21:39:42', '2021-04-07 21:40:42');
INSERT INTO `miaosha_order_info` VALUES ('437', '20210407214110447892', '12', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-07 21:41:10', '2021-04-07 21:42:10');
INSERT INTO `miaosha_order_info` VALUES ('438', '20210407214233643845', '12', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-07 21:42:33', '2021-04-07 21:43:33');
INSERT INTO `miaosha_order_info` VALUES ('439', '20210408101235781878', '13', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-08 10:12:35', '2021-04-08 10:13:35');
INSERT INTO `miaosha_order_info` VALUES ('440', '20210408105558900727', '14', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-08 10:55:58', '2021-04-08 10:56:58');
INSERT INTO `miaosha_order_info` VALUES ('441', '20210408190117716284', '15', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-08 19:01:17', '2021-04-08 19:02:17');
INSERT INTO `miaosha_order_info` VALUES ('442', '20210408193437768940', '16', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-08 19:34:37', '2021-04-08 19:35:37');
INSERT INTO `miaosha_order_info` VALUES ('443', '20210409154115569904', '11', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-09 15:41:15', '2021-04-09 15:42:15');
INSERT INTO `miaosha_order_info` VALUES ('444', '20210409154227790603', '11', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-09 15:42:27', '2021-04-09 15:43:27');
INSERT INTO `miaosha_order_info` VALUES ('445', '20210411151544582029', '11', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-11 15:15:44', '2021-04-11 15:16:44');
INSERT INTO `miaosha_order_info` VALUES ('446', '20210412105922754338', '17', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-12 10:59:22', '2021-04-12 11:00:22');
INSERT INTO `miaosha_order_info` VALUES ('447', '20210414200706313888', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-14 20:07:06', '2021-04-14 20:08:06');
INSERT INTO `miaosha_order_info` VALUES ('448', '20210414200819096236', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-14 20:08:19', '2021-04-14 20:09:19');
INSERT INTO `miaosha_order_info` VALUES ('449', '20210414200848814888', '1', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-14 20:08:48', '2021-04-14 20:09:48');
INSERT INTO `miaosha_order_info` VALUES ('450', '20210415102539352800', '11', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-15 10:25:39', '2021-04-15 10:26:04');
INSERT INTO `miaosha_order_info` VALUES ('451', '20210415102611389437', '11', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-15 10:26:11', '2021-04-15 10:27:11');
INSERT INTO `miaosha_order_info` VALUES ('452', '20210416161300357830', '18', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-16 16:13:00', '2021-04-16 16:14:00');
INSERT INTO `miaosha_order_info` VALUES ('453', '20210416161319319952', '18', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-16 16:13:19', '2021-04-16 16:13:45');
INSERT INTO `miaosha_order_info` VALUES ('454', '20210419221227728284', '19', '8', '棒棒糖（猕猴桃味）', '/static/bbt.jpg', '10', '-1', '2021-04-19 22:12:27', '2021-04-19 22:12:37');
INSERT INTO `miaosha_order_info` VALUES ('455', '20210419221314211165', '19', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-19 22:13:14', '2021-04-19 22:13:45');
INSERT INTO `miaosha_order_info` VALUES ('456', '20210421134546022758', '1', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-21 13:45:46', '2021-04-21 13:46:46');
INSERT INTO `miaosha_order_info` VALUES ('457', '20210425153524592228', '20', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-25 15:35:24', '2021-04-25 15:35:51');
INSERT INTO `miaosha_order_info` VALUES ('458', '20210425153607397891', '20', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-04-25 15:36:07', '2021-04-25 15:37:07');
INSERT INTO `miaosha_order_info` VALUES ('459', '20210425153728091642', '20', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-04-25 15:37:28', '2021-04-25 15:38:28');
INSERT INTO `miaosha_order_info` VALUES ('460', '20210425180816468856', '21', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-04-25 18:08:16', '2021-04-25 18:09:16');
INSERT INTO `miaosha_order_info` VALUES ('461', '20210617101854232811', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-06-17 10:18:54', '2021-06-17 10:19:55');
INSERT INTO `miaosha_order_info` VALUES ('462', '20210618152652474770', '23', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-06-18 15:26:52', '2021-06-18 15:27:26');
INSERT INTO `miaosha_order_info` VALUES ('463', '20210623212121208268', '24', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-06-23 21:21:21', '2021-06-23 21:22:21');
INSERT INTO `miaosha_order_info` VALUES ('464', '20210628221556745865', '25', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-06-28 22:15:56', '2021-06-28 22:16:29');
INSERT INTO `miaosha_order_info` VALUES ('465', '20210628221644592541', '25', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-06-28 22:16:44', '2021-06-28 22:17:44');
INSERT INTO `miaosha_order_info` VALUES ('466', '20210628221803509647', '25', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-06-28 22:18:04', '2021-06-28 22:19:04');
INSERT INTO `miaosha_order_info` VALUES ('467', '20210629162355373737', '25', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-06-29 16:23:55', '2021-06-29 16:24:55');
INSERT INTO `miaosha_order_info` VALUES ('468', '20211101170452969337', '1', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-11-01 17:04:52', '2021-11-01 17:05:52');
INSERT INTO `miaosha_order_info` VALUES ('469', '20211104202355506778', '38', '7', '棒棒糖（西瓜味）', '/static/bbt1.jpg', '30', '-1', '2021-11-04 20:23:55', '2021-11-04 20:24:55');
INSERT INTO `miaosha_order_info` VALUES ('470', '20211105101314131331', '39', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-11-05 10:13:14', '2021-11-05 10:13:35');
INSERT INTO `miaosha_order_info` VALUES ('471', '20211105101350781434', '39', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-11-05 10:13:50', '2021-11-05 10:14:50');
INSERT INTO `miaosha_order_info` VALUES ('472', '20211121003840673526', '41', '9', '棒棒糖（柠檬味）', '/static/bbt1.jpg', '15', '-1', '2021-11-21 00:38:40', '2021-11-21 00:39:40');
INSERT INTO `miaosha_order_info` VALUES ('473', '20211121003940670987', '41', '6', '棒棒糖（原味）', '/static/bbt.jpg', '20', '-1', '2021-11-21 00:39:40', '2021-11-21 00:40:40');

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

-- ----------------------------
-- Records of miaosha_user
-- ----------------------------
INSERT INTO `miaosha_user` VALUES ('1', '18800000001', '2021-01-05 16:49:50', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('2', '18800000002', '2021-01-06 16:56:47', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('3', '18800000003', '2021-01-06 18:18:38', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('4', '18800000004', '2021-01-07 14:27:53', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('5', '18800000005', '2021-01-09 15:32:33', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('6', '18800000006', '2021-01-15 16:02:49', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('7', '18800000007', '2021-01-15 18:24:00', '0000-00-00 00:00:00');
INSERT INTO `miaosha_user` VALUES ('8', '18800000008', '2021-02-02 16:39:38', '2021-02-02 16:39:41');
INSERT INTO `miaosha_user` VALUES ('9', '18800000009', '2021-03-03 13:54:29', '2021-03-03 13:54:29');
INSERT INTO `miaosha_user` VALUES ('10', '18742014997', '2021-04-05 11:02:51', '2021-04-05 11:02:51');
INSERT INTO `miaosha_user` VALUES ('11', '17362574466', '2021-04-06 14:14:24', '2021-04-06 14:14:24');
INSERT INTO `miaosha_user` VALUES ('12', '13800138000', '2021-04-07 20:53:22', '2021-04-07 20:53:22');
INSERT INTO `miaosha_user` VALUES ('13', '15955125519', '2021-04-08 10:10:42', '2021-04-08 10:10:42');
INSERT INTO `miaosha_user` VALUES ('14', '13510948631', '2021-04-08 10:55:54', '2021-04-08 10:55:54');
INSERT INTO `miaosha_user` VALUES ('15', '15261561703', '2021-04-08 19:01:15', '2021-04-08 19:01:15');
INSERT INTO `miaosha_user` VALUES ('16', '15312353122', '2021-04-08 19:34:36', '2021-04-08 19:34:36');
INSERT INTO `miaosha_user` VALUES ('17', '18788910523', '2021-04-12 10:59:19', '2021-04-12 10:59:19');
INSERT INTO `miaosha_user` VALUES ('18', '13043490706', '2021-04-16 16:12:47', '2021-04-16 16:12:47');
INSERT INTO `miaosha_user` VALUES ('19', '15186828064', '2021-04-19 22:12:23', '2021-04-19 22:12:23');
INSERT INTO `miaosha_user` VALUES ('20', '15159703397', '2021-04-25 15:35:03', '2021-04-25 15:35:03');
INSERT INTO `miaosha_user` VALUES ('21', '18501733868', '2021-04-25 18:08:12', '2021-04-25 18:08:12');
INSERT INTO `miaosha_user` VALUES ('22', '18705692120', '2021-05-12 19:44:22', '2021-05-12 19:44:22');
INSERT INTO `miaosha_user` VALUES ('23', '15922941760', '2021-06-18 15:26:50', '2021-06-18 15:26:50');
INSERT INTO `miaosha_user` VALUES ('24', '13060618257', '2021-06-23 21:21:11', '2021-06-23 21:21:11');
INSERT INTO `miaosha_user` VALUES ('25', '17806289060', '2021-06-28 22:15:45', '2021-06-28 22:15:45');
INSERT INTO `miaosha_user` VALUES ('26', '13715365208', '2021-07-02 23:39:15', '2021-07-02 23:39:15');
INSERT INTO `miaosha_user` VALUES ('27', '13042030777', '2021-07-16 10:39:17', '2021-07-16 10:39:17');
INSERT INTO `miaosha_user` VALUES ('28', '13588878963', '2021-07-21 23:02:15', '2021-07-21 23:02:15');
INSERT INTO `miaosha_user` VALUES ('29', '13386488912', '2021-07-28 15:48:50', '2021-07-28 15:48:50');
INSERT INTO `miaosha_user` VALUES ('30', '15239138610', '2021-08-07 16:10:19', '2021-08-07 16:10:19');
INSERT INTO `miaosha_user` VALUES ('31', '13561373631', '2021-08-20 22:24:07', '2021-08-20 22:24:07');
INSERT INTO `miaosha_user` VALUES ('32', '15346975431', '2021-08-23 16:14:31', '2021-08-23 16:14:31');
INSERT INTO `miaosha_user` VALUES ('33', '12345678911', '2021-08-24 11:38:30', '2021-08-24 11:38:30');
INSERT INTO `miaosha_user` VALUES ('34', '13036605003', '2021-09-12 19:01:35', '2021-09-12 19:01:35');
INSERT INTO `miaosha_user` VALUES ('35', '15583875995', '2021-09-12 19:02:51', '2021-09-12 19:02:51');
INSERT INTO `miaosha_user` VALUES ('36', '18011411064', '2021-09-26 18:52:23', '2021-09-26 18:52:23');
INSERT INTO `miaosha_user` VALUES ('37', '15800000000', '2021-10-07 00:51:46', '2021-10-07 00:51:46');
INSERT INTO `miaosha_user` VALUES ('38', '15671248002', '2021-11-04 20:23:52', '2021-11-04 20:23:52');
INSERT INTO `miaosha_user` VALUES ('39', '17736775050', '2021-11-05 10:13:08', '2021-11-05 10:13:08');
INSERT INTO `miaosha_user` VALUES ('40', '13025566561', '2021-11-05 11:23:18', '2021-11-05 11:23:18');
INSERT INTO `miaosha_user` VALUES ('41', '17895592870', '2021-11-21 00:38:34', '2021-11-21 00:38:34');
