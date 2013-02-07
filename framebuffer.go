// Interface to linux framebuffer device.
package fb

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	fbdev   *os.File
	finfo   FixedScreenInfo
	vinfo   VariableScreenInfo
	fbdata  []byte
	fbdata2 []byte
)

func Write(x, y, red, green, blue, alpha int) {
	offset := (int(vinfo.Xoffset)+x)*(int(vinfo.Bits_per_pixel)/8) + (int(vinfo.Yoffset)+y)*int(finfo.Line_length)
	fbdata[offset] = byte(blue)
	fbdata[offset+1] = byte(green)
	fbdata[offset+2] = byte(red)
	fbdata[offset+3] = byte(alpha)
}

func Init() {
	fbdev, _ = os.OpenFile("/dev/fb0", os.O_RDWR, os.ModeDevice)
	syscall.Syscall(syscall.SYS_IOCTL, fbdev.Fd(), GetFixedScreenInfo, uintptr(unsafe.Pointer(&finfo)))
	syscall.Syscall(syscall.SYS_IOCTL, fbdev.Fd(), GetVariableScreenInfo, uintptr(unsafe.Pointer(&vinfo)))
	fbdata, _ = syscall.Mmap(int(fbdev.Fd()), 0, int(finfo.Smem_len+uint32(finfo.Smem_start&4095)), ProtocolRead|ProtocolWrite, MapShared)
	fbdata2 = make([]byte, len(fbdata))
	for i := range fbdata {
		fbdata2[i] = fbdata[i]
	}
}

func Close() {
	for i := range fbdata2 {
		fbdata[i] = fbdata2[i]
	}
	fbdev.Close()
	syscall.Munmap(fbdata)
}

func Clear() {
	w, h := Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			Write(i, j, 0, 0, 0, 0)
		}
	}
}

func Size() (width, height int) {
	return int(vinfo.Xres), int(vinfo.Yres)
}
