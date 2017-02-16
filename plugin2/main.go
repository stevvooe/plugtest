package main

import (
	"fmt"
	"plugtest"
)

func init() {
	plugtest.Register("plugin2", func() {
		fmt.Println("plugin2")
	})
}
