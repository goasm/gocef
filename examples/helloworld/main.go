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
	appSettings := &gocef.Settings{
		NoSandbox: true,
	}
	gocef.Initialize(appSettings)
	info := &gocef.WindowInfo{
		WindowName: "www.baidu.com",
		Width:      500,
		Height:     300,
		X:          710,
		Y:          390,
	}
	client := gocef.NewClient(&myClient{})
	url := "https://www.baidu.com"
	settings := &gocef.BrowserSettings{}
	gocef.CreateBrowser(info, client, url, settings)
	gocef.RunMessageLoop()
	gocef.Shutdown()
	client.Destroy()
}
