## proto 转java源代码
product_info.proto 文件放在 client/src/main/proto 

``` Gradle 6.5
gradle build 
```
client/build/generated/source/proto/main/grpc/ecommerce/ProductInfoGrpc.java
client/build/generated/source/proto/main/java/ecommerce/ProductInfoQuterClass.java
## build client.jar
cd ch02/productinfo/java/client
 
the following shell command:
```
gradle build
```