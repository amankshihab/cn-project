package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	log.Println("Starting TCP server on localhost:8080")

	l, err := net.Listen("tcp", "192.168.1.53:8080")
	if err != nil {
		log.Println("Error Listening :", err.Error())

		os.Exit(1)
	}
	defer l.Close()

	log.Println("Listening on port 8080..")

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Error Connecting:", err.Error())
			return
		}

		log.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	for {
		buffer, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Println("Client " + c.RemoteAddr().String() + " left")
			return
		}

		bufferLower := strings.ToLower(string(buffer))

		operation := bufferLower[:2]

		switch operation {

			case "cr": createFile(bufferLower, c) // creates a new file

			case "ca": catCommand(bufferLower, c) // same as "cat" in bash

			case "de": deleteFile(bufferLower, c) // deletes a file

			case "ed": editFile(bufferLower, c) // append content to a file

			default: c.Write([]byte("Command not recognized"))
		}
	}
}
