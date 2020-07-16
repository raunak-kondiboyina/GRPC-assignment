runserver:
	brew services start redis
	go run base/server/server.go

runclient:
	go run base/client/client.go

startRDB:
	brew services start redis

printDBlist:
	redis-cli lrange client 0 `redis-cli llen client`

installprotoc:
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go

installgrpc:
	go get -u google.golang.org/grpc

installredis:
	brew install redis

installredigo:
	go get github.com/gomodule/redigo 

UT_isvalid:
	go test base/server/*.go -v
	
	