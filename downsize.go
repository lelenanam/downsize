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

//Accuracy for calculating specified file size maxSize
//for maxSize result might be in range [maxSize-Accuracy; maxSize]
const Accuracy = 0.05

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

//Downsize reads image from reader r, changes size to make it <=maxSize and writes result to writer w
func Downsize(maxSize int, r io.Reader, w io.Writer) error {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	originSize, err := io.Copy(buf, r)
	if err != nil {
		return err
	}

	if int(originSize) <= maxSize {
		if _, err := io.Copy(w, buf); err != nil {
			return err
		}
		return nil
	}

	img, format, err := image.Decode(buf)
	if err != nil {
		return err
	}

	buf.Reset()
	min := 0
	max := img.Bounds().Dx()
	//Removing metadata by resizing with the same width
	if err := changeWidht(max, img, format, buf); err != nil {
		return err
	}

	noMetadataSize := buf.Len()
	if noMetadataSize <= maxSize {
		if _, err := io.Copy(w, buf); err != nil {
			return err
		}
		return nil
	}

	for min < max {
		buf.Reset()
		newWidth := (min + max) / 2
		if err := changeWidht(newWidth, img, format, buf); err != nil {
			return err
		}
		newSize := buf.Len()
		if newSize > maxSize {
			max = newWidth - 1
		} else {
			newAccur := 1 - float64(newSize)/float64(maxSize)
			if newAccur <= Accuracy {
				break
			}
			min = newWidth + 1
		}
	}
	_, err = io.Copy(w, buf)
	if err != nil {
		return err
	}
	return nil
}

func changeWidht(width int, img image.Image, format string, w io.Writer) error {
	m := resize.Resize(uint(width), 0, img, resize.Lanczos3)
	switch format {
	case "jpeg":
		jpeg.Encode(w, m, &jpeg.Options{Quality: 98})
	case "png":
		png.Encode(w, m)
	case "gif":
		gif.Encode(w, m, nil)
	default:
		return fmt.Errorf("Unknown image format %q", format)
	}
	return nil
}
