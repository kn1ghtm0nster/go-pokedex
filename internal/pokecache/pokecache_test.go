package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestNewCache(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	if cache == nil {
		t.Errorf("expected to create cache")
		return
	}

	if cache.entries == nil {
		t.Errorf("expected to create entries map")
		return
	}

	if len(cache.entries) != 0 {
		t.Errorf("expected entries map to be empty")
		return
	}
}

func TestGetMissingKey(t *testing.T) {
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("doesn-not-exist")
	if ok {
		t.Errorf("expected to not find key")
	}
}

func TestAddDuplicateKey(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("https://example.com", []byte("testdata1"))
	cache.Add("https://example.com", []byte("some other data here"))

	val, ok := cache.Get("https://example.com")

	// Ensure that the key is actually present
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	// Ensure that the value is the updated one
	if string(val) != "some other data here" {
		t.Errorf("expected to find updated value")
		return
	}
}