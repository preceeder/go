package base

import (
	"sync"
	"testing"
	"time"
)

func TestSpinLock_MultiGoroutines(t *testing.T) {
	sl := &SpinLock{}
	counter := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id string) {
			for j := 0; j < 1000; j++ {
				sl.Lock(id)
				counter++
				sl.Unlock(id)
			}
			wg.Done()
		}("goroutine-" + string(rune(i)))
	}

	wg.Wait()
	if counter != 10000 {
		t.Errorf("Expected counter 10000, got %d", counter)
	}
}

func TestSpinLock_Reentrant(t *testing.T) {
	sl := &SpinLock{}
	id := "task-1"

	sl.Lock(id)
	sl.Lock(id)
	sl.Unlock(id)
	sl.Unlock(id)

	// If not panic, test passed.
}

func TestSpinLock_TryLockTimeout(t *testing.T) {
	sl := &SpinLock{}
	id1 := "owner1"
	id2 := "owner2"

	sl.Lock(id1)

	success := sl.TryLock(id2, 10*time.Millisecond)
	if success {
		t.Error("TryLock should have failed due to timeout")
	}

	sl.Unlock(id1)
}

// ------------------- BENCHMARK -------------------

func BenchmarkSpinLock(b *testing.B) {
	sl := &SpinLock{}
	ownerID := "bench"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sl.Lock(ownerID)
			sl.Unlock(ownerID)
		}
	})
}
