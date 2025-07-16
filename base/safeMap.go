package base

import "sync"

type SafeMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: make(map[K]V)}
}

func (s *SafeMap[K, V]) Set(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap[K, V]) Get(key K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

// 只拿值，不存在时返回零值
func (m *SafeMap[K, V]) MustGet(key K) V {
	val, _ := m.Get(key)
	return val
}

func (s *SafeMap[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap[K, V]) Has(key K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.m[key]
	return ok
}

// 这种方法在 遍历的时候不能 对自身执行其他操作
func (s *SafeMap[K, V]) Range(f func(K, V) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.m {
		if !f(k, v) {
			break
		}
	}
}

// 这里是拷贝快照，然后在执行操作
func (s *SafeMap[K, V]) Snapshot() map[K]V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	copyData := make(map[K]V, len(s.m))
	for k, v := range s.m {
		copyData[k] = v
	}
	return copyData
}

func (s *SafeMap[K, V]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}
