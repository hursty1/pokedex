package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		input string
		val []byte
	}{
		{
		input:"http://boot.dev",
		val:[]byte("testdata"),
		},
		{
		input:"http://boot.dev/2",
		val:[]byte("testdata2"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.input, c.val)
			val, ok := cache.Get(c.input)
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
	return
}

func TestReap(t *testing.T) {
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
	return
}