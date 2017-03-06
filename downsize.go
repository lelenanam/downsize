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

var defaultFormat = "jpeg"
var defaultJpegOptions = &jpeg.Options{Quality: 100}
var defaultOptions = &Options{Format: defaultFormat, JpegOptions: defaultJpegOptions}

//Accuracy for calculating specified file size Options.Size
//for Options.Size result might be in range [Options.Size - Options.Size*Accuracy; Options.Size]
const Accuracy = 0.05

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Encode changes size of Image m (result size<=o.Size)
// and writes the Image m to w with the given options.
// Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, m image.Image, o *Options) error {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

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

	if err := encode(buf, m, opts); err != nil {
		return err
	}
	originSize := buf.Len()

	if opts.Size <= 0 {
		opts.Size = originSize
	}

	if originSize <= opts.Size {
		if _, err := io.Copy(w, buf); err != nil {
			return err
		}
		return nil
	}

	buf.Reset()
	min := 0
	max := m.Bounds().Dx()

	for min < max {
		buf.Reset()
		newWidth := (min + max) / 2
		newm := resize.Resize(uint(newWidth), 0, m, resize.Lanczos3)
		if err := encode(buf, newm, opts); err != nil {
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
	if err != nil {
		return err
	}
	return nil
}

func encode(w io.Writer, m image.Image, o *Options) error {
	switch o.Format {
	case "jpeg":
		jpeg.Encode(w, m, o.JpegOptions)
	case "png":
		png.Encode(w, m)
	case "gif":
		gif.Encode(w, m, o.GifOptions)
	default:
		return fmt.Errorf("Unknown image format %q", o.Format)
	}
	return nil
}
