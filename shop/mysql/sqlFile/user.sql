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
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `Info`;
CREATE TABLE `Info` (
  `id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '�û�id��������',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '�û�����',
  `real_name` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '�û���ʵ����',
  `role_id` tinyint(20) NOT NULL DEFAULT '1' COMMENT '�û���ɫ��1��ʾ��ͨ�û�',
  `password` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '�û�����',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '�û��绰',
  `balance` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '�û����',
  `status` tinyint(20) NOT NULL DEFAULT '1' COMMENT '�û�״̬��1��ʾ������0��ʾ��ͣ',
  `created` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '����ʱ��',
  `updated` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100037 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `Info` VALUES ('oUT385ZLmRr6R_a9xKSfSW9SekYI', 'ʱ�ⲻ����', 'admin', 1,'admin@123','13300003333',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('aUT385ZLmRr6R_a9xKSfSW9SekUK', '�����޺�', 'HAHA', 1,'admin@123','13708061236',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('aUT385ZLmRr6R_dsadasdW9SekUK', '�����ͽ', '����', 1,'admin@123','18708061231',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('aUT385Zdasdas_a9xKSfSW9SekUK', '��ëʨ��', '����', 1,'admin@123','15108061236',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('dasdadsadacxz_a9xKSfSW9SekUK', '�������', '����', 1,'admin@123','17708061236',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('aUT385ZLmRr6R_dfmskdmfsdfksd', '�������', '����', 1,'admin@123','13708061238',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('dasdsadvcbxgd_dsadDkdfFDSFdd', '��ң��', '����', 1,'admin@123','13708061232',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info` VALUES ('cUT234ZLmRr7R_a0xKSfSW1SekUK', '��ë��', '�ܰ�', 1,'admin@123','13708061233',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('dUT381ZLmRr6R_a1xKSfSW2SekUK', '�Ƕ���', '���', 1,'admin@123','13708061235',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('eUT482ZLmRr6R_a2xKSfSW3SekUK', '��ʱ��', '��ʵ', 1,'admin@123','13708061239',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('fUT583ZLmRr6R_a3xKSfSW4SekUK', '�콭��', 'ʮһ', 1,'admin@123','15508061236',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('hUT684ZLmRr6R_a4xKSfSW5SekUK', 'С��ɵ�', '��Ѱ��', 1,'admin@123','13708061136',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('iUT786ZLmRr6R_a5xKSfSW6SekUK', '��ң��ʦ', '������', 1,'admin@123','13708061256',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('jUT887ZLmRr6R_a6xKSfSW7SekUK', '����ʮ����', '�Ƿ�', 1,'admin@123','13708065555',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('hUT988ZLmRr6R_a7xKSfSW8SekUK', '����', '����', 1,'admin@123','13708061111',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('qUT089ZLmRr6R_a8xKSfSU1SekUK', '��������', '����', 1,'admin@123','13708062235',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('wUT171ZLmRr6R_b1xKSfSU2SekUK', 'ѩɽ�ɺ�', '���', 1,'admin@123','13708062233',0, 1, '2021-12-03 10:35:02', NULL);
INSERT INTO `Info`  VALUES ('eUT272ZLmRr6R_b2xKSfSU3SekUK', '�ܶ���', '�¼���', 1,'admin@123','13708061666',0, 1, '2021-12-03 10:35:02', NULL);
