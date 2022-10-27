package internal

import (
	"net"
	"sync"
)

const (
	Pinguin    = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n	 `-'       `--'\n[ENTER YOUR NAME]: "
	TimeFormat = "2006-01-02 15:04:05"
	NewLine    = "\033[2K\r"
)

type Server struct {
	mu      sync.Mutex
	room    map[string]net.Conn
	history string
}

func NewServer() *Server {
	return &Server{
		room: make(map[string]net.Conn),
	}
}
