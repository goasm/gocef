package main

import (
	"os"

	"github.com/goasm/gocef"
)

type Client struct{}
type LifeSpanHandler struct{}

func (h LifeSpanHandler) OnBeforeClose(b *gocef.Browser) {
	gocef.QuitMessageLoop()
}

func (c Client) GetLifeSpanHandler() gocef.LifeSpanHandler {
	return LifeSpanHandler{}
}

func main() {
	if retcode := gocef.ExecuteProcess(); retcode >= 0 {
		os.Exit(retcode)
	}
	gocef.Initialize()
	client := Client{}
	gocef.CreateBrowser(client)
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
