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

	fmt.Println("Conncting to TCP server @ localhost:8080...")

	conn, err := net.Dial("tcp", "192.168.1.53:8080") // the address of the server can be changed based on requirements
	if err != nil {
		fmt.Println("Error connecting to server, exiting..")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	reader2 := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("\n>>")

		input, _ := reader.ReadString('\n')

		// To exit from the prompt
		if strings.Trim(input, " \n\r") == "exit" {

			conn.Close()
			break
		}

		conn.Write([]byte(input))

		if input[:4] == "edit" {

			reply := make([]byte, 1024)
			_, _ = conn.Read(reply)

			if strings.Compare(string(reply), "err") == 0 {
				fmt.Println("Not happening")
			} else {

				content, err := reader2.ReadString('#')
				if err != nil {
					fmt.Println("Reader not wokring")
				}

				conn.Write([]byte(content))
			}
		}

		reply := make([]byte, 1024)
		_, err := conn.Read(reply)
		if err != nil {
			fmt.Println("Not happening")
		}
		fmt.Println()
		log.Println(string(reply))
	}
}
