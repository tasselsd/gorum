package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/tasselsd/gorum/pkg/core"
)

type SessionManager interface {
	Save(*Session, time.Duration)
	LoadSession(token string) *Session
	Remove(token string)
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
	sm.store.Set(s.token, map[string]any{
		"t": time.Now().UnixMilli(),
		"s": s,
	})
}

func (sm *InMemorySessionManager) LoadSession(token string) *Session {
	s, ret := sm.store.Get(token)
	if ret {
		t := s.(map[string]any)["t"].(int64)
		if t < time.Now().Add(-time.Hour*24).UnixMilli() {
			sm.store.Remove(token)
			ws := s.(map[string]any)["s"].(*Session).ws
			if ws != nil {
				ws.Close()
			}
			return nil
		}
		sS := s.(map[string]any)
		sS["t"] = time.Now().UnixMilli()
		return sS["s"].(*Session)
	}
	return nil
}

func (sm *InMemorySessionManager) Remove(token string) {
	sm.store.Remove(token)
}

var (
	sessionManager = newInMemorySessionManager()
	NaS            *Session
)

type Session struct {
	core.User
	token string
	ws    *websocket.Conn
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

func RemoveSession(token string) {
	sessionManager.Remove(token)
}

func (s *Session) Token() string {
	return s.token
}

func (s *Session) JSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s *Session) SetWs(ws *websocket.Conn) {
	s.ws = ws
}

func (s *Session) GetWs() (*websocket.Conn, error) {
	if s.ws == nil {
		return nil, errors.New("not connected")
	}
	return s.ws, nil
}

func NewTokenString() string {
	return core.NewSha1Object(uuid.NewString()).Sha1()
}

func init() {
	core.SHUTDOWN_HOOKS = append(core.SHUTDOWN_HOOKS, func() {

		for _, item := range sessionManager.store.Items() {
			ws := item.(map[string]any)["s"].(*Session).ws
			if ws != nil {
				ws.WriteMessage(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseGoingAway, "server shutdown"))
				fmt.Printf("[INFO] close websocket connection\n")
			}
		}

		v, err := sessionManager.store.MarshalJSON()
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(".gorum-sessions", v, fs.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[INFO] sessions persisted\n")
	})
	core.STARTUP_HOOKS = append(core.STARTUP_HOOKS, func() {
		b, err := ioutil.ReadFile(".gorum-sessions")
		if err != nil {
			fmt.Printf("[WARN] sessions load error [ %s ]\n", err.Error())
			return
		}
		var m map[string]interface{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			fmt.Printf("[WARN] sessions load error [ %s ]\n", err.Error())
			return
		}
		for k, v := range m {
			t := v.(map[string]any)["t"].(float64)
			if int64(t) < time.Now().Add(-time.Hour*24).UnixMilli() {
				continue
			}
			sMap := v.(map[string]any)["s"]
			s, _ := json.Marshal(sMap)
			var sess Session
			json.Unmarshal(s, &sess)
			sessionManager.store.Set(k, map[string]any{
				"t": int64(t),
				"s": &sess,
			})
		}
	})
}
