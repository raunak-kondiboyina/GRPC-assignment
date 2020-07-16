package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/GRPC-assignment/base/proto"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

type server struct {
}

func Isvalid(name string) bool {
	for i := 0; i < len(name)-1; i++ {
		//fmt.Println(string(c))
		if (name[i] >= 97 && name[i] <= 122) || (name[i] >= 65 && name[i] <= 90) || (name[i] >= 48 && name[i] <= 57) {
			//fmt.Println(name[i])
		} else {
			return false
		}
	}
	return true
}

func (*server) Send(ctx context.Context, request *proto.SendRequest) (*proto.SendResponse, error) {
	name := request.Name
	fmt.Println(len(name))
	match := Isvalid(name)

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).

	if match == true {
		_, err = conn.Do("lpush", "client", name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("added to DB")

		response := &proto.SendResponse{

			Inserted: true,
		}
		return response, nil
	} else {
		response := &proto.SendResponse{
			Inserted: false,
			//"invalid input, enter A-Z,a-z,0-9",
		}
		return response, nil
	}

}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	proto.RegisterSendServiceServer(s, &server{})

	s.Serve(lis)
}
