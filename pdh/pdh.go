package pdh

// #cgo LDFLAGS: -L./lib -lpdh
// #include <pdh.h>
import "C"
import (
	"unsafe"
)

func PdhValidatePath(path string) uint32 {
	cPath := C.CString(path)
	ret := C.PdhValidatePath(cPath)
	defer C.free(unsafe.Pointer(cPath))
	return uint32(ret)
}