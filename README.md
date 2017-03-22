# downsize

[![Build Status](https://travis-ci.org/lelenanam/downsize.svg?branch=master)](https://travis-ci.org/lelenanam/downsize)
[![GoDoc](https://godoc.org/github.com/lelenanam/downsize?status.svg)](https://godoc.org/github.com/lelenanam/downsize)

Reduces an image to a specified file size in bytes.

# Installation

```bash
$ go get github.com/lelenanam/downsize
```

# Usage

```go
import "github.com/lelenanam/downsize"
```

The `downsize` package provides a function `downsize.Encode`:

```go
func Encode(w io.Writer, m image.Image, o *Options) error 
```

This function:

* takes any image type that implements `image.Image` interface as an input `m`
* reduces an image's dimensions to achieve a specified file size `Options.Size` in bytes
* writes result Image `m` to writer `w` with the given options
* default parameters are used if a `nil` `*Options` is passed

```go
// Options are the encoding parameters.
type Options struct {
	// Size is desired output file size in bytes
	Size int
	// Format is image format to encode
	Format string
	// JpegOptions are the options for jpeg format
	JpegOptions *jpeg.Options
	// GifOptions are the options for gif format
	GifOptions *gif.Options
}
```

By default an image encodes with `jpeg` format and with the highest quality `&jpeg.Options{Quality: 100}`.
All metadata is stripped after encoding.

```go
var defaultFormat = "jpeg"
var defaultJpegOptions = &jpeg.Options{Quality: 100}
var defaultOptions = &Options{Format: defaultFormat, JpegOptions: defaultJpegOptions}
```

# Example

```go
package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/lelenanam/downsize"
)

func main() {
	file, err := os.Open("img.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error: %v, cannot decode file %v", err, file.Name())
	}

	out, err := os.Create("resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	opt := &downsize.Options{Size: 1048576, Format: format}
	if err = downsize.Encode(out, img, opt); err != nil {
		log.Fatalf("Error: %v, cannot downsize image to size: %v", err, opt.Size)
	}
}
```

# License

[MIT License](LICENSE.md)
