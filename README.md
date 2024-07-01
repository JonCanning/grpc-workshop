Install the protobuf compiler  
https://grpc.io/docs/protoc-installation/

These languages are supported natively
- C++
- C#
- Java
- Kotlin
- Objective-C
- PHP
- Python
- Ruby
- Rust

For the languages we're interested in we'll install plugins

For GO
```
go install github.com/golang/protobuf/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

For TS
```
npm add ts-protoc-gen
```

See the respective `run.sh` files for how to compile the `.proto` for each language  

You'll need to install Go and [Bun](https://bun.sh/) for TS

From a `.proto` file we can generate protobuf type definitions which we can use for message brokers and other asynchronous activities, and we can easily generate GRPC clients and servers with support for security and backoff/retry.

`./run.sh` will give you a choice in each language to run the client or server  

Have a look at [Buf](https://buf.build/docs/introduction) as an alternative to protoc  

[gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) is an interesting project that extends protoc to create a JSON API too
