# CloudDisk

本着从项目中了解go-zero的用发，所以跟着视频，写了部分内容笔记

> 轻量级云盘系统，基于go-zero、xorm实现。

## goctl
### goctl简介

goctl是go-zero微服务框架下的代码生成工具。使用 goctl 可显著提升开发效率，让开发人员将时间重点放在业务开发上，其功能有：

- api服务生成
- rpc服务生成
- model代码生成
- 模板管理

很多人会把 ```goctl``` 读作 ```go-C-T-L```，这种是错误的念法，应参照 ```go control```来读。


### goctl安装
```shell
# goctl安装
go install github.com/zeromicro/go-zero/tools/goctl@latest
# 安装好之后查看版本号
goctl -v
```
### 单体core服务
goctl api是goctl中的核心模块之一，其可以通过.api文件一键快速生成一个api服务，说白了就是自动文件结构生成、文件生成、逻辑管理，极大的减少了代码量和项目逻辑的书写和考虑时间。



使用到的命令
```text
goctl api new core
# 先安装依赖
go mod tidy
# 再运行
go run core.go -f etc/greet-api.yaml
# windows 稍微有点区别
go run core.go -f .\etc\core-api.yaml
# 在.api文件中定义好结构体后，使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```
