package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "grpcgittest/proto"
)

func listenToClient(sendQ chan pb.Message, reader *bufio.Reader, name string, groupnum int32) {
	for {
		msg, _ := reader.ReadString('\n')
		sendQ <- pb.Message{Sender: name, Text: msg, Group: groupnum}
	}
}

func receiveMessages(stream pb.Chat_TransferMessageClient, mailbox chan pb.Message) {
	for {
		msg, _ := stream.Recv()
		mailbox <- *msg
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

	//Recieve stream
	stream, err := client.TransferMessage(context.Background())
	if err != nil {
		return err
	}

	//Input name to construct Message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	clientName, err := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)
	if err != nil {
		return err
	}

	//Send initial message with Sender and Register=TRUE
	stream.Send(&pb.Message{Sender: clientName, Register: true})

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
	fmt.Printf("%s> %s | Group %d", mess.Sender, mess.Text, mess.Group)
	groupid = mess.Group //Updated GroupID

	//Make buffered mailbox recieve message from server
	mailBox := make(chan pb.Message, 100)
	go receiveMessages(stream, mailBox)

	//Make send queue buffered message
	sendQ := make(chan pb.Message, 100)
	go listenToClient(sendQ, reader, clientName, groupid)

	//Forever
	for {
		select {

		//If send channel is active, send to server
		case toSend := <-sendQ:
			stream.Send(&toSend)

		//If mailbox has something, print.
		case received := <-mailBox:
			fmt.Printf("Group %d | %s  > %s", received.Group, received.Sender, received.Text)
		}
	}
}

func main() {
	err := Connect("127.0.0.1:10000")
	if err != nil {
		log.Println(err)
	}
}
