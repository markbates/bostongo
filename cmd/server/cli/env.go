package cli

import (
	"os"
	"sync"
)

// snippet: type

type Env struct {
	data map[string]string
	mu   sync.RWMutex
}

// snippet: type

func (e *Env) Setenv(key string, value string) {
	if e == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.data == nil {
		e.data = map[string]string{}
	}
	e.data[key] = value
}

// snippet: getenv

func (e *Env) Getenv(key string) string {
	if e == nil || e.data == nil {
		return os.Getenv(key)
	}

	e.mu.RLock()
	defer e.mu.RUnlock()

	if k, ok := e.data[key]; ok {
		return k
	}

	return os.Getenv(key)
}

// snippet: getenv
