/*
 Navicat Premium Data Transfer

 Source Server         : mariaDB
 Source Server Type    : MariaDB
 Source Server Version : 100508
 Source Host           : localhost:3306
 Source Schema         : ginblog

 Target Server Type    : MariaDB
 Target Server Version : 100508
 File Encoding         : 65001

 Date: 12/24/2021 17:06:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`            bigint(20) UNSIGNED                                     NOT NULL AUTO_INCREMENT,
    `created_at`    datetime(3)                                             NULL     DEFAULT NULL,
    `title`         varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `desc`          varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL     DEFAULT NULL,
    `content`       longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL,
    `img`           varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL     DEFAULT NULL,
    `comment_count` bigint(20)                                              NOT NULL DEFAULT 0,
    `read_count`    bigint(20)                                              NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 574
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article`
VALUES (1, '2021-12-24 22:47:34.425', 'haha', 'hag',
        'test',
        'test', 2, 9);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`            bigint(20) UNSIGNED                                     NOT NULL AUTO_INCREMENT,
    `created_at`    datetime(3)                                             NULL DEFAULT NULL,
    `user_id`       bigint(20) UNSIGNED                                     NULL DEFAULT NULL,
    `article_id`    bigint(20) UNSIGNED                                     NULL DEFAULT NULL,
    `content`       varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `status`        tinyint(4)                                              NULL DEFAULT 2,
    `article_title` longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL,
    `username`      longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL,
    `title`         longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 13
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment`
VALUES (1, '2021-12-24 22:47:58.055', 1, 1, '测试测试', 1, NULL, 'admin', '');
INSERT INTO `comment`
VALUES (2, '2021-12-24 22:49:04.785', 2, 1, '测试测试', 1, NULL, 'weject', '');


-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint(20) UNSIGNED                                     NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)                                             NULL DEFAULT NULL,
    `username`   varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL,
    `password`   varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `role`       bigint(20)                                              NULL DEFAULT 2,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user`
VALUES (1, '2021-12-25 17:05:14.764', 'admin',
        '$2a$10$YGL5a9z7ykG6BWOo.XhJU.h8r98BD5IvAmLISBB9rFIefbDzrv58O', 1);

SET FOREIGN_KEY_CHECKS = 1;
