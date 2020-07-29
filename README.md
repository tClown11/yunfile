# 项目情况

这是yunfile的微服务版本，暂时还没有将下载与上传并入

技术栈：golang、gin、redis、mysql、protobuf、rabbitmq、阿里云oss、grpc、go-micro

暂时工作进度：实现上传接口、完成登录注册功能、完成token验证、文件信息获取、文件下载、文件更新、文件删除

尚未实现功能： redis实现断点续传、反向路由代理

## go-micro应用与学习

- 本项目的微服务的服务端实现的是对于数据库的交互，客户端实现的是前端的数据交互与对后端返回的进行处理并返回给前端

- 基本上各个微服务的客户端可以统合成一个客户端，用一个客户端去调用多个服务端

go-micro是个管理微服务的框架，go-micro服务端与客户端之间的通过rpc进行数据交互

go-micro的默认注册的方式是mdns，在生成端最后使用consul或etcd来进行服务的发现与注册

## 设计(非微服务架构)

- go-micro设计

  - 将数据库操作与前端的数据判断、返回操作相隔离

- 包的设计

  - dao：对数据库的数据进行交互方法与数据结构

  - dto：返回前端展示的数据结构与方法

  - services：与前端url相响应的方法集

  - config：数据库链接的信息

  - util：一些其他功能的实现(token验证、加密等)

- 表的设计

  - 用户表：增删查

  - 用户文件表：增加(上传)、删除、修改(update)、查询(基于用户id的所有已上传文件的查询)

  - 文件表：增加(基于文件hash不同)、查询(主要找到相应文件hash的下载地址，单个文件查询)

- 文件存储的方式

  - 数据库只存文件的下载路径，文件内容保存到本地硬盘或云盘中，倘若保存到数据库中，会拖慢数据库的执行效率，更坏的可能会导致数据库崩溃

  - 文件hash值与文件下载路径成立一张数据表，使文件数据仅存一份，可以减少资源开销

  - 用户id与数据文件hash存到一张表

  - 上传文件可以先查询文件hash是否存在，再进行数据存储

- 实现文件上传到MySQL

  - 服务端获取上传文件必须在服务端开启一个新的文件写入上传的文件流(必须将`multipart.File`通过文件写入的方式将其转成`os.file`的服务端本地文件)

  - 随后通过sha1算法将文件转成一个sha1的string，随后便可调用mysql的insert方法插入到mysql数据库中

  - 最后可以删除本地文件,减少不必要的内存浪费，下次下载重数据库中获取

- 登录时设置uid与token的到cookie 

  - 在gin中实现setcookie可以使用`func (*gin.Context).SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)`

  - 设置多个cookie暂时使用的是多次调用SetCookie方法

### 关于分块上传于断点上传、秒传

- 由于前端并不太熟悉，该项目只做后端实现，通过postman验证

- 关于文件秒传问题：

  秒传需要后端判断其上传文件的hash值(MD5、sha1等)，倘若前端没有对上传文件进行hash，并将其hash传到后端，这样后端就需要接收文件在进行hash值判断，这样就无法实现秒传的意义，所以本项目在前端没有对文件进行hash的情况下，井仔后端对文件进行hash保存时进行hash判断，有重复的就删除新上传的，保证hash值文件的唯一性

- 分块上传

  - 需要在客户端进行分块的划分

## 参考

- gin 中间件执行逻辑参考：https://www.imooc.com/article/282992

- gorm的使用：
  
  - 文档：http://gorm.io/zh_CN/docs/index.html

  - github：https://github.com/go-gorm/gorm

- go-redis： https://github.com/go-redis/redis

- protobuf and grpc：

  - [protobuf官网](https://developers.google.com/protocol-buffers)

## 遇到的问题

- protoc-gen-micro执行报错

  - 来自https://github.com/micro/micro/tree/master/cmd/protoc-gen-micro 的答案

    Errors
    If you see an error about protoc-gen-micro not being found or executable, it's likely your environment may not be configured correctly. If you've already installed protoc, protoc-gen-go, and protoc-gen-micro ensure you've included $GOPATH/bin in your PATH.

    出现上述的错误可以手动编译protoc-gen-micro放到`$GOPATH/bin/`目录下

Alternative specify the Go plugin paths as arguments to the protoc command

- DevTools failed to load SourceMap: Could not load content for chrome-extension://gighmmpiobklfepjocnamgkkbiglidom/include.preload.js.map: HTTP error: status code 404, net::ERR_UNKNOWN_URL_SCHEME

  - 参考：https://juejin.im/post/5ea7c86a6fb9a0436545a11c 的解决方法

  - 本项目出现的问题是因为cdn已压缩的资源文件在引用的时候发生错误导致页面无法正常展示，需要清除缓存方可以继续访问

  - chrome关于插件报错的讨论：https://stackoverflow.com/questions/60302529/devtools-failed-to-parse-sourcemap-chrome-extension

  - 用随机数，随机数也是避免缓存的一种很不错的方法！   `URL 参数后加上 "?ran=" + Math.random(); //当然这里参数 ran可以任意取了`
