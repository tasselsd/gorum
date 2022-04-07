package session

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/tasselsd/gorum/pkg/core"
)

type SessionManager interface {
	Save(*Session, time.Duration)
	LoadSession(token string) *Session
}

type InMemorySessionManager struct {
	store cmap.ConcurrentMap
}

func newInMemorySessionManager() *InMemorySessionManager {
	inMem := InMemorySessionManager{}
	inMem.store = cmap.New()
	return &inMem
}
func (sm *InMemorySessionManager) Save(s *Session, d time.Duration) {
	sm.store.Set(s.token, s)
}

func (sm *InMemorySessionManager) LoadSession(token string) *Session {
	s, ret := sm.store.Get(token)
	if ret {
		return s.(*Session)
	}
	return nil
}

var (
	sessionManager = newInMemorySessionManager()
	NaS            *Session
)

type Session struct {
	core.User
	token string
}

func NewSession(user *core.User) *Session {
	s := Session{User: *user}
	s.token = NewTokenString()
	sessionManager.Save(&s, 24*time.Hour)
	return &s
}

func SessionFromToken(token string) (*Session, error) {
	s := sessionManager.LoadSession(token)
	if s != nil {
		return s, nil
	}
	return nil, errors.New("Unauthorized")
}

func (s *Session) Token() string {
	return s.token
}

func (s *Session) JSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func NewTokenString() string {
	return core.NewSha1Object(uuid.NewString()).Sha1()
}
