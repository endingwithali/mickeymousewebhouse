package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

//terminal runs -> prints booting. then starts listening for tcp type connections in 1928
// if listen fails, throw error
//if not infinite loop waiting fo rconnectiopn
func main() {
	fmt.Println("Goofy: hyuck - booting up!")

	port, err := net.Listen("tcp", ":1928")

	if err != nil { //this is on server failing to set up
		fmt.Println("Mickey: Oh no Goofy! It looks like there was an error starting up the server! ")
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := port.Accept() //this function stalls / blocks until a new incoming request is seen
		//loop fo rhandling new connections to tap into server
		//everytime we get a new client that comes in we accept it and spinn off a new handleconnection  to deal with that client
		if err != nil { //this is on sender failing to connect (other webserver)/something failed with the connection between server and client
			//trying to set up link but it flopped
			// localhost:1928 on my browser
			// as long as its a valid tcp connection
			fmt.Println(err.Error())
			return
		}
		// spins off new thread and new instance to handle all communications to thread
		go handleConnection(conn) //multithreaded asynchronisity vs single threaded async
		// dont care what thread its on, just make it happen -> go will maintain the pool of threads for the work, it wont acquire a new thread each process but inst4ead dynamically allocate thread availability

		fmt.Println("Welcome to the Mickey Mouse Web House!!!1!1!!")
		// prints twice because when you go to any website it will first request the favicon which is an incoiming request and then the website

		// how to do this without go threads later???
		// remove go -> requrests that block will block everything and you wont be able to serve as many incoming requests at the same time
		// benefit of go routine -> make them happen by just putting go infront of it
		// handle connecgtion -> must be self contained to make threading work
	}

}

// send something back to client - what happens when you request a webpage
func handleConnection(connection net.Conn) {
	defer connection.Close()
	//put at the start because ensures that you call the close at the end of the fxn no matter where the function ends
	request, err := bufio.NewReader(connection).ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println(request)

	requestParts := strings.Split(request, " ")
	fmt.Println(requestParts[1]) //by splitting we can see that it is requesting the page /

	// we need to  close th econnectionn, and it needs to be done at the end after everything else. tell go you want a piece of code to run after eveyrthing else in the fuinction is run

	// connection.Write([]byte("HTTP/1.1 200 OK\r\n"))
	// connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
	// connection.Write([]byte("Content-Length: " + strconv.Itoa(len(message)) + "\r\n\r\n"))
	// connection.Write([]byte(message + "\n"))

	if requestParts[1] == "/clubhouse" {
		message := "if goofy has a dog, and goofy is a dog....????"

		connection.Write([]byte("HTTP/1.1 200 OK\r\n"))
		connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
		connection.Write([]byte("Content-Length: " + strconv.Itoa(len(message)) + "\r\n\r\n"))
		connection.Write([]byte(message + "\n"))
		return
	}

	connection.Write([]byte("HTTP/1.1 404 Not Found\r\n"))
	connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
	connection.Write([]byte("Content-Length: 0\r\n\r\n"))

	//what we need to send
	// HTTP/1.1 200 OK\r\n
	// Content-Type: text/plain; charset=UTF-8\r\n
	// Content-Length: <length>\r\n\r\n
	// Hello World!\n

	// HTTP/1.1 200 OK\r\n
	//tells us what version of http we are using, the status code, and the text associate4d with the status code
	//each one of these lines end with \r\n to help with separating new lines

	// Content-Type: text/plain; charset=UTF-8\r\n
	//tells server the incoming type of content so we know how to handle it
	// lets later handle different content types!!!

	// Content-Length: <length>\r\n\r\n
	//ends in 4 because the section after is the start of the body content
	// universal servers expect 4 sparation characters between headers and content
	// every new header ends with 2

	// Hello World!\n

}
