package text32x36

// Width returns the number of columns.
//
// Note that this is NOT the width of the image.Image that this represents.
// For image.Image width information use the Bounds method.
func (receiver Slice) Width() int {
	return Width
}

// Height returns the number of rows.
//
// Note that this is NOT the height of the image.Image that this represents.
// For image.Image height information use the Bounds method.
func (receiver Slice) Height() int {
	return Height
}
