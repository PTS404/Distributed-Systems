package main

import (
	"fmt"
	"net" //https://pkg.go.dev/net#pkg-overview
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Start the server in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		server()
	}()

	// Start the client in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		client()
	}()

	//Wait for the program to finalize
	wg.Wait()
}

func server() {
	//Creates the server
	server, err := net.Listen("tcp", "localhost:8080") //Listens for connections on localhost and port 8080
	if err != nil {
		panic(err) //Throws an exception - Prints message and function then crashes - https://www.educative.io/answers/what-is-panic-in-golang
	}
	defer server.Close() //Closes the connection once the program terminates/finishes

	//Split the server address into its host and port components
	host, port, err := net.SplitHostPort(server.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	//Waits for the next connection and returns it to the listener (server)
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go handleSequence(conn)
	}
}

func handleSequence(conn net.Conn) {
	//Makes a buffer of 1 kb in size
	buffer := make([]byte, 1024)

	//Simulate first step in 3 way handshake (Receive seq from client)
	msg, err := conn.Read(buffer) //Reads the connection and saves the message received
	if err != nil {
		fmt.Printf("Error trying to read: %#v\n", err)
		return
	}

	seq, _ := strconv.Atoi(string(buffer[:msg]))   //Turns the received string to an int
	fmt.Printf("Server received: seq = %d\n", seq) //Prints out the received seq

	//Simulate second step in 3 way handshake (Send ack and seq to client)
	ack := seq + 1                           //Updates ack with previous seq value + 1
	seq = 42                                 //Updates seq to a new value
	ackseq := fmt.Sprintf("%d %d", ack, seq) //Creates a string that contains both ack and seq
	_, err = conn.Write([]byte(ackseq))      //Writes back to the client syn ack and seq (ack = x+1 | seq = y)

	if err != nil { //Error handling
		fmt.Println("Error sending the data:", err)
		return
	}

	//Simulate third step in 3 way handshake (Recieve ack and seq)
	msg, err = conn.Read(buffer) //Reads the connection and saves the message received

	if err != nil { //Error handling
		fmt.Printf("Error trying to read: %#v\n", err)
		return
	}

	newAckSeq := strings.Split(string(buffer[:msg]), " ") //Splits the string that contains both ack and seq
	newAck, _ := strconv.Atoi(newAckSeq[0])
	newSeq, _ := strconv.Atoi(newAckSeq[1])

	fmt.Printf("Server received: ack = %d | seq = %d\n", newAck, newSeq)

	//Send confirmation of last receival
	conn.Write([]byte("Received"))

	//Receive actual data
	msg, err = conn.Read(buffer) //Reads the connection and saves the message received

	if err != nil { //Error handling
		fmt.Printf("Error trying to read: %#v\n", err)
		return
	}

	fmt.Println("Server received (data):", string(buffer[:msg])) //Outputs data

	//Closes the connection
	conn.Close()
}

func client() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil { //Error handling
		fmt.Println("Error trying to connect to the server:", err)
		return
	}

	defer conn.Close() //Closes the connection when the program finishes

	//Creates a buffer of 1kb in size
	buffer := make([]byte, 1024)

	//Simulate first step in 3 way handshake (Send seq to server)
	seq := "1"
	_, err = conn.Write([]byte(seq))

	if err != nil { //Error handling
		fmt.Println("Error sending the data:", err)
		return
	}

	//Simulate second step in 3 way handshake (Receive ack and seq)
	msg, err := conn.Read(buffer)

	if err != nil { //Error handling
		fmt.Println("Error receiving the data:", err)
		return
	}

	ackseq := strings.Split(string(buffer[:msg]), " ")
	newAck, _ := strconv.Atoi(ackseq[0])
	newSeq, _ := strconv.Atoi(ackseq[1])

	fmt.Printf("Client received: ack = %d | seq = %d \n", newAck, newSeq)

	//Simulate third step in 3 way handshake (Send ack and seq)
	twoAck := newSeq + 1
	twoSeq := newAck + 1
	newAckSeq := fmt.Sprintf("%d %d", twoAck, twoSeq)
	conn.Write([]byte(newAckSeq)) //Writes back to the server syn ack and seq (ack = y+1 | seq = x+1)

	//Receive last confirmation from server
	msg, err = conn.Read(buffer)

	if err != nil { //Error handling
		fmt.Println("Error receiving the data:", err)
		return
	}

	confirmation := string(buffer[:msg])

	//Send data to the server if confirmation is true (connection is established through the 3 way handshake)
	if confirmation == "Received" {
		message := "Hello, Server!" //This is the final message being sent to the server
		_, err = conn.Write([]byte(message))

		if err != nil { //Error handling
			fmt.Println("Error sending the data:", err)
			return
		}
	}
}
