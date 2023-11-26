# miaosha-server

> 服务端，技术栈：Golang、Gin、SQLX、MySQL、Redis 等。

## Build Setup

```shell
# 1、数据库导入 miaosha.sql 脚本

# 2、修改 conf.example.ini 配置信息并将文件重名为 conf.ini

# 3、如果是 docker 环境，执行下面脚本
./deploy.sh

# 4、如果非 docker 环境，执行下面命令
go run main.go

```

