# 依赖安装
执行一下命令，构建项目，自动安装依赖
```go
go build
go run cmd/server/main.go
```

或者手动下载依赖
```go
go mod download
```
或者下载并且验证依赖
```go
go mod tidy
go mod verify
```

运行服务
```go
go run cmd/server/main.go   
```

接口测试

使用的是 [ApiFox](https://www.apifox.cn/)

导入接口文档 [go-blog.Apifox.json](go-blog.Apifox.json)