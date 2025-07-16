package base

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	lock     int32      // 0: unlocked, 1: locked
	ownerMu  sync.Mutex // 保护下面两个字段
	ownerID  string     // 当前持有锁的 goroutine/任务 ID
	recCount int32      // 重入次数
}

func (s *SpinLock) Lock(ownerID string) {
	// 是否是当前锁的拥有者？
	s.ownerMu.Lock()
	if s.ownerID == ownerID {
		s.recCount++
		s.ownerMu.Unlock()
		return
	}
	s.ownerMu.Unlock()

	// 自旋直到获得锁
	for !atomic.CompareAndSwapInt32(&s.lock, 0, 1) {
		runtime.Gosched()
	}

	s.ownerMu.Lock()
	s.ownerID = ownerID
	s.recCount = 1
	s.ownerMu.Unlock()
}

func (s *SpinLock) Unlock(ownerID string) {
	s.ownerMu.Lock()
	defer s.ownerMu.Unlock()

	if s.ownerID != ownerID {
		panic("SpinLock: unlock called by non-owner")
	}

	s.recCount--
	if s.recCount == 0 {
		s.ownerID = ""
		atomic.StoreInt32(&s.lock, 0)
	}
}

func (s *SpinLock) TryLock(ownerID string, timeout time.Duration) bool {
	s.ownerMu.Lock()
	if s.ownerID == ownerID {
		s.recCount++
		s.ownerMu.Unlock()
		return true
	}
	s.ownerMu.Unlock()

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if atomic.CompareAndSwapInt32(&s.lock, 0, 1) {
			s.ownerMu.Lock()
			s.ownerID = ownerID
			s.recCount = 1
			s.ownerMu.Unlock()
			return true
		}
		runtime.Gosched()
	}
	return false
}
