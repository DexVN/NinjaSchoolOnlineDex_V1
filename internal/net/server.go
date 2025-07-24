package net

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"

	"go.uber.org/fx"
	"nso-server/internal/pkg/config"
	"nso-server/internal/pkg/logger"
)

type Server struct {
	port     string
	listener net.Listener
	router   RouterFunc
	wg       sync.WaitGroup
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
			s.handleConn(conn)
		}()
	}
}

func (s *Server) Stop() error {
	if s.listener == nil {
		logger.Warn("🛑 Server was never started or already stopped.")
		return nil
	}

	logger.Info("🛑 Stopping server...")
	if err := s.listener.Close(); err != nil {
		logger.WithError(err).Error("❌ Error closing listener")
		return err
	}
	s.wg.Wait()
	logger.Info("✅ Server shutdown complete")
	return nil
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
			return server.Stop()
		},
	})
}

// isNetClosedError kiểm tra lỗi từ listener đã đóng
func isNetClosedError(err error) bool {
	return errors.Is(err, net.ErrClosed) ||
		strings.Contains(err.Error(), "use of closed network connection")
}
