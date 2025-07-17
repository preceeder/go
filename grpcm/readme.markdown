### 执行命令
```cmd
protoc --go_out=./src --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./src --go-grpc_opt=paths=source_relative *.proto
protoc -I=. -.gothon_out=. --grpc.gothon_out=. 
```