package utils

import "sync"

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m Map[K, V]) KeyExists(key K) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

func (m Map[K, V]) Map() map[K]V {
	return m
}

func (m Map[K, V]) Del(key K) {
	delete(m, key)
}

func (m Map[K, V]) Value(field K) V {
	return m[field]
}

// Set ItemSet the set of Items
type Set[V comparable] struct {
	items map[V]bool
	lock  sync.RWMutex
}

func (s *Set[V]) SliceAsSet(t []V) *Set[V] {
	for _, item := range t {
		s.Add(item)
	}
	return s
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *Set[V]) Add(t V) *Set[V] {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[V]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *Set[V]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[V]bool)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *Set[V]) Delete(item V) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *Set[V]) Has(item V) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *Set[V]) Items() []V {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var items []V
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *Set[V]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

func NewSet[V comparable]() *Set[V] {
	return &Set[V]{}
}
