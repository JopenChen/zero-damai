# zero-damai
基于 go-zero 框架仿写的大麦 App 服务，旨在开发交流学习。

对应的后台服务 zero-damai-admin，可关注公众号回复获取项目权限：zero-damai-admin

## 依赖 go-zero 版本

| 工具名   | 版本     | 仓库地址                                  |
|-------|--------|---------------------------------------|
| goctl | v1.5.6 | https://github.com/zeromicro/go-zero  |

特别鸣谢 go-zero 开源社区，官方地址：https://go-zero.dev/

## 开发步骤
说明：goctl 代码生成模板部分修改过，目录为：template/1.5.6，文件名带有“_backup”后缀的说明有修改过。

1、配置 go-zero 开发环境，详情请参考：https://go-zero.dev/docs/tasks/installation/goctl

2、拉取本仓库代码
```shell
# 拉取
git clone https://github.com/JopenChen/zero-damai.git
# 切换目录
cd zero-damai
# 加载依赖包
go mod tidy
```

3、修改配置文件

4、启动 RPC 服务
```shell
go run service.go -f etc/service.yaml
```

5、启动 API 服务
```shell
go run gateway.go -f etc/service.yaml
```

6、测试

## 设计指标
1、以 1 千万活跃用户量为基准；

2、仅考虑 APP 客户端。

## 项目架构概述

## 目录结构

## 数据库设计

## 公众号：Golang 进阶栈
<img height="100" src="doc/images/公众号二维码.jpg" title="Golang 进阶栈公众号二维码" width="100" alt="公众号二维码"/>

1、API 文档开放地址, 对话框回复：API 文档；

2、微信交流群，对话框回复：zero-damai；

3、zero-damai-admin 项目权限，对话框回复：zero-damai-admin。




