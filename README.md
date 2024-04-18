# Cron Admin

一个零依赖的基于Golang + React.js 的定时任务管理平台

基本逻辑执行任务：通过api同步调用，同步返回

Cron Admin 仅负责定时任务的调度，不负责定时任务的执行

可以配置任何编程语言实现定任务的执行，例如：

- Python: Flask、Django、FastAPI
- Golang: Gin
- Node.js: Koa、Express
- Java: Spring Boot
- ...

## 项目截图

![](https://cdn.jsdelivr.net/gh/mouday/img/2024/04/18/t3qfbrg.png)

![](https://cdn.jsdelivr.net/gh/mouday/img/2024/04/18/sfofgut.png)

## 安装

下载对应平台的可以执行文件，运行即可，不需要安装任何依赖

1、下载适合所在运行平台的二进制文件

[https://github.com/mouday/cron-admin/releases](https://github.com/mouday/cron-admin/releases)

2、解压文件并启动服务

```bash
# macos: 
./cron-admin

# linux: 
./cron-admin

# windows: 
cron-admin.exe
```

启动后控制台会输出如下内容

```
username: admin
password: 8),fua+jMN
********************************************
* Welcome Use Cron Admin  v1.0.0
* server run at:  http://127.0.0.1:8000
* issue: https://github.com/mouday/cron-admin
********************************************
```

查看控制台输出的账号密码，即可登录应用

默认地址：http://127.0.0.1:8000

## 配置文件

可以使用默认配置，也可以自定义配置

```bash
cp env.example .env
```

配置说明

```bash
# == 应用配置 ==
# 运行模式 debug test release (默认：release)
GIN_MODE=release

# 监听端口 (默认：127.0.0.1:8000）
APP_RUN_ADDRESS=127.0.0.1:8000

# == 账号配置 ==
# 管理员账号
APP_ADMIN_USERNAME=admin
# 管理员密码
APP_ADMIN_PASSWORD=
```
