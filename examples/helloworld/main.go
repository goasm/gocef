package main

import (
	"os"

	"github.com/goasm/gocef"
)

type myClient struct{}
type myLifeSpanHandler struct{}

func (h *myLifeSpanHandler) OnBeforeClose(b *gocef.Browser) {
	gocef.QuitMessageLoop()
}

func (c *myClient) GetLifeSpanHandler() gocef.LifeSpanHandler {
	return &myLifeSpanHandler{}
}

func main() {
	if retcode := gocef.ExecuteProcess(); retcode >= 0 {
		os.Exit(retcode)
	}
	settings := &gocef.Settings{}
	settings.NoSandbox = true
	gocef.Initialize(settings)
	client := gocef.NewClient(&myClient{})
	gocef.CreateBrowser(client)
	gocef.RunMessageLoop()
	gocef.Shutdown()
}
