/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : sf_mall

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '��Ŀid',
  `name` char(50) DEFAULT NULL COMMENT '��Ŀ����',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '������Ŀid',
  `level` int(5) DEFAULT NULL COMMENT '��Ŀ�㼶',
  `sort` int(5) DEFAULT NULL COMMENT '��Ŀ����',
  `created` char(20) DEFAULT NULL COMMENT '����ʱ��',
  `updated` char(20) DEFAULT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1106 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1077, 'Ҷ����', 1, 1, 100, '2021-11-03 10:10:05', '');
INSERT INTO `category` VALUES (1078, 'С�ײ�', 1077, 2, 80, '2021-11-03 10:10:21', '');
INSERT INTO `category` VALUES (1079, '�۲�', 1077, 2, 35, '2021-11-03 10:11:38', '');
INSERT INTO `category` VALUES (1080, '���Ĳ�', 1077, 2, 32, '2021-11-03 10:12:16', '');
INSERT INTO `category` VALUES (1081, '��ѿ��', 1, 1, 60, '2021-11-03 10:12:46', '');
INSERT INTO `category` VALUES (1082, '��ѿ��', 1081, 2, 35, '2021-11-03 10:12:55', '');
INSERT INTO `category` VALUES (1083, '�Ϲ���', 1, 1, 34, '2021-11-03 10:24:58', '');
INSERT INTO `category` VALUES (1084, '�ƹ�', 1083, 2, 55, '2021-11-03 10:25:08', '');
INSERT INTO `category` VALUES (1085, '����', 1083, 2, 77, '2021-11-03 10:25:32', '');
INSERT INTO `category` VALUES (1086, '�ѹ���', 1, 1, 55, '2021-11-03 10:25:55', '');
INSERT INTO `category` VALUES (1087, '����', 1086, 2, 67, '2021-11-03 10:26:05', '');
INSERT INTO `category` VALUES (1088, '������', 1, 1, 12, '2021-11-03 10:26:19', '');
INSERT INTO `category` VALUES (1089, '���ܲ�', 1088, 2, 33, '2021-11-03 10:26:29', '');
INSERT INTO `category` VALUES (1090, '������', 1, 1, 35, '2021-11-03 10:26:40', '');
INSERT INTO `category` VALUES (1091, '������', 1090, 2, 12, '2021-11-03 10:26:47', '');
INSERT INTO `category` VALUES (1092, '������', 1, 1, 35, '2021-11-03 10:27:00', '');
INSERT INTO `category` VALUES (1093, '�㶹', 1092, 2, 58, '2021-11-03 10:27:13', '');
INSERT INTO `category` VALUES (1094, '������', 1, 1, 55, '2021-11-03 10:27:48', '');
INSERT INTO `category` VALUES (1095, '�㹽', 1094, 2, 12, '2021-11-03 10:27:58', '');
INSERT INTO `category` VALUES (1096, 'ˮ����', 1, 1, 12, '2021-11-03 10:28:44', '');
INSERT INTO `category` VALUES (1097, '�ϲ�', 1096, 2, 12, '2021-11-03 10:28:47', '');
INSERT INTO `category` VALUES (1098, '����', 1077, 2, 55, '2021-11-05 09:44:49', '');
INSERT INTO `category` VALUES (1100, '������', 1086, 2, 39, '2021-11-05 10:23:49', '');
INSERT INTO `category` VALUES (1101, '������', 1, 1, 30, '2021-11-05 10:31:33', '');
INSERT INTO `category` VALUES (1102, '�����', 1101, 2, 55, '2021-11-05 10:32:04', '');
INSERT INTO `category` VALUES (1103, '����', 1077, 2, 29, '2021-11-05 10:44:47', '');
INSERT INTO `category` VALUES (1104, '�����ܲ�', 1088, 2, 36, '2021-11-05 14:20:39', '');
INSERT INTO `category` VALUES (1105, '��ײ�', 1077, 2, 80, '2021-11-05 14:28:24', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
