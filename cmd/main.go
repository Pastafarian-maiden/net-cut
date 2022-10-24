package main

import (
	"fmt"
	"log"
	"net"
	"net-cut/internal"
	"os"
	"strconv"
)

func main() {
	port := ""
	switch len(os.Args) {
	case 1:
		port = "8989"
		fmt.Printf("Listening on the port :%v", port)
	case 2:
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argument shoud be integer for port variable")
			return
		}
		fmt.Printf("Listening on the port :%v", port)
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
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
