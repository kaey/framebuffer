// Copyright 2013 Konstantin Kulikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package framebuffer is an interface to linux framebuffer device.
package framebuffer

import (
	"os"
	"syscall"
	"unsafe"
)

// Framebuffer contains information about framebuffer.
type Framebuffer struct {
	dev      *os.File
	finfo    fixedScreenInfo
	vinfo    variableScreenInfo
	data     []byte
	restData []byte
}

// Init opens framebuffer device, maps it to memory and saves its current contents.
func Init(dev string) (*Framebuffer, error) {
	var (
		fb  = new(Framebuffer)
		err error
	)

	fb.dev, err = os.OpenFile(dev, os.O_RDWR, os.ModeDevice)
	if err != nil {
		return nil, err
	}

	err = ioctl(fb.dev.Fd(), getFixedScreenInfo, unsafe.Pointer(&fb.finfo))
	if err != nil {
		fb.dev.Close()
		return nil, err
	}

	err = ioctl(fb.dev.Fd(), getVariableScreenInfo, unsafe.Pointer(&fb.vinfo))
	if err != nil {
		fb.dev.Close()
		return nil, err
	}

	fb.data, err = syscall.Mmap(int(fb.dev.Fd()), 0, int(fb.finfo.Smem_len+uint32(fb.finfo.Smem_start&uint64(syscall.Getpagesize()-1))), protocolRead|protocolWrite, mapShared)
	if err != nil {
		fb.dev.Close()
		return nil, err
	}

	fb.restData = make([]byte, len(fb.data))
	for i := range fb.data {
		fb.restData[i] = fb.data[i]
	}

	return fb, nil
}

// Close closes framebuffer device and restores its contents.
func (fb *Framebuffer) Close() {
	for i := range fb.restData {
		fb.data[i] = fb.restData[i]
	}
	syscall.Munmap(fb.data)
	fb.dev.Close()
}

// WritePixel changes pixel at x, y to specified color.
func (fb *Framebuffer) WritePixel(x, y, red, green, blue, alpha int) {
	offset := (int(fb.vinfo.Xoffset)+x)*(int(fb.vinfo.Bits_per_pixel)/8) + (int(fb.vinfo.Yoffset)+y)*int(fb.finfo.Line_length)
	fb.data[offset] = byte(blue)
	fb.data[offset+1] = byte(green)
	fb.data[offset+2] = byte(red)
	fb.data[offset+3] = byte(alpha)
}

// Clear fills screen with specified color
func (fb *Framebuffer) Clear(red, green, blue, alpha int) {
	w, h := fb.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fb.WritePixel(i, j, red, green, blue, alpha)
		}
	}
}

// Size returns dimensions of a framebuffer.
func (fb *Framebuffer) Size() (width, height int) {
	return int(fb.vinfo.Xres), int(fb.vinfo.Yres)
}

func ioctl(fd uintptr, cmd uintptr, data unsafe.Pointer) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, uintptr(data))
	if errno != 0 {
		return os.NewSyscallError("IOCTL", errno)
	}
	return nil
}
