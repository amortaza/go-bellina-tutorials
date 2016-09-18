package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux"
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
			bl.Id("child")
			bl.Pos(20, 40)
			bl.Dim(130, 160)

			bl.OnMouseMove(func(e *bl.MouseMoveEvent) {
				fmt.Println("Mouse moved on ", e.Target.Id, "at", e.X, e.Y)
			})

			border.Fill(0, 50, 0)
			border.Wire()
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



