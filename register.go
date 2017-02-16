package plugtest

import (
	"fmt"
	"log"
	"path/filepath"
	"plugin"
	"sync"
)

var (
	plugins = map[string]func(){}
	mu      sync.Mutex
)

func Register(name string, fn func()) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := plugins[name]; ok {
		panic(fmt.Sprintln(name, "already registered"))
	}

	log.Println(name, "registered")
	plugins[name] = fn
}

func Get(name string) func() {
	mu.Lock()
	defer mu.Unlock()

	return plugins[name]
}

func LoadPlugins(dir string) error {
	paths, err := filepath.Glob(filepath.Join(dir, "*.so"))
	if err != nil {
		return err
	}

	for _, path := range paths {
		if _, err := plugin.Open(path); err != nil {
			return err
		}
	}

	return nil
}
