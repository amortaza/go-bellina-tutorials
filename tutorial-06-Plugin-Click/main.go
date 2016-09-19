package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-bellina-plugins/click"
)

func initialize() {
	fmt.Println("Initializing!")

	go_dark_ux.Init()
}

func uninitialize() {
	fmt.Println("Uninitializing!")
	go_dark_ux.Uninit()
}

func tick() {

	bl.Root()
	{
		bl.Dim( bl.Window_Width, bl.Window_Height)

		border.Fill(50, 0, 0)
		border.Wire()

		bl.Div()
		{
			bl.Id("one")
		    	bl.Pos(100,100)
		    	bl.Dim(200,200)
		    	border.Fill(0,50,0)

		    	click.On(func(e interface{}) {
		        	clickEvent := e.(click.Event)

		    		fmt.Println("Clicked on", clickEvent.Target.Id, clickEvent.X, clickEvent.Y)
		    	})
		}
		bl.End()
	}
	bl.End()
}


func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.New(), 640, 480, "Bellina Tutorial 03 - MouseMove", initialize, tick, uninitialize )

	fmt.Println("bye!")
}



