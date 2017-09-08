package model

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type Session struct {
	ID     string
	UserID string
}

var sessionStore struct {
	sync.RWMutex
	data map[string]*Session
}

func generateID() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return base64.URLEncoding.EncodeToString(buf)
	//return base64.StdEncoding.EncodeToString(buf)
}

func CreateSession() *Session {
	return &Session{
		ID: generateID(),
	}
}

func SaveSession(s *Session) {
	sessionStore.Lock()
	defer sessionStore.Unlock()
	if sessionStore.data == nil { //ถ้่า session ไม่มีค่า
		sessionStore.data = make(map[string]*Session) //ให้สร้างขึ้นมาใหม่
	}
	sessionStore.data[s.ID] = s
}

func GetSession(id string) *Session {
	sessionStore.RLock()
	defer sessionStore.RUnlock()
	if sessionStore.data == nil {
		return nil
	}
	return sessionStore.data[id]
}
