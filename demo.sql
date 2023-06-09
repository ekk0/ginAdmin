/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 10/06/2023 01:35:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `age` int(2) NOT NULL,
  `sex` int(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'chenwebo', 1, 2);
INSERT INTO `user` VALUES (2, '陈短发的', 11, 1);
INSERT INTO `user` VALUES (3, '张三', 18, 1);
INSERT INTO `user` VALUES (4, '张三', 18, 1);
INSERT INTO `user` VALUES (5, '张三', 18, 1);

SET FOREIGN_KEY_CHECKS = 1;
