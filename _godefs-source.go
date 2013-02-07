package fb

/*
#include <linux/fb.h>
#include <sys/mman.h>
*/
import "C"

type FixedScreenInfo C.struct_fb_fix_screeninfo
type VariableScreenInfo C.struct_fb_var_screeninfo
type BitField C.struct_fb_bitfield

const (
	GetFixedScreenInfo    uintptr = C.FBIOGET_FSCREENINFO
	GetVariableScreenInfo uintptr = C.FBIOGET_VSCREENINFO
)

const (
	ProtocolRead  int = C.PROT_READ
	ProtocolWrite int = C.PROT_WRITE
	MapShared     int = C.MAP_SHARED
)
