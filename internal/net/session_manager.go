package net

import (
	"sync"
	"time"

	logger "nso-server/internal/infra"
	"nso-server/internal/lang"
	"nso-server/internal/proto"
)

type SessionManagerStruct struct {
	sessions      map[int]*Session
	pendingByUser map[string][]*Session
	mu            sync.RWMutex
}

var SessionManager = &SessionManagerStruct{
	sessions:      make(map[int]*Session),
	pendingByUser: make(map[string][]*Session),
}

func (m *SessionManagerStruct) Add(userID int, session *Session) {
	logger.Log.Infof("🔗 Adding session for user %d", userID)

	m.mu.Lock()
	defer m.mu.Unlock()

	if old, ok := m.sessions[userID]; ok && old != session {
		// 🔴 Gửi thông báo tới session cũ
		go func(s *Session) {
			s.SendMessageWithCommand(proto.CmdServerDialog, dialogWriter(lang.Get("account.logged_in_elsewhere")))
			time.Sleep(3000 * time.Millisecond)
			s.Kick(true)
			s.Cleanup()
		}(old)

		// 🟢 Gửi thông báo tới session mới (đang login)
		go func(s *Session) {
			s.SendMessageWithCommand(proto.CmdServerDialog, dialogWriter(lang.Get("account.logged_in_elsewhere_new")))
			time.Sleep(3000 * time.Millisecond)
			s.Kick(true)
			s.Cleanup()
		}(session)
		return // Không ghi đè session
	}

	// ✅ Nếu không có session cũ thì gán session mới
	m.sessions[userID] = session
}

func (s *Session) OnLoginSuccess(userID int) {
	logger.Log.Infof("✅ User %d logged in successfully", userID)
	s.ClientSessionID = &userID
	SessionManager.Add(userID, s)
}

func (m *SessionManagerStruct) Remove(userID int) {
	logger.Log.Infof("🔌 Removing session for user %d", userID)
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, userID)
}

func (m *SessionManagerStruct) IsOnline(userID int) bool {
	logger.Log.Infof("🔍 Checking if user %d is online", userID)
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.sessions[userID]
	return ok
}

func (m *SessionManagerStruct) GetSession(userID int) (*Session, bool) {
	logger.Log.Infof("🔎 Getting session for user %d", userID)
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[userID]
	return s, ok
}

func dialogWriter(text string) *proto.Writer {
	w := proto.NewWriter()
	w.WriteUTF(text)
	return w
}
