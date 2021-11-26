package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func broadcastDial(message string) {
	conn, err := net.Dial("tcp", "192.168.100.103:9000")
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		io.WriteString(conn, message)
		
		conn.Close()
}

func keyboard() {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		broadcastDial(message)		
	}
}

func main() {
  fmt.Println("Broadcast Messenger v. alpha0.1")

	go keyboard()

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		
		for {
			bs := make([]byte, 1024)
			_, err := conn.Read(bs)
			if err != nil {
				break
			}

			fmt.Print("[", conn.RemoteAddr(), "] ", string(bs[:]))
		}

		conn.Close()
	}
}
