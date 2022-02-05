## ``ProductInfo`` Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module root directory location (productinfo/go/server) and execute the following
 shell command,
```
go build -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (productinfo/go/server) and execute the following
shell command,

```
./bin/server
```

## Building and Running Client   

In order to build, Go to ``Go`` module root directory location (productinfo/go/client) and execute the following
 shell command,
```
go build -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (productinfo/go/client) and execute the following
shell command,

```
./bin/client
```

## Additional Information

### Generate Server and Client side code 
Pre-generated stub file is included in the go project. If you need to generate the stub files please use the below
 command from the root directory(inside productinfo directory)
``` 
protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/server/ecommerce

升级 go.mod 为:
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1

去掉末尾的ecommerce，修改为
protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:./go/server/    // 生成go语言，需要预装 protoc-gen-go ; go install github.com/golang/protobuf/protoc-gen-go

``` 
