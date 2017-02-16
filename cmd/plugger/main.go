package main

import (
	"flag"
	"log"
	"os"
	"plugtest"
)

func main() {
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if err := plugtest.LoadPlugins(dir); err != nil {
		log.Fatalln("error loading plugins:", err)
	}

	fn := plugtest.Get(flag.Arg(0))
	if fn == nil {
		log.Fatalln("plugin not found")
	}

	fn()
}
