package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"

	pb "grpcgittest/proto"
)

//GLOBALS:
var usersLock = &sync.Mutex{}
var usersMap = make(map[string]chan pb.Message, 100)
var usersID = make(map[string]string)
var groupLock = &sync.Mutex{}
var groups []group

type group struct {
	name          string
	channels      []chan pb.Message
	messageMemory string
}

// UsersDB has username and password
var UsersDB = make(map[string]string)

type chatServer struct {
}

func newGroup(gname string) {
	g := group{name: gname, messageMemory: ""}
	groups = append(groups, g)
}

//Add channel to groupID, handle GroupID Adder, return Valid GroupID
func addToGroup(gid int32, c chan pb.Message, gname string) int32 {

	//fmt.Println("THIS NUMBER READ : ")
	//fmt.Println(gid)
	if num, ok := groupExists(gid); ok == false {
		newGroup(gname)
		fmt.Println("Created new group :" + gname)
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
	newGroup("default")
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

func hasUser(name string) bool {
	groupLock.Lock()
	defer groupLock.Unlock()
	_, exists := UsersDB[name]
	return exists
}

//BROADCASTS messages to respective mailboxes belonging to the same group
func broadcast(sender string, msg pb.Message) {
	usersLock.Lock()
	defer usersLock.Unlock()

	//Update Group Chat Stored
	groups[msg.GetGroup()].messageMemory = groups[msg.GetGroup()].messageMemory + sender + ">" + msg.GetText()
	for _, q := range groups[msg.GetGroup()].channels {
		if q != usersMap[sender] {
			q <- msg
		}
	}
}

//When messages come from Client stream
func listenToClient(stream pb.Chat_TransferMessageServer, messages chan<- pb.Message) {
	defer close(messages)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			//
		}
		if err != nil {
			//
			break
		}
		//Handling logout error
		if cmp := strings.Compare(msg.GetText(), ""); cmp <= 0 {
			continue
		}
		messages <- *msg
	}
}

// newLoginID
func newLoginID(login *pb.Login) string {
	s := login.GetUsername() + login.GetPassword()
	usersID[login.GetUsername()] = s
	return s
}

//Send Login Credentials
func (s *chatServer) LoginCred(ctx context.Context, login *pb.Login) (*pb.LoginResponse, error) {

	//Check if user present
	if pass, ok := UsersDB[login.GetUsername()]; ok {
		if pass == login.GetPassword() {

			//Has connection?
			if hasUser(login.GetUsername()) {

				//Online
				if hasListener(login.GetUsername()) {
					fmt.Println(login.GetUsername() + " : Mode 3 : Already Online")
					return &pb.LoginResponse{Mode: 3}, nil
				}

				//No listener but user exists, allow
				fmt.Println(login.GetUsername() + " : Mode 5 : Relogged in")
				return &pb.LoginResponse{Mode: 5, ID: newLoginID(login)}, nil
			}
			//No previous connection, accept
			fmt.Println(login.GetUsername() + " : Mode 4 : Accepted")
			return &pb.LoginResponse{Mode: 4}, nil
		}

		//Denied Login
		fmt.Println(login.GetUsername() + " : Mode 2 : Denied Login")
		return &pb.LoginResponse{Mode: 2}, nil
	}

	//Create new user credentials
	UsersDB[login.GetUsername()] = login.GetPassword()
	fmt.Println(login.GetUsername() + " : Mode 4 : Accepted")
	return &pb.LoginResponse{Mode: 4, ID: newLoginID(login)}, nil

}

//Transfer message over stream address
func (s *chatServer) TransferMessage(stream pb.Chat_TransferMessageServer) error {
	//Recieve from stream
	InMessage, err := stream.Recv()
	//Variable for client name
	clientName := InMessage.GetSender()

	//Make Client Buffer Mailbox
	clientMailbox := make(chan pb.Message)
	if err != nil {
		return err
	}

	//Assign mailbox to userZZZz (Setup Mailbox)
	addListener(clientName, clientMailbox)
	//Create a listener to client
	clientMessages := make(chan pb.Message)

	//GetGroupNUmber
	groupID := sendGroup(stream)

	//Add a group, update groupID if out of bounds
	groupID = addToGroup(groupID, clientMailbox, (clientName + "'s Group"))

	//Group Confirmation
	stream.Send(&pb.Message{
		Sender: "[SERVER]", Text: "You have joined " + groups[groupID].name, Group: groupID,
	})

	//Send Group's Previous chat
	stream.Send(&pb.Message{Sender: "[SERVER]", Text: groups[groupID].messageMemory, Group: groupID})

	//Starts chat
	go listenToClient(stream, clientMessages)

	for {
		select {
		case messageFromClient := <-clientMessages:
			if strings.Contains(messageFromClient.GetText(), "quit") {
				stream.Send(&pb.Message{Sender: "server", Text: "quit"})

				//Check if valid login ID
			} else if messageFromClient.GetLoginID() == usersID[messageFromClient.GetSender()] {
				broadcast(clientName, messageFromClient)
			}
		case messageFromOthers := <-clientMailbox:
			stream.Send(&messageFromOthers)
		}
	}
}

//Logout
func (s *chatServer) LogoutCred(ctx context.Context, logout *pb.Logout) (*empty.Empty, error) {
	removeListener(logout.GetUsername())
	fmt.Println(logout.GetUsername() + " has logged out.")
	return &empty.Empty{}, nil
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
