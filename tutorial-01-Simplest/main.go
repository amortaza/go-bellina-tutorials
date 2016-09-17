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
	}
	bl.End()
}


func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.New(), 1280, 1024, "Bellina v0.2", initialize, tick, uninitialize )

	fmt.Println("bye!")
}



