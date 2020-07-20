package text32x36

import (
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// Slice represents a text matrix.
type Slice []uint8

// Clear sets all elements of the text matrix to the space character.
func (receiver Slice) Clear() {
	runes := receiver.Runes()
	if nil == runes {
		return
	}

	for i:=0; i<len(runes); i++ {
		runes[i] = ' '
	}
}

// LineFeed moved every line up one row.
func (receiver Slice) LineFeed() {
	runes := receiver.Runes()
	if nil == runes {
		return
	}

	for y:=1; y<=Height; y++ {
		for x:=0; x<Width; x++ {
			srcOffset := receiver.runesOffset(x,y)

			dstOffset := receiver.runesOffset(x,y-1)

			runes[dstOffset] = runes[srcOffset]
		}
	}

	{
		y := Height-1

		for x:=0; x<Width; x++ {
			offset := receiver.runesOffset(x,y)

			runes[offset] = ' '
		}
	}
}

func (receiver Slice) Publish(s string) {
	runes := receiver.Runes()
	if nil == runes {
		return
	}

	if "" == s {
		receiver.LineFeed()
		return
	}

	p := s

	Loop: for 0 < len(p) {
		receiver.LineFeed()

		y := Height-1
		for x:=0; x<Width; x++ {
			offset := receiver.runesOffset(x,y)

			r, size := utf8.DecodeRuneInString(p)
			p = p[size:]
			if utf8.RuneError == r  && 0 < size {
				continue
			}
			if utf8.RuneError == r {
				break Loop
			}

			runes[offset] = r
		}

	}
}

// Runes returns the data in the text matrix in []rune form.
//
// Where the top-left corner is offset 0. And the botom-right
// corner is offset  1,151 ( = 32×36 - 1).
func (receiver Slice) Runes() []rune {
	if nil == receiver {
		return nil
	}
	if 0 >= len(receiver) {
		return nil
	}
	if ByteSize != len(receiver) {
		return nil
	}

	h := (*reflect.SliceHeader)(unsafe.Pointer(&receiver))

	var header reflect.SliceHeader
	header.Data = h.Data
	header.Len  = h.Len/4
	header.Cap  = h.Cap/4

	return *(*[]rune)(unsafe.Pointer(&header))
}

// String returns contents of the textmatrix  as a string.
func (receiver Slice) String() string {
	runes := receiver.Runes()

	if nil == runes {
		return  ""+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "+"\n"+
			"                                "
	}

	var buffer strings.Builder

	for offset:=0; offset<len(runes); offset++ {
		if x := offset % Width; 0 == x && 0 != offset {
			buffer.WriteRune('\n')
		}

		buffer.WriteRune(runes[offset])
	}

	return buffer.String()
}

func (receiver Slice) runesOffset(x int, y int) int {
	x = x % Width
	if 0 > x {
		x += Width
	}

	y = y % Height
	if 0 > y {
		y += Height
	}

	return y*Width + x
}
