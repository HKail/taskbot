# Task Bot

Task Bot 是一款用于提醒和记录用户制定的每日任务完成情况的 QQ 频道机器人～👾👾

## Usage

1. 拉取代码到本地

    ```shell
    # 拉取工程代码
    $ git clone https://github.com/HKail/taskbot.git
    $ cd taskbot
    ```

2. 修改配置文件

    将`conf/prod/config.yml`中的机器人配置和数据库配置修改为自己的配置

3. 执行`env/sql/*`下的 SQL 文件进行建表

4. 部署

    1. 自行编译`app/cmd/taskbot/main.go`文件
    2. 直接使用利用根目录的`Dockerfile`文件进行容器化部署

## Feature

- [x] 每日签到
- [ ] 任务制定
- [ ] 任务打卡
- [ ] 任务通知
- [ ] 任务完成统计
- [ ] ...

## 已支持指令

- /打卡
- ...

## TODO

- 依赖注入引入
- 日志完善
- 单测完善
- 增加缓存
- 增加监控、报警
- WS 链接支持分片

## 打卡功能设计（指令：/打卡）

### 存储结构

```sql
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='频道用户签到表';
```

使用 bitmap 的方式来存储用户每个月的签到天数情况，由于一个月最多不超过 31 天，因此使用 32 位的无符号整型就足够表示用户一个月的签到情况。

## 体验方式

1. 扫描以下二维码进入测试频道

![image-20220325150352667](README.assets/image-20220325150352667.png)

2. 在频道中对小花机器人发送"/打卡"指令进行打卡