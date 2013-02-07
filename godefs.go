// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _godefs-source.go

package fb

type FixedScreenInfo struct {
	Id           [16]int8
	Smem_start   uint64
	Smem_len     uint32
	Type         uint32
	Type_aux     uint32
	Visual       uint32
	Xpanstep     uint16
	Ypanstep     uint16
	Ywrapstep    uint16
	Pad_cgo_0    [2]byte
	Line_length  uint32
	Pad_cgo_1    [4]byte
	Mmio_start   uint64
	Mmio_len     uint32
	Accel        uint32
	Capabilities uint16
	Reserved     [2]uint16
	Pad_cgo_2    [2]byte
}
type VariableScreenInfo struct {
	Xres           uint32
	Yres           uint32
	Xres_virtual   uint32
	Yres_virtual   uint32
	Xoffset        uint32
	Yoffset        uint32
	Bits_per_pixel uint32
	Grayscale      uint32
	Red            BitField
	Green          BitField
	Blue           BitField
	Transp         BitField
	Nonstd         uint32
	Activate       uint32
	Height         uint32
	Width          uint32
	Accel_flags    uint32
	Pixclock       uint32
	Left_margin    uint32
	Right_margin   uint32
	Upper_margin   uint32
	Lower_margin   uint32
	Hsync_len      uint32
	Vsync_len      uint32
	Sync           uint32
	Vmode          uint32
	Rotate         uint32
	Colorspace     uint32
	Reserved       [4]uint32
}
type BitField struct {
	Offset uint32
	Length uint32
	Right  uint32
}

const (
	GetFixedScreenInfo    uintptr = 0x4602
	GetVariableScreenInfo uintptr = 0x4600
)

const (
	ProtocolRead  int = 0x1
	ProtocolWrite int = 0x2
	MapShared     int = 0x1
)
