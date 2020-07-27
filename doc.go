/*
Package text32x36 provides a type for dealing with a 32×36 text matrix,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

A text matrix is a 2-dimensional matrix of UNICODE characters.

This can be used to create the output end of command-line interface (CLI).
Other labels some might associate with this are: terminal, console, and even shell.

Example

Here is an example:

	import "github.com/reiver/go-text32x36"
	
	// ...
	
	var memory [text32x36.ByteSize]uint8
	
	var terminal text32x36.Slice = text32x36.Slice(memory[:])
	
	// ...
	
	terminal.Publish("Hello world!")
	terminal.LineFeed()
	terminal.Publish("Goodbye · Khodafez · 안녕 · 再見")
	
	// ...
	
	draw.Draw(dst, terminal.Bounds, terminal, terminal.Bounds().Min, draw.Over)
*/
package text32x36
