package cli

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Commands struct {
	routes map[string]Commander
	mu     sync.RWMutex
}

func (c *Commands) Find(name string) (Commander, bool) {
	if c == nil {
		return nil, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	cmd, ok := c.routes[name]
	if !ok || cmd == nil {
		return nil, false
	}

	return cmd, true
}

func (c *Commands) Set(name string, cmd Commander) {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.routes == nil {
		c.routes = map[string]Commander{}
	}

	c.routes[name] = cmd
}

func (c *Commands) Print(w io.Writer) error {
	if c == nil {
		return fmt.Errorf("nil commands")
	}

	if w == nil {
		w = os.Stdout
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, cmd := range c.routes {
		fmt.Fprintf(w, "%T\n", cmd)
	}

	return nil
}

func (c *Commands) Map() map[string]Commander {
	if c == nil {
		return nil
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	m := c.routes
	if m == nil {
		m = map[string]Commander{}
	}

	return m
}
