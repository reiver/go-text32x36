# go-text32x36

Package text32x36 provides a type for dealing with a 32×36 text matrix,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

A text matrix is a 2-dimensional matrix of UNICODE characters.

This can be used to create the output end of command-line interface (CLI).
Other labels some might associate with this are: terminal, console, and even shell.

The image.Image that this represents has 8x8 pixel characers.
Which means the image.Image has a width×height of 256×288.

Example

Here is an example:
```go
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
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-text32x36

[![GoDoc](https://godoc.org/github.com/reiver/go-text32x36?status.svg)](https://godoc.org/github.com/reiver/go-text32x36)
