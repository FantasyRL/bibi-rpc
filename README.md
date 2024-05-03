# MyGOWork~~4~~ ~~5~~ 6!

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

使用：将docs/swagger.* 丢到apifox/postman，然后就能用了



## Recent:
进行rpc重构


目录树生成:
```bash
treer -e tree.txt -i "/.idea|.git|data/"
```

## Todo...

~~在idl中添加optional以优化response(done)~~

~~es管理日志(done)~~

将会改进comment缓存的逻辑(todo)

~~将会进行重构rpc以改进混沌的handler层(done)~~

~~将会更加贴合接口文档需求(doing)~~

~~将会添加双token(done)~~

gormopentracing,Snowflake(todo)

## 已知问题

