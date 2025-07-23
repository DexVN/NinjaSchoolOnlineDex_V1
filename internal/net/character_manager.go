package net

import (
	"sync"
)

type CharacterManagerStruct struct {
	characters map[int]*Session
	mu         sync.RWMutex
}

var CharacterOnlineManager = &CharacterManagerStruct{
	characters: make(map[int]*Session),
}

func (m *CharacterManagerStruct) Add(characterID int, s *Session) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.characters[characterID] = s
}

func (m *CharacterManagerStruct) Remove(characterID int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.characters, characterID)
}

func (m *CharacterManagerStruct) IsOnline(characterID int) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.characters[characterID]
	return ok
}

func (m *CharacterManagerStruct) GetSession(characterID int) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.characters[characterID]
	return s, ok
}
