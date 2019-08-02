package main

import "github.com/goasm/gocef"

func main() {
	gocef.Initialize()
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
