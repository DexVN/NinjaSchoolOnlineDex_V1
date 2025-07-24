package net

import (
	"net"

	"nso-server/internal/pkg/logger"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	logger.Infof("🔌 New client connected: %s", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			logger.Warnf("❌ Disconnected: %s", conn.RemoteAddr())
			return
		}
		logger.Infof("📨 Received %d bytes from %s: %x", n, conn.RemoteAddr(), buf[:n])
	}
}
