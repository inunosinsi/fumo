package main

import (	
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	editor := theme.CreateCodeEditor()
	
	window := theme.CreateWindow(580, 300, "Code Editor")
	
	color := gxui.White
	
	brush := gxui.CreateBrush(color)
	window.SetBackgroundBrush(brush)
	
	window.AddChild(editor)
	
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}