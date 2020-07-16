package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GRPC-assignment/base/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...ENTER NAME")
	reader := bufio.NewReader(os.Stdin) //added
	name, _ := reader.ReadString('\n')  //added

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := proto.NewSendServiceClient(cc)
	request := &proto.SendRequest{Name: name}

	resp, _ := client.Send(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Inserted)
}
