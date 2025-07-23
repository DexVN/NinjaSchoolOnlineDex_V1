package net

import (
	"net"

	logger "nso-server/internal/infra"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	logger.Log.Infof("🔌 New client connected: %s", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			logger.Log.Warnf("❌ Disconnected: %s", conn.RemoteAddr())
			return
		}
		logger.Log.Infof("📨 Received %d bytes from %s: %x", n, conn.RemoteAddr(), buf[:n])
	}
}
