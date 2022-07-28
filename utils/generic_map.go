package utils

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
