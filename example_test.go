// Copyright 2013 Konstantin Kulikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package framebuffer_test

import (
	"fmt"
	"log"

	"github.com/kaey/framebuffer"
)

func Example() {
	fb, err := framebuffer.Init("/dev/fb0")
	if err != nil {
		log.Fatalln(err)
	}
	defer fb.Close()
	fb.Clear(0, 0, 0, 0)
	fb.WritePixel(200, 100, 255, 0, 0, 0)
	fmt.Scanln()
}
