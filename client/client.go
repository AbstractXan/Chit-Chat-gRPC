package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "grpcgittest/proto"
)

func listenToClient(sendQ chan pb.Message, reader *bufio.Reader, name string, groupnum int32, id string) {
	defer close(sendQ)
	for {
		msg, _ := reader.ReadString('\n')
		sendQ <- pb.Message{Sender: name, Text: msg, Group: groupnum, LoginID: id}

		if strings.Contains(msg, "quit") {
			fmt.Println("Closing Client Sender.")
		}

	}
}

func receiveMessages(stream pb.Chat_TransferMessageClient, mailbox chan pb.Message) {
	defer close(mailbox)
	for {
		msg, _ := stream.Recv()
		mailbox <- *msg

		if msg.GetSender() == "server" && strings.Contains(msg.GetText(), "quit") {
			fmt.Println("Closing Client Reciever.")
			return
		}

	}
}

// Connect : Connects to server
func Connect(address string) error {

	//Dial to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	//Setup new client
	client := pb.NewChatClient(conn)

	//Input name to construct Message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	clientName, err := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)
	if err != nil {
		return err
	}
	fmt.Print("Enter your password: ")
	clientPass, err := reader.ReadString('\n')
	clientPass = strings.TrimSpace(clientName)
	if err != nil {
		return err
	}

	//Get Login Thingy
	LoginStats, err := client.LoginCred(context.Background(), &pb.Login{Username: clientName, Password: clientPass})
	if err != nil {
		return err
	}

	switch LoginStats.Mode {
	case 2:
		err = errors.New("Login Denied. Wrong Credential")
		return err
	case 3:
		err = errors.New("User already online")
		return err
	case 4:
		fmt.Println("Login successful")
	case 5:
		fmt.Println("Welcome.")
	default:
		err = errors.New("??? ERROR reading loginstats")
		return err
	}

	//Define stream
	stream, err := client.TransferMessage(context.Background())
	if err != nil {
		return err
	}

	//GET PrivateID
	id := LoginStats.ID

	//Send username. Client name
	stream.Send(&pb.Message{Sender: clientName})

	//Getgroupmessage
	groupmessage, err := stream.Recv()
	if err != nil {
		return err
	}
	fmt.Printf("%s> %s", groupmessage.Sender, groupmessage.Text)

	//Send group num as Group int32
	groupnum, err := reader.ReadString('\n')
	groupnum = strings.TrimSpace(groupnum)
	gnum, _ := strconv.Atoi(groupnum)
	groupid := int32(gnum)

	//Send groupnum and Server registers member to groupnum
	stream.Send(&pb.Message{Sender: clientName, Group: groupid})

	//Recieve confirmation to group
	mess, err := stream.Recv()
	if err != nil {
		return err
	}

	//Print Group Confirmation
	fmt.Printf("%s> %s | Group %d\n", mess.Sender, mess.Text, mess.Group)
	groupid = mess.Group //Updated GroupID

	//Get Prev Group Chat
	gmess, err := stream.Recv()
	if err != nil {
		return err
	}

	//Print Prev Group Chat
	fmt.Println(gmess.Text)

	//Make buffered mailbox recieve message from server
	mailBox := make(chan pb.Message)
	go receiveMessages(stream, mailBox)

	//Make send queue buffered message
	sendQ := make(chan pb.Message)
	go listenToClient(sendQ, reader, clientName, groupid, id)

	//Forever
	for {
		select {

		//If send channel is active, send to server
		case toSend := <-sendQ:
			stream.Send(&toSend)

		//If mailbox has something, print.
		case received := <-mailBox:
			if received.GetSender() == "server" && strings.Contains(received.GetText(), "quit") {
				fmt.Println("Exiting server.")
				client.LogoutCred(context.Background(), &pb.Logout{Username: clientName})
				return nil
			}
			fmt.Printf("%s> %s", received.Sender, received.Text)
		}
	}
}

func main() {
	err := Connect("127.0.0.1:10000")
	if err != nil {
		log.Println(err)
	}
}
