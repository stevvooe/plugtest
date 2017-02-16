package main

import (
	"fmt"
	"plugtest"
)

func init() {
	plugtest.Register("plugin1", func() {
		fmt.Println("plugin1")
	})
}
