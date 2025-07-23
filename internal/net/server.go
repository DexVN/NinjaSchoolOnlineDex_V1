package net

import (
	"net"

	logger "nso-server/internal/infra"
)

// RouterFunc kiểu hàm xử lý message
type Server struct {
	listener net.Listener
	router   RouterFunc
}

// NewServer tạo server với router callback
func NewServer(addr string, router RouterFunc) (*Server, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Server{listener: ln, router: router}, nil
}

func (s *Server) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logger.Log.WithError(err).Warn("⚠️ Accept error")
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	session := NewSession(conn, s.router) // truyền router vào session
	session.Start()
}
