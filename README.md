## 项目介绍
该项目主要用于提供内容敏感词检测服务，提供 http 以及 grpc 调用方式。
采用分层模型，文本检测目前基于 ac自动机算法 和 百度提供的检测服务。
图片检测基于百度提供的图片检测服务。

## 环境依赖
go v1.17  
mongodb  
[百度智能云](https://cloud.baidu.com/solution/censoring?track=b07f3ea2be2cdc5cf7052e4f5b47af1c236b6019a950ac15&bd_vid=12333161148280427575e)

## 部署步骤
1. 安装go环境
2. 下载项目 
3. 配置mongodb
4. 开通百度智能云内容检测服务
5. 修改 conf 下配置文件


## 目录结构描述
```
├─app
│  ├─domain
│  │  ├─dtos
│  │  ├─models
│  │  │  └─mongo_models
│  │  ├─proto
│  │  ├─repositories
│  │  └─services
│  ├─grpc
│  ├─http
│  │  ├─controllers
│  │  └─middlewares
│  └─providers
├─bootstrap
├─conf                   //配置文件
├─configs
├─docs
├─pkg
│  └─utils
└─routes

```