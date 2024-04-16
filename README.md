# MyGOWork~~4~~ 5!

**bibi-demo** is a small video website backend using hertz(hz-gen、jwt、websocket)+gorm(mysql)+redis+oss(aliyun)

## deploy by docker(net=host)

(使用前请先关闭本机的mysql与redis服务)

(由于重构为rpc架构，暂时不能快速启动)
`快速启动`
```bash
#oss与email的配置需自行填写
mv config/config-example.yaml config/config.yaml
docker-compose up -d # 启动相关容器
docker build -t bibi-demo . # 构建镜像
docker run -d --net=host bibi-demo go run bibi # 运行程序
```

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了(**Header:Authorization格式**:Bearer {token})

本项目的构建历程：抄项目架构、抄结构体、看demo遇到不会的学一下然后继续抄...

真是一场酣畅淋漓的ctrl+c.jpg

对于结构体加密存储redis使用了msgp(就一个地方偷懒直接用了JSON) 

(commit 都是瞎写的不要在意...)

## 完成情况：
重构=0


## Recent:

进行rpc重构


## 重构相关:
### 0324
花了一整天学习rpc、kitex-demo、tiktok(west-2 online)

### 0325
优化了config(感谢强大的viper),添加了constants包

### 0326
以Register为例:
1. [api/biz/handler/api/user_handler.go](api/biz/handler/api/user_handler.go)暴露api，接收请求，打包发送至中转
2. [api/biz/rpc](api/biz/rpc)作为中转，向rpc服务器发送请求
3. [rpc/user/handler.go](rpc/user/handler.go)接收请求，作为rpc服务器中的handler

### 0330
学了一下用shell自动运行命令(太好用了)

实现了双token:创建了两个hertz_jwt中间件，一个负责access，一个负责refresh，并添加一个用于get access by refresh的接口

### 0416

在video服务接入es,并使用钩子自动上传([参考仓库链接](https://github.com/CocaineCong/eslogrus))

同时video服务展示了rpc架构下并发远程调用其他服务

目录树生成:
```bash
treer -e tree.txt -i "/.idea|.git|data/"
```

## Todo...

在idl中添加optional以优化response(done)

es管理日志

将会改进comment缓存的逻辑

将会进行重构rpc以改进混沌的handler层(done)

将会更加贴合接口文档需求(doing)

将会添加双token(done)

gormopentracing,Snowflake