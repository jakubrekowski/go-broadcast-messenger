package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func broadcastDial(message string) {
	conn, err := net.Dial("udp", "192.168.100.255:9000")
		if err != nil {
			panic(err)
		}

		defer conn.Close()

		io.WriteString(conn, message)
		
		conn.Close()
}

func keyboard(nick string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		if message == "\n" {
			continue
		}

		broadcastDial(fmt.Sprint("\033[33m", nick, "% \033[0m", message))
	}
}

func main() {
  fmt.Println("Broadcast Messenger v. alpha0.3")

// Nick selector

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter nick: ")
	nick, _ := reader.ReadString('\n')

	if nick == "\n" {
		nick = "anonymous"
	}

	nick = strings.Replace(nick, "\n", "", -1)

	broadcastDial(fmt.Sprint("\033[32m> ", nick, " has join the conversation.\n"))

// main loop

	go keyboard(nick)

	addr := net.UDPAddr{
		Port: 9000,
		IP:  	net.ParseIP("192.168.100.255"),
	}

	ln, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		// conn, err := ln.Accept()
		// if err != nil {
		// 	panic(err)
		// }
		
		for {
			bs := make([]byte, 1024)
			_, err := ln.Read(bs)
			if err != nil {
				break
			}

			fmt.Print(string(bs[:]))
		}

		ln.Close()
	}
}
