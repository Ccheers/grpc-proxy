# gRPC Proxy

fork from https://github.com/mwitkow/grpc-proxy 

## GRPC 代理 无需解码编码，代理 GRPC 

## 使用说明

```go 


var unknownServiceHandler =  func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
    ....
}

// 必须注册这个方法才能运行
grpc.NewServer(GrpcProxyOptions(unknownServiceHandler)...)

```
