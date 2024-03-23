# Cron Admin

基于Golang + React.js 的定时任务管理平台

https://zhuanlan.zhihu.com/p/480841357

- 后端代码：https://github.com/mouday/cron-admin
- 前端代码：https://github.com/mouday/cron-admin-web

如果需要前端代码，issue中 留下github用户名，我会将代码权限分享给你

运行逻辑

1. 启动Cron Admin
2. 配置定时任务
3. Cron Admin定时任务启动，调用`用户api`，执行用户自定义定时任务
4. 任务执行中，调用`Cron Admin api`，更新定时任务状态，推送执行日志
5. 定时任务状态更新完毕，调用`Cron Admin api`，更新定时任务执行结果

可以配置任何编程语言实现定任务的执行，例如：

- Python: Flask、Django、FastAPI
- Golang: Gin
- Node.js: Koa、Express
- Java: Spring Boot

Cron Admin 仅负责定时任务的调度，不负责定时任务的执行

## 安装

下载对应平台的可以执行文件，运行即可，不需要安装任何依赖

