package main

import (
	"os"

	"github.com/goasm/gocef"
)

func main() {
	if retcode := gocef.ExecuteProcess(); retcode >= 0 {
		os.Exit(retcode)
	}
	gocef.Initialize()
	gocef.CreateBrowser()
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
