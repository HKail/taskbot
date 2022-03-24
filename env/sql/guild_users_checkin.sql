CREATE TABLE `guild_users_checkin` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID, 自增主键',
    `guild_id` bigint unsigned NOT NULL COMMENT '频道 ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户 ID',
    `yearmonth` int unsigned NOT NULL COMMENT '年月, 如: 200601',
    `days` int unsigned NOT NULL DEFAULT '0' COMMENT '当月签到情况, bitmap 格式, 32 位足够表示一个月的 31 天',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_gid_uid_yearmonth` (`guild_id`,`user_id`,`yearmonth`) COMMENT '用户月签到记录唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='频道用户签到表'