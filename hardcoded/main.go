package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// terminal runs -> prints booting. then starts listening for tcp type connections in 1928
// if listen fails, throw error
// if not infinite loop waiting fo rconnectiopn

func main() {
	fmt.Println("Goofy: hyuck - booting up!")

	port, err := net.Listen("tcp", ":1928")

	if err != nil { // on server failing to set up
		fmt.Println("Mickey: Oh no Goofy! It looks like there was an error starting up the server! ")
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := port.Accept() // this function stalls / blocks until a new incoming request is seen
		// everytime we get a new client that comes in we accept it and spin off a new handleconnection to deal with that client

		if err != nil { // on sender failing to connect (other webserver)/something failed with the connection between server and client
			// connect either via localhost:1928 on my browser or any valid tcp connection
			fmt.Println(err.Error())
			return
		}

		// spins off new thread and new instance to handle all communications to thread
		go handleConnection(conn)

		fmt.Println("Welcome to the Mickey Mouse Web House!")
		// prints twice because when you go to any website it will first request the favicon which is an incoiming request and then the website

		// if you remove go -> requests that block will block everything and you wont be able to serve as many incoming requests at the same time
		// benefit of go routine -> make them happen by just putting go infront of it
		// handle connection -> must be self contained to make threading work
	}

}

// send something back to client - what happens when you request a webpage
func handleConnection(connection net.Conn) {
	defer connection.Close()
	// put at the start because ensures that you call the close at the end of the fxn no matter where the function ends
	// we need to  close the connectionn, and it needs to be done at the end after everything else - tells go you want a piece of code to run after eveyrthing else in the fuinction is run

	request, err := bufio.NewReader(connection).ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	requestParts := strings.Split(request, " ") //by splitting we can see that it is requesting the page /clubhouse

	if requestParts[1] == "/clubhouse" {
		message := "If Goofy has a dog, and Goofy is a dog....????"

		connection.Write([]byte("HTTP/1.1 200 OK\r\n"))
		connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
		connection.Write([]byte("Content-Length: " + strconv.Itoa(len(message)) + "\r\n\r\n"))
		connection.Write([]byte(message + "\n"))
		return
	}

	connection.Write([]byte("HTTP/1.1 404 Not Found\r\n"))
	connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
	connection.Write([]byte("Content-Length: 0\r\n\r\n"))

	// what we need to send:
	// HTTP/1.1 200 OK\r\n
	// Content-Type: text/plain; charset=UTF-8\r\n
	// Content-Length: <length>\r\n\r\n
	// Hello World!\n

	// HTTP/1.1 200 OK\r\n
	// tells us what version of http we are using, the status code, and the text associated with the status code
	// each one of these lines end with \r\n to help with separating new lines

	// Content-Type: text/plain; charset=UTF-8\r\n
	// tells server the incoming type of content so we know how to handle it
	// lets later handle different content types!!!

	// Content-Length: <length>\r\n\r\n
	// ends in 4 because the section after is the start of the body content
	// universal servers expect 4 sparation characters between headers and content
	// every new header ends with 2

	// Hello World!\n

}
