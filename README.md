# Go Event Source

## Problem Statement
Simple project of how to impelemet grpc protobuf in golang as client and server. For further documentation you can click [here](https://grpc.io/docs/tutorials/basic/go/)

## How it works
The fundamental idea of grpc is automated mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages. You can even update your data structure without breaking deployed programs that are compiled against the "old" format.

In gRPC a client application can directly call methods on a server application on a different machine as if it was a local object, making it easier for you to create distributed applications and services. As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.

![Github Logo](https://grpc.io/img/landing-2.svg)

gRPC clients and servers can run and talk to each other in a variety of environments - from servers inside Google to your own desktop - and can be written in any of gRPC’s supported languages. So, for example, you can easily create a gRPC server in Java with clients in Go, Python, or Ruby. In addition, the latest Google APIs will have gRPC versions of their interfaces, letting you easily build Google functionality into your applications.

In this case, I am making client and server in golang. The context is just to greet some text to the client.

## Dependencies
In this time, I use mac as an os to running the system.

#### 1. Golang
First of all export some paths, and save them in your `.zshrc` or `.bashrc` files for easy use. Use sudo if you get error.

```
# Go development
```

```
export GOPATH="${HOME}/.go"
```

```
export GOROOT="$(brew --prefix golang)/libexec"
```

```
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

```
test -d "${GOPATH}" || mkdir "${GOPATH}"
```

And then, you can install the Go
```
brew install go

```

Now the main part you must following command below to install grpc
```
go get -u google.golang.org/grpc
```

After that, you have to install the latest protocol buffer. Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(protoc-`version`-`platform`.zip) from here: https://github.com/google/protobuf/releases

- Unzip this file.
- Update the environment variable PATH to include the path to the protoc binary file.

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.
```
export PATH=$PATH:$GOPATH/bin
```
If you're confusing with the explained above. You can follow someone's instruction on [here](https://medium.com/@erika_dike/installing-the-protobuf-compiler-on-a-mac-a0d397af46b8) which is very simple in my perspective.

## How to use
First, you have to generate the proto file to go dependencion with this command:
```
protoc -I api/proto/v1 api/proto/v1/todo-service.proto --go_out=plugins=grpc:api/proto/v1
```

After that, you can run the server and client main in different terminal.

### Server Side :
```
go run greeter_server/main.go
```
### Client Side :
```
go run greeter_client/main.go
```
