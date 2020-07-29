package text32x36

// ByteSize represents the number of bytes of the text matrix — i.e., how many uint8 are in the text matrix.
//
// This is used in this way:
//
//	import "github.com/reiver/go-text32x36"
//	
//	// ...
//	
//	var memory [text32x36.ByteSize]uint8 // <--------- ”text32x36.ByteSize” is used here.
//	
//	var terminal text32x36.Slice = text32x36.Slice(memory[:])
const ByteSize = Width * Height * Depth
