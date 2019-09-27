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
	info := &gocef.WindowInfo{}
	info.WindowName = "www.baidu.com"
	info.Width = 500
	info.Height = 300
	info.X = 710
	info.Y = 390
	client := gocef.NewClient(&myClient{})
	gocef.CreateBrowser(info, client)
	gocef.RunMessageLoop()
	gocef.Shutdown()
	client.Destroy()
}
