package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	water    int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		water:    0,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) leak(now time.Time) {
	if lb.leakRate <= 0 {
		return
	}

	elapsed := now.Sub(lb.lastLeak)
	leaked := int(elapsed / lb.leakRate)
	if leaked <= 0 {
		return
	}

	lb.water -= leaked
	if lb.water < 0 {
		lb.water = 0
	}
	lb.lastLeak = lb.lastLeak.Add(time.Duration(leaked) * lb.leakRate)
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.leak(time.Now())

	if lb.water < lb.capacity {
		lb.water++
		return true
	}
	return false
}

func main() {
	lb := NewLeakyBucket(5, 500*time.Millisecond)

	println("--- burst (should allow 5, deny the rest) ---")
	for range 8 {
		if lb.Allow() {
			println("Request allowed")
		} else {
			println("XXX----- Request denied")
		}
	}

	println("--- wait for 3 units to leak ---")
	time.Sleep(3 * 500 * time.Millisecond)

	println("--- after leak (should allow 3 more) ---")
	for range 5 {
		if lb.Allow() {
			println("Request allowed")
		} else {
			println("XXX----- Request denied")
		}
	}
}
