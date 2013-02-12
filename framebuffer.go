// Interface to linux framebuffer device.
package framebuffer

import (
	"os"
	"syscall"
	"unsafe"
)

type Framebuffer struct {
	dev      *os.File
	finfo    fixedScreenInfo
	vinfo    variableScreenInfo
	data     []byte
	restData []byte
}

func Init(dev string) (*Framebuffer, error) {
	var (
		fb    = new(Framebuffer)
		err   error
		errno syscall.Errno
	)

	fb.dev, err = os.OpenFile(dev, os.O_RDWR, os.ModeDevice)
	if err != nil {
		return nil, err
	}

	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, fb.dev.Fd(), getFixedScreenInfo, uintptr(unsafe.Pointer(&fb.finfo)))
	if errno != 0 {
		return nil, errno
	}

	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, fb.dev.Fd(), getVariableScreenInfo, uintptr(unsafe.Pointer(&fb.vinfo)))
	if errno != 0 {
		return nil, errno
	}

	fb.data, err = syscall.Mmap(int(fb.dev.Fd()), 0, int(fb.finfo.Smem_len+uint32(fb.finfo.Smem_start&uint64(syscall.Getpagesize()-1))), protocolRead|protocolWrite, mapShared)
	if err != nil {
		return nil, err
	}

	fb.restData = make([]byte, len(fb.data))
	for i := range fb.data {
		fb.restData[i] = fb.data[i]
	}

	return fb, nil
}

func (fb *Framebuffer) WritePixel(x, y, red, green, blue, alpha int) {
	offset := (int(fb.vinfo.Xoffset)+x)*(int(fb.vinfo.Bits_per_pixel)/8) + (int(fb.vinfo.Yoffset)+y)*int(fb.finfo.Line_length)
	fb.data[offset] = byte(blue)
	fb.data[offset+1] = byte(green)
	fb.data[offset+2] = byte(red)
	fb.data[offset+3] = byte(alpha)
}

func (fb *Framebuffer) Close() {
	for i := range fb.restData {
		fb.data[i] = fb.restData[i]
	}
	syscall.Munmap(fb.data)
	fb.dev.Close()
}

func (fb *Framebuffer) Clear(red, green, blue, alpha int) {
	w, h := fb.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fb.WritePixel(i, j, red, green, blue, alpha)
		}
	}
}

func (fb *Framebuffer) Size() (width, height int) {
	return int(fb.vinfo.Xres), int(fb.vinfo.Yres)
}
