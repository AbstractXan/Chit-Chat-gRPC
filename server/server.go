package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"

	"google.golang.org/grpc"

	pb "grpcgittest/proto"
)

//GLOBALS:
var usersLock = &sync.Mutex{}
var usersMap = make(map[string]chan pb.Message, 100)

var groupLock = &sync.Mutex{}
var groups []group

type group struct {
	name     string
	channels []chan pb.Message
}

type chatServer struct {
}

func newGroup(gname string) {
	g := group{name: gname}
	groups = append(groups, g)
}

//Add channel to groupID, handle GroupID Adder, return Valid GroupID
func addToGroup(gid int32, c chan pb.Message, gname string) int32 {

	fmt.Println("THIS NUMBER READ : ")
	fmt.Println(gid)
	if num, ok := groupExists(gid); ok == false {
		newGroup(gname)
		fmt.Println("Created new group")
		gid = num
	}
	groups[gid].channels = append(groups[gid].channels, c)
	return gid
}

func groupExists(gid int32) (int32, bool) {
	if gid > int32(len(groups)-1) {
		return int32(len(groups)), false
	}
	exists := groups[gid]
	if exists.name != "" {
		return gid, true
	}
	return gid, false

}

//Sends group list to client
func sendGroup(stream pb.Chat_TransferMessageServer) int32 {
	s := "Groups available: "

	//Send available groups
	for i, v := range groups {
		s = s + "\n" + strconv.Itoa(i) + ":" + v.name
	}

	//Send
	stream.Send(&pb.Message{
		Sender: "[SERVER]",
		Text:   s + "\n",
	})

	//Recieve
	mess, err := stream.Recv()
	if err != nil {
		return 0
	}
	return mess.Group
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

	for _, q := range groups[msg.GetGroup()].channels {
		if q != usersMap[sender] {
			q <- msg
		}
	}
}

//When messages come from Client stream
func listenToClient(stream pb.Chat_TransferMessageServer, messages chan<- pb.Message) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			//
		}
		if err != nil {
			fmt.Println(err)
			return
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

		//Set Client name as in MESSAGE
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

	//GetGroupNUmber
	groupID := sendGroup(stream)

	//Add a group, update groupID if out of bounds
	groupID = addToGroup(groupID, clientMailbox, (clientName + "'s Group"))

	//Group Confirmation
	stream.Send(&pb.Message{
		Sender: "[SERVER]", Text: "You have joined " + groups[groupID].name, Group: groupID,
	})

	//Starts chat
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
	newGroup("default")
	Serve()
}
