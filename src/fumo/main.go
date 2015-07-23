package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"os"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type FileIO struct {
	qml.Object
	text string
}

func (io *FileIO) EchoText() string{
	return "test"
}

func run() error {
	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(f *FileIO, obj qml.Object) { f.Object = obj },
	}})
	
	engine := qml.NewEngine()
	
	component, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}
	win := component.CreateWindow(nil)
	win.Show()
	win.Wait()
	
	return nil
}

