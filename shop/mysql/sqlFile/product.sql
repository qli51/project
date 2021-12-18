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
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '��Ʒ���',
  `category_id` bigint(20) DEFAULT NULL COMMENT '��Ŀ���',
  `title` varchar(50) DEFAULT NULL COMMENT '��Ʒ����',
  `description` varchar(80) DEFAULT NULL COMMENT '��Ʒ����',
  `price` decimal(20,2) DEFAULT NULL COMMENT '��Ʒ�۸�',
  `amount` int(10) DEFAULT NULL COMMENT '��Ʒ����',
  `sales` int(10) DEFAULT NULL COMMENT '��Ʒ����',
  `main_image` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '��Ʒ��ͼ',
  `delivery` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '��Ʒ����',
  `assurance` varchar(30) DEFAULT NULL COMMENT '��Ʒ����',
  `name` varchar(30) DEFAULT NULL COMMENT '��Ʒ����',
  `weight` double(20,0) DEFAULT NULL COMMENT '��Ʒ����',
  `brand` varchar(10) DEFAULT NULL COMMENT '��ƷƷ��',
  `origin` varchar(80) DEFAULT NULL COMMENT '��Ʒ����',
  `shelf_life` int(20) DEFAULT NULL COMMENT '��Ʒ������',
  `net_weight` double(20,0) DEFAULT NULL COMMENT '��Ʒ������',
  `use_way` varchar(20) DEFAULT NULL COMMENT 'ʹ�÷�ʽ',
  `packing_way` varchar(20) DEFAULT NULL COMMENT '��װ��ʽ',
  `storage_conditions` varchar(20) DEFAULT NULL COMMENT '�洢����',
  `detail_image` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '����ͼƬ',
  `status` int(10) DEFAULT NULL COMMENT '��Ʒ״̬',
  `created` varchar(50) DEFAULT NULL COMMENT '����ʱ��',
  `updated` varchar(50) DEFAULT NULL COMMENT '����ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (9, 1098, '����300g', '���ʲ���', 6.00, 351, 330, 'http://localhost:8000/image/IMG_1296.JPG', '�����µ���Ԥ��3Сʱ���ʹ�', '֧��6Сʱ���˻���', '����', 300, '��Ӫ', '����ʡ�人��', 2, 300, 'ʳ��', '��װ', '����', 'http://localhost:8000/image/IMG_1298.JPG', 2, '2021-11-05 10:17:05', '2021-11-09 20:43:45');
INSERT INTO `product` VALUES (10, 1100, '��������������ѣ�400g', '֭ˮ��ӯ ����ɿ�', 4.00, 680, 120, 'http://localhost:8000/image/IMG_1299.JPG', '�����µ���Ԥ��2Сʱ�ʹ�', '֧��6Сʱ���˻�', '�����������ѣ�', 400, '��Ӫ', '����ʡ�人��', 3, 400, 'ʳ��', '��װ', '����', 'http://localhost:8000/image/IMG_1300.JPG', 1, '2021-11-05 10:29:39', '');
INSERT INTO `product` VALUES (11, 1102, '�����450g', '�������� ΢�����', 4.00, 233, 36, 'http://localhost:8000/image/IMG_1301.JPG', '���ڷ�����Ԥ��3Сʱ�ʹ�', '֧��6Сʱ���˻�', '�����', 450, '��Ӫ', '�㶫ʡ������', 7, 450, '����ʳ��', '��װ', '����', 'http://localhost:8000/image/IMG_1302.JPG', 1, '2021-11-05 10:40:59', '2021-11-05 21:23:25');
INSERT INTO `product` VALUES (12, 1103, '����300g', '�������� ���ü�', 4.00, 106, 300, 'http://localhost:8000/image/IMG_1303.JPG', '�������أ�Ԥ��2Сʱ�ʹ�', '��֧���˻���', '����', 300, '��Ӫ', '����ʡ�人��', 3, 300, 'ʳ��', '��װ', '���', 'http://localhost:8000/image/IMG_1304.JPG', 1, '2021-11-05 10:49:00', '2021-11-05 21:04:53');
INSERT INTO `product` VALUES (13, 1104, '�����ܲ�600g', '�������� ��ζ���ɵ�', 5.00, 355, 550, 'http://localhost:8000/image/IMG_1305.JPG', '�����µ���Ԥ��2Сʱ���ʹ�', '֧��5Сʱ���˻�', '�����ܲ�', 600, '��Ӫ', '����ʡ�ϲ���', 5, 600, 'ʳ��', '��װ', '����', 'http://localhost:8000/image/IMG_1306.JPG', 1, '2021-11-05 14:27:12', '');
INSERT INTO `product` VALUES (14, 1105, '��ײ�300g', '���ʴ�ײ� ���ó�', 5.00, 124, 8802, 'http://localhost:8000/image/IMG_1307.JPG', '�����µ���Ԥ��2Сʱ�ʹ�', '֧��3Сʱ���˻�', '��ײ�', 300, '��Ӫ', '����ʡ�人��', 2, 300, 'ʳ��', '��װ', '����', 'http://localhost:8000/image/IMG_1308.JPG', 1, '2021-11-05 14:33:09', '2021-11-05 21:02:58');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
