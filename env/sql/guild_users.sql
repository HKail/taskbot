CREATE TABLE `guild_users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID，自增主键',
    `guild_id` bigint unsigned NOT NULL COMMENT '频道 ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户 ID',
    `cont_checkin_cnt` int unsigned NOT NULL DEFAULT '0' COMMENT '连续签到次数',
    `experience` int unsigned NOT NULL DEFAULT '0' COMMENT '经验值',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_guild_users_gid_uid` (`guild_id`,`user_id`) COMMENT '频道用户唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='频道用户表'