package internal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func (s *Server) Chat(conn net.Conn) {
	io.WriteString(conn, Pinguin)

	scanner := bufio.NewScanner(conn)
	var username string

	for scanner.Scan() {
		username = scanner.Text()
		if username == "" {
			io.WriteString(conn, "Please enter username.\n")
			continue
		}
		_, exist := s.room[username]
		if exist {
			io.WriteString(conn, "Username "+username+" is taken. Try something else.\n")
			continue
		}
		break
	}

	s.mu.Lock()
	s.room[username] = conn
	s.mu.Unlock()

	status := "newUser"
	message := MessageFormat(username, "", status)
	s.Send(message, username)
	LineFormat(username, conn)

	for scanner.Scan() {
		text := scanner.Text()
		status = ""
		message = MessageFormat(username, text, status)
		s.Send(message, username)
		LineFormat(username, conn)
	}
}

func LineFormat(username string, conn net.Conn) {
	time := time.Now().Format(TimeFormat)
	format := NewLine + "[" + time + "]" + "[" + username + "]: "
	io.WriteString(conn, format)
}

func MessageFormat(username, text, status string) (message string) {
	if status == "newUser" {
		message = NewLine + username + " has joined our chat...\n"
	} else {
		time := time.Now().Format(TimeFormat)
		message = fmt.Sprintf("%v[%v][%v]: %v\n", NewLine, time, username, text)
	}
	return
}

func (s *Server) Send(message, sender string) {
	s.mu.Lock()
	for username, conn := range s.room {
		if username != sender {
			io.WriteString(conn, message)
			LineFormat(username, conn)
		}
	}
	s.history += message
	s.mu.Unlock()
}
