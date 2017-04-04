package downsize

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"sync"

	"github.com/nfnt/resize"
)

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

// DefaultQuality is default quality to encode image
const DefaultQuality = 80

// defaultFormat is default output format
var defaultFormat = "jpeg"

// defaultJpegOptions is default options to encode jpeg format
var defaultJpegOptions = &jpeg.Options{Quality: DefaultQuality}

// defaultOptions is default options for downsize
var defaultOptions = &Options{Format: defaultFormat, JpegOptions: defaultJpegOptions}

//Accuracy for calculating specified file size Options.Size
//for Options.Size result might be in range [Options.Size - Options.Size*Accuracy; Options.Size]
const Accuracy = 0.05

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func setOptions(o *Options) *Options {
	opts := defaultOptions
	if o != nil {
		opts = o
	}
	if opts.Format == "" {
		opts.Format = defaultFormat
		if opts.GifOptions != nil {
			opts.Format = "gif"
		}
	}
	if opts.Format == defaultFormat && opts.JpegOptions == nil {
		opts.JpegOptions = defaultJpegOptions
	}
	return opts
}

// Encode changes size of Image img (result size<=o.Size)
// and writes the Image img to w with the given options.
// Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, img image.Image, o *Options) error {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	opts := setOptions(o)

	if err := encode(buf, img, opts); err != nil {
		return err
	}
	originSize := buf.Len()

	if opts.Size <= 0 {
		opts.Size = originSize
	}

	if originSize <= opts.Size {
		_, err := io.Copy(w, buf)
		return err
	}

	min := 0
	max := img.Bounds().Dx()

	for min < max {
		buf.Reset()
		newWidth := (min + max) / 2
		newImg := resize.Resize(uint(newWidth), 0, img, resize.Lanczos3)
		if err := encode(buf, newImg, opts); err != nil {
			return err
		}
		newSize := buf.Len()
		if newSize > opts.Size {
			max = newWidth - 1
		} else {
			newAccur := 1 - float64(newSize)/float64(opts.Size)
			if newAccur <= Accuracy {
				break
			}
			min = newWidth + 1
		}
	}
	_, err := io.Copy(w, buf)
	return err
}

func encode(w io.Writer, img image.Image, o *Options) error {
	switch o.Format {
	case "jpeg":
		return jpeg.Encode(w, img, o.JpegOptions)
	case "png":
		return png.Encode(w, img)
	case "gif":
		return gif.Encode(w, img, o.GifOptions)
	default:
		return fmt.Errorf("Unknown image format %q", o.Format)
	}
}
