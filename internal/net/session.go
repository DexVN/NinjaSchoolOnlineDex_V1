package net

import (
	"encoding/binary"
	"io"
	"net"
	"time"

	"nso-server/internal/pkg/logger"
	"nso-server/internal/proto"
)

const (
	CmdKeyHandshake = byte(0xE5)
)

type Session struct {
	conn            net.Conn
	router          RouterFunc
	key             []byte
	readIndex       int
	writeIndex      int
	ClientSessionID *int
	CharacterID     *int
	ServerID        *int
	ServerCode      string
}

type RouterFunc func(msg *proto.Message, s *Session)

func NewSession(conn net.Conn, router RouterFunc) *Session {
	return &Session{
		conn:   conn,
		router: router,
	}
}

func (s *Session) Start() {
	defer func() {
		s.Cleanup()    // 🧹 cleanup khỏi SessionManager, CharacterManager,...
		s.conn.Close() // đóng kết nối sau cleanup
	}()
	logger.Infof("🔄 New connection from %s", s.conn.RemoteAddr())

	// Đọc gói đầu tiên (thường là handshake)
	opcodeBuf := make([]byte, 1)
	if _, err := io.ReadFull(s.conn, opcodeBuf); err != nil {
		logger.WithError(err).Error("❌ Failed to read opcode")
		return
	}
	opcode := int8(opcodeBuf[0])

	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(s.conn, lenBuf); err != nil {
		logger.WithError(err).Error("❌ Failed to read length")
		return
	}
	length := binary.BigEndian.Uint16(lenBuf)

	if length > 0 {
		drop := make([]byte, length)
		if _, err := io.ReadFull(s.conn, drop); err != nil {
			logger.WithError(err).Error("❌ Failed to read payload")
			return
		}
	}

	if opcode == proto.CmdGetSessionId {
		if err := s.sendHandshake(); err != nil {
			logger.WithError(err).Error("❌ Send handshake failed")
			return
		}
	}

	logger.Info("✅ Handshake complete, start encrypted message loop")

	for {
		msg, err := proto.ReadMessage(s.conn, s.key, &s.readIndex)
		if err != nil {
			if err != io.EOF {
				logger.WithError(err).Warn("⚠️ ReadMessage error")
			}
			break
		}
		s.router(msg, s)
	}
}

func (s *Session) sendHandshake() error {
	key := []byte{'D'}
	payload := append([]byte{byte(len(key))}, key...)

	buf := make([]byte, 1+2+len(payload))
	buf[0] = CmdKeyHandshake
	binary.BigEndian.PutUint16(buf[1:3], uint16(len(payload)))
	copy(buf[3:], payload)

	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}

	s.key = make([]byte, len(key))
	copy(s.key, key)
	for i := 1; i < len(s.key); i++ {
		s.key[i] ^= s.key[i-1] // giống client Unity
	}

	logger.Infof("🔐 XOR key activated: % X", s.key)
	return nil
}

func (s *Session) SendMessage(msg *proto.Message) error {
	return proto.WriteMessage(s.conn, msg, s.key, &s.writeIndex)
}

func (s *Session) SendMessageWithCommand(cmd int8, w *proto.Writer) error {
	msg := proto.NewMessage(cmd)
	msg.Writer().WriteBytes(w.GetData())
	return s.SendMessage(msg)
}

func (s *Session) Kick(forceClose bool) {
	if forceClose {
		go func() {
			time.Sleep(500 * time.Millisecond)
			s.conn.Close()
		}()
	}
}

func (s *Session) Cleanup() {
	logger.Info("🧹 Cleaning up session resources")
	if s.ClientSessionID != nil {
		SessionManager.Remove(*s.ClientSessionID)
	}
	if s.CharacterID != nil {
		CharacterOnlineManager.Remove(*s.CharacterID)
	}
}

func (s *Session) Conn() net.Conn {
	return s.conn
}
