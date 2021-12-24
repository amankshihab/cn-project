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
	fmt.Println("Listening on port 8080..")

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

		switch operation {

		// Creates a new file
		case "cr":
			file, err := os.Create(strings.Trim(bufferLower[7:], " \n"))
			if err != nil {
				c.Write([]byte("Error while creating file, try again later."))
			}
			file.Close()
			c.Write([]byte("File created"))

		// same effect as executing cat command on a file
		case "ca":
			content, err := os.ReadFile(strings.Trim(bufferLower[4:], " \n"))
			if err != nil {
				c.Write([]byte("Won't work mate"))
			}
			if len(content) != 0 {
				c.Write(content)
			} else {
				c.Write([]byte(" "))
			}
		
		// deletes a file
		case "de":
			err := os.Remove(strings.Trim(bufferLower[7:], " \n"))
			if err != nil {
				c.Write([]byte("Error deleting the files"))
			}
			c.Write([]byte("File Deleted"))

		default:
			c.Write([]byte("Command not recognized"))
		}
	}
}
