package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Starting server at", SERVER_HOST, ":", SERVER_PORT)
	ln, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Listening on", SERVER_HOST, ":", SERVER_PORT)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			os.Exit(1)
		}
		fmt.Println("Client connected.")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	mLenght, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
	}
	fmt.Println("Received data:", string(buffer[:mLenght]))
	request := string(buffer[:mLenght])

	lines := strings.Split(request, "\n")
	requestLine := lines[0]

	parts := strings.Split(requestLine, " ")

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nRequested path: %s", parts[1])

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing:", err)
	}
	conn.Close()
}
