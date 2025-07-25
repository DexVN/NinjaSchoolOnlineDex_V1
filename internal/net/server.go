package net

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"nso-server/internal/pkg/config"
	"nso-server/internal/pkg/logger"

	"go.uber.org/fx"
)

type Server struct {
	port     string
	listener net.Listener
	router   RouterFunc
	wg       sync.WaitGroup
	sessions sync.Map
}

func NewServer(cfg *config.Config, router RouterFunc) (*Server, error) {
	port := cfg.ServerPort
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	logger.Infof("🌐 Listening on %s", port)

	return &Server{
		port:     port,
		listener: ln,
		router:   router,
	}, nil
}

func (s *Server) Start() {
	logger.Info("🚀 Server started, waiting for connections...")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if isNetClosedError(err) {
				logger.Info("🛑 Listener closed, stop accepting new connections.")
				break
			}
			logger.WithError(err).Warn("⚠️ Accept failed")
			continue
		}

		s.wg.Add(1)

		go func() {
			defer s.wg.Done()
			session := NewSession(conn, s.router)
			s.sessions.Store(session, struct{}{})
			defer s.sessions.Delete(session)
			session.Start()
		}()
	}
}

func (s *Server) Stop() {
	logger.Info("🛑 Stopping server...")
	_ = s.listener.Close()

	// Kick tất cả session
	s.sessions.Range(func(k, _ any) bool {
		if sess, ok := k.(*Session); ok {
			sess.Kick(true) // Đóng kết nối nhẹ nhàng
		}
		return true
	})

	// Chờ tất cả session đóng
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logger.Info("✅ All sessions closed")
	case <-time.After(5 * time.Second):
		logger.Warn("⚠️ Timeout: Some sessions may still be hanging")
	}
}

func (s *Server) handleConn(conn net.Conn) {
	session := NewSession(conn, s.router)
	session.Start()
}

func Serve(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			go server.Stop()
			return nil
		},
	})
}

// isNetClosedError kiểm tra lỗi từ listener đã đóng
func isNetClosedError(err error) bool {
	return errors.Is(err, net.ErrClosed) ||
		strings.Contains(err.Error(), "use of closed network connection")
}
