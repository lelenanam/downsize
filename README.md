# downsize

[![Build Status](https://travis-ci.org/lelenanam/downsize.svg?branch=master)](https://travis-ci.org/lelenanam/downsize)
[![GoDoc](https://godoc.org/github.com/lelenanam/downsize?status.svg)](https://godoc.org/github.com/lelenanam/downsize)

Reduces an image to a specified file size in bytes.

# Installation

```bash
$ go get -u github.com/lelenanam/downsize
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

# Sample

The original image `2.4 MB`:

![flower](https://cloud.githubusercontent.com/assets/4003503/24270582/f352a102-0fd2-11e7-852e-7ea77c4eae82.jpg)

Downsize to `200 KB`, `png` format for result image:

```sh
$ downsize -s=204800 -f=png flower.jpg flower200kb.png
```

Resized result `200 KB`:

![flower200kb](https://cloud.githubusercontent.com/assets/4003503/24270862/1126aace-0fd4-11e7-8c06-769162a93abe.png)

Downsize to `100 KB`, `png` format for result image:

```sh
$ downsize -s=102400 -f=png flower.jpg flower100kb.png
```

Resized result `100 KB`:

![flower100kb](https://cloud.githubusercontent.com/assets/4003503/24270871/1b7d8e7a-0fd4-11e7-8b27-8b055a60201b.png)


The original image `3.4 MB`:

![leaves](https://cloud.githubusercontent.com/assets/4003503/24270590/ffc8b070-0fd2-11e7-949f-3f76364ac252.jpg)

Downsize to `200 KB`, auto determine format for result image:

```sh
$ downsize -s=204800 leaves.jpg leaves200kb.jpg
```

Resized result `200 KB`:

![leaves200kb](https://cloud.githubusercontent.com/assets/4003503/24270881/245cb76e-0fd4-11e7-86a4-b3547010e4f6.jpg)

Downsize to `100 KB`, auto determine format for result image:

```sh
$ downsize -s=102400 leaves.jpg leaves100kb.jpg
```

Resized result `100 KB`:

![leaves100kb](https://cloud.githubusercontent.com/assets/4003503/24271855/02c5bcfa-0fd8-11e7-8bbc-b1cf86751350.jpg)

# License

[MIT License](LICENSE.md)
