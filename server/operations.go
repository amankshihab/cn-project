package main

import (
	"bufio"
	"bytes"
	"net"
	"os"
	"strings"
)

func createFile(filename string, c net.Conn) {
	if _, err := os.Stat(strings.Trim(filename[7:], " \n\r")); os.IsNotExist(err) {

		file, err := os.Create(strings.Trim(filename[7:], " \n\r"))
		if err != nil {
			c.Write([]byte("Error while creating file, try again later."))
			return
		}

		file.Close()

		c.Write([]byte("File created"))
	} else {
		c.Write([]byte("File with the same name already exists."))
	}
}

func catCommand(filename string, c net.Conn) {
	content, err := os.ReadFile(strings.Trim(filename[4:], " \n\r"))
	if os.IsNotExist(err) {
		c.Write([]byte("File doesn't exist."))
		return
	} else if err != nil {
		c.Write([]byte("Error while reading file."))
		return
	}

	if len(content) != 0 {
		c.Write(content)
	} else {
		c.Write([]byte(" "))
	}
}

func deleteFile(filename string, c net.Conn) {
	err := os.Remove(strings.Trim(filename[7:], " \n\r"))
	if os.IsNotExist(err) {
		c.Write([]byte("Error deleting the files"))
		return
	} else if err != nil {
		c.Write([]byte("Error deleting files"))
		return
	}
	c.Write([]byte("File Deleted"))
}

func editFile(filename string, c net.Conn) {
	file, err := os.OpenFile(strings.Trim(filename[5:], " \n\r"), os.O_RDWR|os.O_APPEND, 0666)
	if os.IsNotExist(err) {
		c.Write([]byte("File doesn't exist."))
		return
	} else if err != nil {
		c.Write([]byte("Error accessing file."))
		return
	}

	c.Write([]byte("nil"))

	content, err := bufio.NewReader(c).ReadBytes('#')
	if err != nil {
		c.Write([]byte("Error retreiving data"))
	}

	_, err = file.Write(bytes.Trim(content, "#"))
	if err != nil {
		c.Write([]byte("Error writing into file."))
	}

	c.Write([]byte("\nFile edited."))

	file.Close()
}
