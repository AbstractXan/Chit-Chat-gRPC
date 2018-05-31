package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "grpcgittest/proto"
)

var usersLock = &sync.Mutex{}

var usersMap = make(map[string]chan pb.Message, 100)

type chatServer struct {
}

func newChatServer() *chatServer {
	return &chatServer{}
}

func addListener(name string, msgQ chan pb.Message) {
	usersLock.Lock()
	defer usersLock.Unlock()
	usersMap[name] = msgQ
}

func removeListener(name string) {
	usersLock.Lock()
	defer usersLock.Unlock()
	delete(usersMap, name)
}

func hasListener(name string) bool {
	usersLock.Lock()
	defer usersLock.Unlock()
	_, exists := usersMap[name]
	return exists
}

func broadcast(sender string, msg pb.Message) {
	usersLock.Lock()
	defer usersLock.Unlock()
	for user, q := range usersMap {
		if user != sender {
			q <- msg
		}
	}
}

//When messages come from Client stream
func listenToClient(stream pb.Chat_TransferMessageServer, messages chan<- pb.Message) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// ?
		}
		if err != nil {
			// ??
		}
		messages <- *msg
	}
}

//Transfer message over stream address
func (s *chatServer) TransferMessage(stream pb.Chat_TransferMessageServer) error {

	//Recieve from stream
	InMessage, err := stream.Recv()

	//Variable for client name
	var clientName string

	//Make Client Buffer Mailbox
	clientMailbox := make(chan pb.Message, 100)

	if err != nil {
		return err
	}

	//If Register == TRUE
	if InMessage.Register {

		//Set Client name as in MESAAGE
		clientName = InMessage.Sender

		//Check if HasListener, i.e. has a channel
		if hasListener(clientName) {
			return fmt.Errorf("name already exists")
		}

		//Else add listener, i.e. new channel
		addListener(clientName, clientMailbox)
	} else {
		return fmt.Errorf("need to register first")
	}

	//Create a listener to client
	clientMessages := make(chan pb.Message, 100)
	go listenToClient(stream, clientMessages)

	for {
		select {
		case messageFromClient := <-clientMessages:
			broadcast(clientName, messageFromClient)
		case messageFromOthers := <-clientMailbox:
			stream.Send(&messageFromOthers)
		}
	}
}

//Serve : Serves at specific address
func Serve() error {

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServer(grpcServer, newChatServer())
	grpcServer.Serve(lis)
	return nil
}

func main() {
	Serve()
}
