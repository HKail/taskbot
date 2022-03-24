# Task Bot

一款用于任务通知的 QQ 频道机器人👾👾

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

## TODO

- 依赖注入引入
- 日志完善
- 单测完善
- 增加缓存
- 增加监控、报警
- WS 链接支持分片