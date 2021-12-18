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
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '��ַid',
  `user_id` varchar(200) DEFAULT NULL COMMENT '�û�id',
  `name` varchar(255) DEFAULT NULL COMMENT '�ջ�������',
  `mobile` char(16) DEFAULT NULL COMMENT '�ֻ���',
  `postal_code` int(6) DEFAULT NULL COMMENT '��������',
  `province` char(30) DEFAULT NULL COMMENT 'ʡ',
  `city` char(30) DEFAULT NULL COMMENT '����',
  `district` char(30) DEFAULT NULL COMMENT '��/��',
  `detailed_address` varchar(200) DEFAULT NULL COMMENT '��ϸ��ַ',
  `is_default` tinyint(1) DEFAULT NULL COMMENT '1ΪĬ�ϣ�0Ϊ��Ĭ��',
  `created` char(20) DEFAULT NULL COMMENT '����ʱ��',
  `updated` char(20) DEFAULT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1098 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1096, 'oUT385ZLmRr6R_a9xKSfSW9SekYI', 'ʱ�ⲻ����', '13300003333', 303033, '�㽭ʡ', '������', '������', '����С��3��303', 1, '2021-11-09 16:28:37', '2021-11-15 18:13:44');
INSERT INTO `address` VALUES (1097, 'aUT385ZLmRr6R_a9xKSfSW9SekUK', '�����޺�', '13708061236', 303045, '�㽭ʡ', '������', '�ϳ���', 'xxxx', 2, '2021-11-15 18:13:26', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
