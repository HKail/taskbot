CREATE TABLE `guild_users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID, 自增主键',
    `guild_id` bigint unsigned NOT NULL COMMENT '频道 ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户 ID',
    `experience` int unsigned NOT NULL DEFAULT '0' COMMENT '经验值',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_gid_uid` (`guild_id`,`user_id`) COMMENT '频道用户唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='频道用户表'