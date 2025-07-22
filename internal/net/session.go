// internal/net/session.go
package net

import (
	"encoding/binary"
	"io"
	"log"
	"net"

	"nso-server/internal/proto"
)

const (
	cmdKeyHandshake = byte(0xE5)
)

type Session struct {
	conn       net.Conn
	router     RouterFunc
	key        []byte
	readIndex  int
	writeIndex int
	ClientSessionID *int
}

type RouterFunc func(msg *proto.Message, s *Session)

func NewSession(conn net.Conn, router RouterFunc) *Session {
	return &Session{
		conn:   conn,
		router: router,
	}
}

func (s *Session) Start() {
	defer s.conn.Close()
	log.Println("🔄 New connection from", s.conn.RemoteAddr())

	// Đọc gói đầu tiên (thường là handshake)
	opcodeBuf := make([]byte, 1)
	if _, err := io.ReadFull(s.conn, opcodeBuf); err != nil {
		log.Println("❌ Failed to read opcode:", err)
		return
	}
	opcode := int8(opcodeBuf[0])

	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(s.conn, lenBuf); err != nil {
		log.Println("❌ Failed to read length:", err)
		return
	}
	length := binary.BigEndian.Uint16(lenBuf)

	if length > 0 {
		drop := make([]byte, length)
		if _, err := io.ReadFull(s.conn, drop); err != nil {
			log.Println("❌ Failed to read payload:", err)
			return
		}
	}

	if opcode == proto.CmdGetSessionID {
		if err := s.sendHandshake(); err != nil {
			log.Println("❌ Send handshake failed:", err)
			return
		}
	}

	log.Println("✅ Handshake complete, start encrypted message loop")

	// Đọc các message XOR sau khi có key
	for {
		msg, err := proto.ReadMessage(s.conn, s.key, &s.readIndex)
		if err != nil {
			if err != io.EOF {
				log.Println("⚠️ ReadMessage error:", err)
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
	buf[0] = cmdKeyHandshake
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

	log.Printf("🔐 XOR key activated: % X", s.key)
	return nil
}

func (s *Session) SendMessage(msg *proto.Message) error {
	return proto.WriteMessage(s.conn, *msg, s.key, &s.writeIndex)
}

func (s *Session) Conn() net.Conn {
	return s.conn
}