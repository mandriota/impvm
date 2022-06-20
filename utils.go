package imp

import "unsafe"

func GetFloat64Ptr[T any](p *T) *float64 {
	return (*float64)(unsafe.Pointer(p))
}
