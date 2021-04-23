package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/endingwithali/mickeymousewebhouse/dynamic/clubhouse"
)

func main() {
	webServer := WebServer{}

	webServer.Init()

	webServer.AddRoute("/clubhouse", clubhouse.ClubhouseRoute)

	webServer.Serve(1928)
}

type WebServer struct {
	routes map[string]func([]string) string
}

func (ws *WebServer) Init() {
	ws.routes = make(map[string]func([]string) string)
}

func (ws *WebServer) AddRoute(route string, fxn func([]string) string) {
	ws.routes[route] = fxn
}

func (ws *WebServer) handleConnection(connection net.Conn) {
	defer connection.Close()
	request, err := bufio.NewReader(connection).ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	requestParts := strings.Split(request, " ")

	reqFunc := ws.routes[requestParts[1]]

	if reqFunc != nil {
		result := reqFunc(requestParts)
		connection.Write([]byte(result))
		return
	}

	connection.Write([]byte("HTTP/1.1 404 Not Found\r\n"))
	connection.Write([]byte("Content-Type: text/plain; charset=UTF-8\r\n"))
	connection.Write([]byte("Content-Length: 0\r\n\r\n"))

}

func (ws *WebServer) Serve(port int) {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("Mickey: Oh no Goofy! It looks like there was an error starting up the server! ")
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go ws.handleConnection(conn)
		fmt.Println("Welcome to the Mickey Mouse Web House!!!1!1!!")
	}
}
