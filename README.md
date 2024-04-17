# MyGOWork~~4~~ 5!

**bibi-demo** is a small video website backend using hertz(hz-gen、jwt、websocket)+gorm(mysql)+redis+oss(aliyun)

## deploy by host(net=host)

(由于非常的不会shell，所以build-all很没道理，暂时不会用docker启动rpc)
`快速启动`
```bash
#oss与email的配置需自行填写
make init
make env-up
make build-all
```

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了(**Header:Authorization格式**:Bearer {token})

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

~~在idl中添加optional以优化response(done)~~

~~es管理日志(done)~~

将会改进comment缓存的逻辑(todo)

~~将会进行重构rpc以改进混沌的handler层(done)~~

将会更加贴合接口文档需求(doing)

~~将会添加双token(done)~~

gormopentracing,Snowflake(todo)