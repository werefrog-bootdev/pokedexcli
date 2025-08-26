package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddGetAndReap(t *testing.T) {
	c := NewCache(50 * time.Millisecond)
	defer c.Close()

	key := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	val := []byte("hello")

	c.Add(key, val)

	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected cache hit")
	}
	if !bytes.Equal(got, val) {
		t.Fatalf("expected %q, got %q", val, got)
	}

	// Mutate the returned slice; ensure the cache keeps its own copy.
	got[0] = 'H'
	got2, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected cache hit after mutation")
	}
	if !bytes.Equal(got2, val) {
		t.Fatalf("cache value should be unchanged; want %q, got %q", val, got2)
	}

	// Wait long enough for the reaper to evict the entry.
	time.Sleep(120 * time.Millisecond)
	if _, ok := c.Get(key); ok {
		t.Fatalf("expected entry to be reaped after interval")
	}
}
