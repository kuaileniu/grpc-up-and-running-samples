## proto 转java源代码
product_info.proto 文件放在 server/src/main/proto 

```
gradle build 
```
server/build/generated/source/proto/main/grpc/ecommerce/ProductInfoGrpc.java
server/build/generated/source/proto/main/java/ecommerce/ProductInfoQuterClass.java
## build server.jar
cd ch02/productinfo/java/server
 
the following shell command:
```
gradle build
```

## run server.jar
cd ch02/productinfo/java/server
 
the following shell command:
```
java -jar build/libs/server.jar
```

CNHQ19M-3zulvdrdeMacBook-Pro:server jtfeng$ java -jar build/libs/server.jar
2月 05, 2022 9:35:49 上午 ecommerce.ProductInfoServer start
信息: Server started, listening on 50051