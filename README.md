## Chit-Chat-gRPC by AbstractXan

### Setup
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc


### Run server and clients:

#### Server
```
$ cd $GOPATH/src/Chit-Chat-gRPC/
$ go run ./server/server.go
```

#### Client
```
$ cd $GOPATH/src/Chit-Chat-gRPC/
$ go run ./client/client.go
```
### Features:
- Client login feature
- Join existing group / New group creation.
- Client logout on typing quit. Closes all readers and Writers.
- Server could identify logouts and close client listeners and update client channels accordingly
- Chats stored in global groups variables. Server sends all past messages to new clients for chosen group.

#### Possible addons:
MAC address validation in case of client side force quit.
