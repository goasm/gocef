package main

import "github.com/goasm/gocef"

func main() {
	if gocef.ExecuteProcess() > -1 {
		return
	}
	gocef.Initialize()
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
