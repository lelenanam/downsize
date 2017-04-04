package main

import (
	"flag"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/lelenanam/downsize"
)

// USAGE: downsize [-s=size] [-f=format] [-q=jpeg quality] infile outfile
// USAGE: downsize [--help]

func main() {
	var size = flag.Int("s", 204800, "desired output file size in bytes")
	var format = flag.String("f", "", "format: jpeg, png or gif, by default the format of an image is determined during decoding")
	var quality = flag.Int("q", downsize.DefaultQuality, "desired output jpeg quality, ranges from 1 to 100 inclusive, higher is better")
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal("Cannot open input file; ", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Cannot close input file: ", err)
		}
	}()

	img, decodedFormat, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Cannot decode file %v: %q", file.Name(), err)
	}

	out, err := os.Create(flag.Arg(1))
	if err != nil {
		log.Fatal("Cannot create output file: ", err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Println("Cannot close output file: ", err)
		}
	}()

	outFormat := *format
	if *format == "" {
		outFormat = decodedFormat
		log.Println("Output format:", decodedFormat)
	}
	if *format == "jpg" {
		outFormat = "jpeg"
	}

	opt := &downsize.Options{Size: *size, Format: outFormat}

	if *quality != 0 && outFormat == "jpeg" {
		opt.JpegOptions = &jpeg.Options{Quality: *quality}
	}

	if err = downsize.Encode(out, img, opt); err != nil {
		log.Fatalf("Cannot downsize image to size: %v, error: %q", opt.Size, err)
	}
}
