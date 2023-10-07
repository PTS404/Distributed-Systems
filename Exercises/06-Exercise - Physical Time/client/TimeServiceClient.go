package main

import (
	"PhysicalTime/proto"
	"bufio"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
)

type TimeServiceClient struct {
	id         int
	portNumber int
}

var (
	TimeServiceClientPort = flag.Int("tscPort", 0, "client port number")
	FirstServerPort       = flag.Int("s1Port", 0, "server port number (should match the port used for the server)")
	SecondServerPort      = flag.Int("s2Port", 0, "server port number (should match the port used for the server)")
)

func main() {
	// Parse the flags to get the port for the TimeServiceClient
	flag.Parse()

	// Create a TimeServiceClient
	TimeServiceClient := &TimeServiceClient{
		id:         1,
		portNumber: *TimeServiceClientPort,
	}

	// Wait for the TimeServiceClient (user) to ask for the time
	go waitForTimeRequestTSC(TimeServiceClient, *FirstServerPort)
	go waitForTimeRequestTSC(TimeServiceClient, *SecondServerPort)

	// Keep the TimeServiceClient running
	for {
	}
}

func waitForTimeRequestTSC(tsc *TimeServiceClient, serverPort int) {
	// Connect to the server
	serverConnection, _ := TimeServiceConnectToServer(serverPort)

	// Wait for input in the TimeServiceClient terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("Client asked for time with input: %s\n", input)

		// Ask the server for the time
		timeReturnMessage, err := serverConnection.AskForTime(context.Background(), &proto.AskForTimeMessage{
			ClientId: int64(tsc.id),
		})

		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Printf("Server %s says the time is %s\n", timeReturnMessage.ServerName, timeReturnMessage.Time)
		}
	}
}

func TimeServiceConnectToServer(serverPort int) (proto.TimeAskClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", serverPort)
	}
	return proto.NewTimeAskClient(conn), nil
}
