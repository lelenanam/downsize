# downsize

[![Build Status](https://travis-ci.org/lelenanam/downsize.svg?branch=master)](https://travis-ci.org/lelenanam/downsize)
[![GoDoc](https://godoc.org/github.com/lelenanam/downsize?status.svg)](https://godoc.org/github.com/lelenanam/downsize)

Reduces an image to a specified file size in bytes.
Also [command line tool](https://github.com/lelenanam/downsize/tree/master/cmd/downsize) available.

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

By default an image encodes with `jpeg` format and with the quality `DefaultQuality = 80`.
All metadata is stripped after encoding.

```go
var DefaultQuality = 80
var defaultFormat = "jpeg"
var defaultJpegOptions = &jpeg.Options{Quality: DefaultQuality}
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
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Cannot close input file: ", err)
		}
	}()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error: %v, cannot decode file %v", err, file.Name())
	}

	out, err := os.Create("resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Println("Cannot close output file: ", err)
		}
	}()

	opt := &downsize.Options{Size: 1048576, Format: format}
	if err = downsize.Encode(out, img, opt); err != nil {
		log.Fatalf("Error: %v, cannot downsize image to size: %v", err, opt.Size)
	}
}
```

# Sample

The original jpeg image `2.4 MB`:

![flower](https://cloud.githubusercontent.com/assets/4003503/24624009/4fcd3962-185f-11e7-8b6b-a28e217cba27.jpg)

Downsize to `200 KB`, `png` format and default quality for result image:

```go
opt := &downsize.Options{Size: 204800, Format: "png"}
err = downsize.Encode(out, img, opt)
```

Resized result `200 KB`:

![flower200kbpng](https://cloud.githubusercontent.com/assets/4003503/24624123/aa7f5f16-185f-11e7-9340-e896ee116bc3.png)

Downsize to `200 KB`, `jpeg` format and default quality for result image:

```go
opt := &downsize.Options{Size: 204800, Format: "jpeg"}
err = downsize.Encode(out, img, opt)
```

Resized result `200 KB`:

![flower200kbjpegq80](https://cloud.githubusercontent.com/assets/4003503/24624188/de20d7b4-185f-11e7-931b-1b2eeb1ab0f0.jpg)

Downsize to `200 KB`, `jpeg` format and quality `50` for result image:

```go
opt := &downsize.Options{Size: 204800, Format: "jpeg", JpegOptions: &jpeg.Options{Quality: 50}}
err = downsize.Encode(out, img, opt)
```

Resized result `200 KB`, quality `50`:

![flower200kbjpegq50](https://cloud.githubusercontent.com/assets/4003503/24624303/3edbcfbe-1860-11e7-947f-16954fd3a872.jpg)


The original image `3.4 MB`:

![leaves](https://cloud.githubusercontent.com/assets/4003503/24270590/ffc8b070-0fd2-11e7-949f-3f76364ac252.jpg)

Downsize to `100 KB`, auto determine format and default quality for result image:

```go
opt := &downsize.Options{Size: 102400}
err = downsize.Encode(out, img, opt)
```

Resized result `100 KB`:

![leaves100kb](https://cloud.githubusercontent.com/assets/4003503/24624461/c86e946e-1860-11e7-8059-c4bb25ad3c49.jpg)

Downsize to `100 KB`, auto determine format and quality `50` for result image:

```go
opt := &downsize.Options{Size: 102400, JpegOptions: &jpeg.Options{Quality: 50}}
err = downsize.Encode(out, img, opt)
```

Resized result `100 KB`, quality `50`:

![leaves100kbjpegq50](https://cloud.githubusercontent.com/assets/4003503/24624590/38ccf520-1861-11e7-964e-7b3411a3fc11.jpg)

Downsize to `50 KB`, auto determine format and default duality for result image:

```go
opt := &downsize.Options{Size: 51200}
err = downsize.Encode(out, img, opt)
```

Resized result `50 KB`:

![leaves50kbjpegq80](https://cloud.githubusercontent.com/assets/4003503/24624690/7b46c0ac-1861-11e7-93d0-b4c87b9765eb.jpg)

# License

[MIT License](LICENSE.md)
