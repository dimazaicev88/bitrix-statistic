package session

import (
	"strconv"
	"sync"
)

type Session struct {
	sync.RWMutex
	session map[string]string
}

func NewSession() *Session {
	return &Session{
		session: make(map[string]string),
	}
}

func (s *Session) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.session[key] = value
}

func (s *Session) Get(key string) string {
	s.Lock()
	defer s.Unlock()
	return s.session[key]
}

func (s *Session) GetAsInt(key string) int {
	s.Lock()
	defer s.Unlock()
	atoi, err := strconv.Atoi(s.session[key])
	if err != nil {
		return 0
	}
	return atoi
}

func (s *Session) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.session, key)
}

func (s *Session) Count() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.session)
}

func (s *Session) KeyExists(key string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.session[key]
	return ok
}

func (s *Session) SetAll(sess *Session) {
	s.Lock()
	defer s.Unlock()
	for key, val := range sess.session {
		s.session[key] = val
	}
}
