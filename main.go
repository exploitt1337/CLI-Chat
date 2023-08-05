package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

const (
	password = "exploit" 
	port = "1337"
)

var (
	clients   = make(map[net.Conn]string) 
	clientsMu sync.Mutex                  
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Enter password: "))
	passwordInput, _ := bufio.NewReader(conn).ReadString('\n')
	passwordInput = strings.TrimSpace(passwordInput)

	if passwordInput != password {
		conn.Write([]byte("Invalid password. Connection terminated.\n"))
		return
	}

	conn.Write([]byte("Username: "))
	username, _ := bufio.NewReader(conn).ReadString('\n')
	username = strings.TrimSpace(username)

	clientsMu.Lock()
	clients[conn] = username
	clientsMu.Unlock()
	fmt.Println(username + " has joined the chat.")
	broadcastMessage(fmt.Sprintf("%s has joined the chat.", username), conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			clientsMu.Lock()
			username := clients[conn]
			delete(clients, conn)
			clientsMu.Unlock()
			fmt.Println(username + " has left the chat.")
			broadcastMessage(fmt.Sprintf("%s has left the chat.", username), conn)
			return
		}
		message = strings.TrimSpace(message)

		clientsMu.Lock()
		username := clients[conn]
		clientsMu.Unlock()
		fmt.Println(username+": "+message)
		broadcastMessage(username+": "+message, conn)

	}
}

func broadcastMessage(message string, sender net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		// if conn != sender { 
			conn.Write([]byte(message + "\n"))
		// }
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:" + port)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("server running on port:", port + "\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
