## ``ProductInfo`` Service and Client - Java Implementation
## 由proto 生成java源代码
- 
(productinfo/java/server) and execute the following
shell command,

``` Gradle 6.5
gradle build

生成 
server/build/generated/source/proto/main/grpc/ecommerce/ProductInfoGrpc.java
和
server/build/generated/source/proto/main/java/ecommerce/ProductInfoQuterClass.java
```
### Building and Running Service

In order to build gradle project, Go to ``Java`` project root directory location (samples) and execute
 the following shell command,
```
./gradlew :ch02:productinfo:java:server:build
```

In order to run, Go to ``Java`` project root directory location (productinfo/java/server) and execute the following
shell command,

```
java -jar build/libs/server.jar
```

### Building and Running Client

In order to build gradle project, Go to ``Java`` project root directory location (inside samples directory) and execute
 the following shell command,
```
./gradlew :ch02:productinfo:java:client:build
```

In order to run, Go to ``Java`` project root directory location (productinfo/java/client) and execute the following
shell command,

```
java -jar build/libs/client.jar
```
