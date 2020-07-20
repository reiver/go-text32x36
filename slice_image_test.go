package text32x36_test

import (
	"github.com/reiver/go-text32x36"

	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"strings"

	"testing"
)

func TestImageImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = text32x36.Slice(nil)

	if nil == datum {
		t.Errorf("This should never happen.")
	}
}

func TestSlice_At(t *testing.T) {

	var buffer [text32x36.ByteSize]uint8

	var text text32x36.Slice = text32x36.Slice(buffer[:])

	{
		runes := text.Runes()

		copy(runes, []rune(
			//           1         2         3
			// 01234567890123456789012345678901
			  "Hello world!                    "+
			  "                                "+
			  "BEWARE I LIVE!                  "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "                                "+
			  "           GAME OVER            ",
		))
	}

	var serialized strings.Builder
	{
		var pngBuffer bytes.Buffer

		err := png.Encode(&pngBuffer, text)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

                encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())

		serialized.WriteString("IMAGE:")
                serialized.WriteString(encoded)
	}

	{
		const expected = "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAAAAAB+tfKMAAAEAUlEQVR4nOzd7XKaQABG4V3H+7/l7fgB6wK2Ma15sLxn2jERwhwOoMCP5dxKLZf/M63cfptev0qfvy/hpQUYTqt36uL1q9TlT+36b++c55/ak5VevH/drLfdptx/fHVf2RWnUtp1O7VS69YW23h/mL/ONaap7f5L/Ygu5zLrf3V/bcs56zh1s+Nu6YfA32yu7T/9gO1fLgdwnY7rm/Kz14l2+5v+GdC/Nab5bzN8xrfA4Vl/DR6MBNACmgTQApoE0AKawwc4387b5/P3Ol/tlTafCD5On/9ydZU4nhG28TJpt2eGp8s12/1kuNbxlHdal+3pi4u9+deN92stu70/MBwCbdjA68vj727CXW76O6frFXzfecvGdfFy+qvs+v7AJcC0qedbWff7A9MKv3CTbLPRDle7MxwCY4h+xNflR+Ez2rCsec+flrvHEudWphtbfVetj3e4+uvqxlgbvwXa8P58U+BhEXsscHQOfyKUAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkCTAFpAkwBaQJMAWkBz+ADn9y26DcMJt/5ImrIakXhjHOK2nO1NvG8PmMblfRhrsw3jDPf5xrFIb9NrefLYm3/M2wK0PiTz6jE2i/GLV4O1/uTgo288BJ7yp836MI7rJx8C8zOX+oOX5lF2f893H/b1Ld4WoD/Cra5W6HGXr9fPhro1DvGPPLLpfYdAbZv2q/GJ633FF+MQ73Ac9v+Tw58IJYAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQJIAW0CSAFtAkgBbQHD7ArwAAAP//sFKpVXrbX/MAAAAASUVORK5CYII="

		actual := serialized.String()

		if expected != actual {
			t.Errorf("The actual image is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
