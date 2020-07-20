package text32x36

import (
	"unsafe"
)

const (
        Width  = 32
        Height = 36
	Depth  = int(unsafe.Sizeof(rune(0)))
)
