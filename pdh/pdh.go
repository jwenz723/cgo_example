package pdh

// #cgo LDFLAGS: -L./lib -lpdh
// #include <pdh.h>
import "C"

func PdhValidatePath(path string) uint32 {
	ret := C.PdhValidatePath(C.CString(path))
	return uint32(ret)
}