package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
	stop     chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		stop:     make(chan struct{}),
	}
	go c.reapLoop()
	return c
}

// Add stores a defensive copy of val.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       append([]byte(nil), val...),
	}
	c.mu.Unlock()
}

// Get returns a defensive copy of the cached value.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	e, ok := c.entries[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return append([]byte(nil), e.val...), true
}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			cutoff := time.Now().Add(-c.interval)
			c.mu.Lock()
			for k, e := range c.entries {
				if e.createdAt.Before(cutoff) {
					delete(c.entries, k)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}

// Close stops the reaper goroutine (useful in tests / shutdown).
func (c *Cache) Close() {
	select {
	case <-c.stop:
		// already closed
	default:
		close(c.stop)
	}
}
