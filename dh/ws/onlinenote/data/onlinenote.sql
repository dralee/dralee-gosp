
create database notedb;

use notedb;

CREATE TABLE `user` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户Id',
    `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
    `is_admin` bit NOT NULL DEFAULT 0 COMMENT '是否是管理员',
    `is_enabled` bit NOT NULL DEFAULT 0 COMMENT '是否启用',
    `creation_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `last_modification_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `note` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '笔记Id',
    `name` varchar(32) NOT NULL DEFAULT '' COMMENT '标题',
    `content` text NOT NULL COMMENT '内容',
    `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户Id',
    `creation_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `last_modification_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` (`id`, `username`, `password`, `is_admin`, `is_enabled`, `creation_time`) VALUES (1, 'lee', 'Lee.123', 1, 1, 0);

