package main

import (
	"github.com/Kaey/framebuffer"
	"fmt"
)

func main() {
	fb.Init()
	defer fb.Close()
	fb.Clear()
	fb.Write(1679, 1049, 255, 0, 0, 0)
	fmt.Scanln()
}
