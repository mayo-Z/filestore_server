# 项目介绍

基于golang实现的一种分布式云存储服务


# 实现功能
- 简单的文件上传服务
- mysql存储文件元数据
- 账号系统, 注册/登录/查询用户或文件数据
- 基于帐号的文件操作接口
- 文件秒传功能
- 文件分块上传/断点续传功能
- 搭建及使用hdfs进行私有云存储
- 使用阿里云OSS对象进行公有云存储
- 使用RabbitMQ实现异步任务队列(todo)

# 项目预览

[![image.png](https://i.postimg.cc/d1R8YRxk/image.png)](https://postimg.cc/zbBLwh58)

# 代码运行

- 首先git clone 项目
```
git clone https://github.com/uptocorrupt/filestore_server.git
```

- 下载类库依赖

```
go mod tidy
```

- 创建 db 并导入数据

```
mysql -u root -p -e "CREATE DATABASE fileserver"
mysql -u root -p go_gateway < table.sql
```

- 调整 配置文件

修改mysql,redis,本地连接的端口和账号密码。

本地连接的端口默认为18080。(如修改端口需同步修改static/view里的前端代码)

- 运行

```
go run main.go
```

- 网页查看
注册：http://127.0.0.1:18080/static/view/signup.html

登录：http://127.0.0.1:18080/static/view/signin.html

主页：http://127.0.0.1:18080/static/view/home.html
(ps:需先登录获取token,登录后自动跳转该页面)

上传页面：http://127.0.0.1:18080/static/view/



