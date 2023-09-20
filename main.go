package main

import (
	"fmt"
	"net" //https://pkg.go.dev/net#pkg-overview
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

	// Split the server address into its host and port components
	host, port, err := net.SplitHostPort(server.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	//Waits for the next connection and returns it to the listener (var server)
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//Makes a buffer of 1 kb in size
	buffer := make([]byte, 1024)

	//Reads the connection and saves the message received
	msg, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error trying to read: %#v\n", err)
		return
	}

	//Prints out the received message
	fmt.Printf("Message received: %s\n", string(buffer[:msg]))

	//Writes back to the client that the message was received
	conn.Write([]byte("Message received.\n"))

	//Closes the connection
	conn.Close()
}

func client() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error trying to connect to the server:", err)
		return
	}
	defer conn.Close() //Closes the connection when the program finishes

	// Send data to the server
	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending the data:", err)
		return
	}

	// Receive and print the response from the server
	buffer := make([]byte, 1024)
	msg, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving the data:", err)
		return
	}
	response := string(buffer[:msg])
	fmt.Println("Server response:", response)
}
