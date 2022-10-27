package main

import (
	"TCPChat/internal"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port := ""
	switch len(os.Args) {
	case 1:
		port = "8989"
	case 2:
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argument shoud be integer for port variable")
			return
		}
		port = os.Args[1]
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	fmt.Printf("Listening on the port : %v", port)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	serv := internal.NewServer()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serv.Chat(conn)
	}
}
