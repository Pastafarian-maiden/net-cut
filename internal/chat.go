package internal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

// Chat is handling connection for each chat user
func (s *Server) Chat(conn net.Conn) {
	if len(s.room) >= 10 {
		io.WriteString(conn, "No spaces in chat room")
		return
	}

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

	io.WriteString(conn, s.history)

	status := "newUser"
	message := MessageFormat(username, "", status)
	s.SendMessage(message, username)
	LineFormat(username, conn)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			LineFormat(username, conn)
			continue
		}
		status = ""
		message = MessageFormat(username, text, status)
		s.SendMessage(message, username)
		LineFormat(username, conn)
	}

	s.mu.Lock()
	delete(s.room, username)
	s.mu.Unlock()

	status = "deleteUser"
	message = MessageFormat(username, "", status)
	s.SendMessage(message, username)
}

// LineFormat is formating new line according to the required template
func LineFormat(username string, conn net.Conn) {
	time := time.Now().Format(TimeFormat)
	format := NewLine + "[" + time + "]" + "[" + username + "]: "
	io.WriteString(conn, format)
}

// MessageFormat is formating message according to the required template
func MessageFormat(username, text, status string) (message string) {
	if status == "newUser" {
		message = NewLine + username + " has joined our chat...\n"
	} else if status == "deleteUser" {
		message = NewLine + username + " has left our chat...\n"
	} else {
		time := time.Now().Format(TimeFormat)
		message = fmt.Sprintf("%v[%v][%v]: %v\n", NewLine, time, username, text)
	}
	return
}

// SendMessage is sending message to chat users
func (s *Server) SendMessage(message, sender string) {
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
