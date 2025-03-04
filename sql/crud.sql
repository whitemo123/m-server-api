/*
 Navicat Premium Dump SQL

 Source Server         : crud
 Source Server Type    : MySQL
 Source Server Version : 80024 (8.0.24)
 Source Host           : 192.168.31.10:13306
 Source Schema         : crud

 Target Server Type    : MySQL
 Target Server Version : 80024 (8.0.24)
 File Encoding         : 65001

 Date: 04/03/2025 16:29:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint NOT NULL COMMENT '系统菜单ID',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父级系统菜单ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `sort` int NULL DEFAULT NULL COMMENT '菜单排序',
  `type` int NOT NULL COMMENT '菜单类型(1=目录 2=菜单 3=按钮)',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单路径',
  `alias` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单别名',
  `keep` int NULL DEFAULT 2 COMMENT '缓存(1=缓存 2=不缓存)',
  `tenant_id` bigint NULL DEFAULT 88888888 COMMENT '租户ID',
  `status` int NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` bigint NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1896393497456414720, NULL, '系统管理', 99, 1, '2', '/system', 'system', 2, 88888888, 1, 1895412907303243776, '2025-03-03 10:53:38', 1895412907303243776, '2025-03-03 10:56:51');
INSERT INTO `sys_menu` VALUES (1896394950698864640, 1896393497456414720, '账号管理', 1, 2, '1', '/system/user', 'systemUser', 2, 88888888, 1, 1895412907303243776, '2025-03-03 10:59:24', 1895412907303243776, '2025-03-03 10:59:24');
INSERT INTO `sys_menu` VALUES (1896395243444506624, 1896393497456414720, '角色管理', 2, 2, '1', '/system/role', 'systemRole', 2, 88888888, 1, 1895412907303243776, '2025-03-03 11:00:34', 1895412907303243776, '2025-03-03 11:00:34');
INSERT INTO `sys_menu` VALUES (1896395442741055488, 1896393497456414720, '菜单管理', 3, 2, '1', '/system/menu', 'systemMenu', 2, 88888888, 1, 1895412907303243776, '2025-03-03 11:01:22', 1895412907303243776, '2025-03-03 11:01:22');
INSERT INTO `sys_menu` VALUES (1896433627596591104, 1896394950698864640, '新增', 1, 3, '', '', 'admin:menu:create', 2, 88888888, 1, 1895412907303243776, '2025-03-03 13:33:06', 1895412907303243776, '2025-03-03 13:33:06');
INSERT INTO `sys_menu` VALUES (1896585636018655232, NULL, '推广', 1, 1, '1', '/promoter', 'promoter', 2, 88888888, 1, 1895412907303243776, '2025-03-03 23:37:07', 1895412907303243776, '2025-03-03 23:37:07');
INSERT INTO `sys_menu` VALUES (1896585782576025600, 1896585636018655232, '推广员', 1, 2, '1', '/promoter/list', 'promoterList', 2, 88888888, 1, 1895412907303243776, '2025-03-03 23:37:42', 1895412907303243776, '2025-03-03 23:37:42');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL COMMENT '主键',
  `tenant_id` bigint NULL DEFAULT 88888888 COMMENT '租户ID',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `status` int NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` bigint NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1895121477213097984, 88888888, '超级管理员', 1, 1895412907303243776, '2025-03-03 13:55:18', 1895412907303243776, '2025-03-03 23:39:36');

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色菜单关联表',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (6, 1895121477213097984, 1896393497456414720);
INSERT INTO `sys_role_menu` VALUES (7, 1895121477213097984, 1896394950698864640);
INSERT INTO `sys_role_menu` VALUES (8, 1895121477213097984, 1896395243444506624);
INSERT INTO `sys_role_menu` VALUES (9, 1895121477213097984, 1896395442741055488);
INSERT INTO `sys_role_menu` VALUES (10, 1895121477213097984, 1896433627596591104);
INSERT INTO `sys_role_menu` VALUES (11, 1895121477213097984, 1896585636018655232);
INSERT INTO `sys_role_menu` VALUES (12, 1895121477213097984, 1896585782576025600);

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant`  (
  `id` bigint NOT NULL,
  `tenant_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` bigint NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------
INSERT INTO `sys_tenant` VALUES (88888888, '默认租户', 1, 1895412907303243776, '2025-02-23 15:49:20', 1895412907303243776, '2025-02-27 14:40:43');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
  `password` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `status` int NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` bigint NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1895412907303243776, '88888888', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', '', 1, 1895412907303243776, '2025-02-28 17:57:07', 1895412907303243776, '2025-02-28 17:57:07');

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户角色关联主键ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `role_id` bigint NULL DEFAULT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (4, 1895412907303243776, 1895121477213097984);

SET FOREIGN_KEY_CHECKS = 1;
