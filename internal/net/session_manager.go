package net

import (
	"log"
	"sync"
	"time"
)

type SessionManagerStruct struct {
	sessions map[int]*Session
	mu       sync.RWMutex
}

var SessionManager = &SessionManagerStruct{
	sessions: make(map[int]*Session),
}

func (m *SessionManagerStruct) Add(userID int, session *Session) {
	log.Printf("🔗 Adding session for user %d", userID)
	m.mu.Lock()
	defer m.mu.Unlock()
	if old, ok := m.sessions[userID]; ok && old != session {
		go func() {
			old.Kick("Tài khoản đã đăng nhập ở nơi khác")
			time.Sleep(2 * time.Second) // Đợi 1 giây để đảm bảo gói tin được gửi
			old.conn.Close()
		}()
	}
	m.sessions[userID] = session
}

func (s *Session) OnLoginSuccess(userID int) {
	log.Printf("✅ User %d logged in successfully", userID)
	s.ClientSessionID = &userID
	SessionManager.Add(userID, s)
}

func (m *SessionManagerStruct) Remove(userID int) {
	log.Printf("🔌 Removing session for user %d", userID)
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, userID)
}

func (m *SessionManagerStruct) IsOnline(userID int) bool {
	log.Printf("🔍 Checking if user %d is online", userID)
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.sessions[userID]
	return ok
}

func (m *SessionManagerStruct) GetSession(userID int) (*Session, bool) {
	log.Printf("🔎 Getting session for user %d", userID)
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[userID]
	return s, ok
}
