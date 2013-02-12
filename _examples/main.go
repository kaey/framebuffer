package main

import (
	"fmt"
	"github.com/Kaey/framebuffer"
	"log"
)

func main() {
	fb, err := framebuffer.Init("/dev/fb0")
	if err != nil {
		log.Fatalln(err)
	}
	defer fb.Close()
	fb.Clear(0, 0, 0, 0)
	fb.WritePixel(200, 100, 255, 0, 0, 0)
	fmt.Scanln()
}
