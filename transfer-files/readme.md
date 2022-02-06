## 参考
- 一个命令行小工具，使用grpc传输文件 https://github.com/dalianzhu/transfile
- https://github.com/arisetransfer/arise
## 传输文件（图片、视频、文档等)


 

### Generate Server and Client side code 
command for the server directory(inside transfer-files/go/server)
```
go mod init transferfiles/server

go get google.golang.org/grpc@v1.44.0
go get google.golang.org/protobuf@v1.27.1
```

command for the server directory(inside transfer-files/go/client)
```
go mod init transferfiles/client
go get google.golang.org/grpc@v1.44.0
go get google.golang.org/protobuf@v1.27.1
```

 
Pre-generated stub file is included in the go project. If you need to generate the stub files please use the below
 command from the root directory(inside transfer-files directory)
``` 

protoc -I proto/ proto/transferfiles_info.proto --go_out=plugins=grpc:./go/server/    // 生成go语言，需要预装 protoc-gen-go ; go install github.com/golang/protobuf/protoc-gen-go


protoc -I proto/ proto/transferfiles_info.proto --go_out=plugins=grpc:./go/client/    // 生成go语言，需要预装 protoc-gen-go ; go install github.com/golang/protobuf/protoc-gen-go

``` 