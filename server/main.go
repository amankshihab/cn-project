package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting TCP server on localhost:8080")

	l, err := net.Listen("tcp", "192.168.1.53:8080")
	if err != nil {
		fmt.Println("Error Listening :", err.Error())

		os.Exit(1)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Connecting:", err.Error())
			return
		}

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

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
		fmt.Println(operation)

		// switch operation {

		// 	case "cr":
		// 		file , err := os.Create(bufferLower[7:])
		// 		if err != nil {
		// 			c.Write([]byte("Error while creating file, try again later."))
		// 		}
		// 		file.Close()
		// 		c.Write([]byte("File created"))

		// 	default: 
		// 		c.Write([]byte("Command not recognized"))
		// }

		c.Write([]byte("Did it"))
	}
}
