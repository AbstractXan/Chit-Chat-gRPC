package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"net"

	pb "chat/chat"
)

type server struct{
	UserDB map[string]string //Username and their pasword
	//UserChan map[string]chan string
	GroupDB []group
}

type group struct{
	Name string
	//Channels slice with text ('string') type
	// !!!Channels should be bufferd or else they'd deadlock !!!
	Channels []chan string 
}

var port = ":8080"

func (s *server) ConnRequest(ctx context.Context, in *pb.Request) (*pb.Confirm, error) {
	n, d := in.Username, in.Password
	log.Printf("%s : %s connected!",n,d)
	//Temporary
	return &pb.Confirm{Ok: true}, nil
}

//Broadcast function
func (s *server) Broadcast(stream pb.Chat_BroadcastServer) error{
	return nil
}
func main() {
	for{
	lis, _ := net.Listen("tcp", port)
	//MIGHT NOT NEED GO ROUTINE
		s := grpc.NewServer()
		pb.RegisterChatServer(s, &server{UserDB: make(map[string]string)})
		s.Serve(lis)
	}
}
