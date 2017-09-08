package model

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
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

func GetSession(r *http.Request) *Session {
	id, err := r.Cookie("session")
	if err != nil {
		return CreateSession()
	}

	sessionStore.RLock()
	defer sessionStore.RUnlock()
	if sessionStore.data == nil {
		return CreateSession()
	}
	s := sessionStore.data[id.Value]
	if s == nil {
		return CreateSession()
	}
	return s
}

func (s *Session) Save(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    s.ID,
		MaxAge:   int(10 * time.Minute / time.Second), //10นาที
		HttpOnly: true,                                //ป้องกันการใช้ JavaScript อ่านค่า
		Path:     "/",
		//Secure:   true,                                //จะส่ง cookie ต่อเมื่อใช้ https เท่านั้น
	})

	sessionStore.Lock()
	defer sessionStore.Unlock()
	if sessionStore.data == nil { //ถ้่า session ไม่มีค่า
		sessionStore.data = make(map[string]*Session) //ให้สร้างขึ้นมาใหม่
	}
	sessionStore.data[s.ID] = s
}
