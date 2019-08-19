package main

import "github.com/goasm/gocef"

func main() {
	if gocef.ExecuteProcess() > -1 {
		return
	}
	gocef.Initialize()
	gocef.CreateBrowser()
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
