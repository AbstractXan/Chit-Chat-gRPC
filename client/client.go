package main

import (
	pb "chat/chat"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func connect(client pb.ChatClient, request *pb.Request) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.ConnRequest(ctx, request)
	if err != nil {

		log.Fatalf("%v.connect(_) = _, %v", client, err)
		return false
	}
	log.Println(resp)
	return true
}
func main() {
	serverAddr := net.JoinHostPort("localhost", "8080")

	//setup insecure connection
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewChatClient(conn)
	ok := connect(client, &pb.Request{Username: "Xan", Password: "123"})
	if ok == false {
		return
	} //Exits if cannot connect

}
