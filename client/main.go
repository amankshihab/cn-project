package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Println("Conncting to TCP server @ localhost:8080...")

	conn, err := net.Dial("tcp", "192.168.1.53:8080") // the address of the server can be changed based on requirements
	if err != nil {
		fmt.Println("Error connecting to server, exiting..")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">>")

		input, _ := reader.ReadString('\n')

		// To exit from the prompt
		if input == "exit\n" {
			break
		}

		conn.Write([]byte(input))

		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Println("Server says: " + message)
	}
}
