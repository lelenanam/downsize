package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/lelenanam/downsize"
)

// USAGE: downsize [-s=size] [-f=format] [-q=jpeg quality] [-i=infile] [-o=outfile]
// USAGE: downsize [--help]

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s [-s=size] [-f=format] [-q=jpeg quality] [-i=infile] [-o=outfile]\n", os.Args[0])
		flag.PrintDefaults()
	}
	var size = flag.Int("s", 204800, "desired output file size in bytes")
	var format = flag.String("f", "", "format: jpeg, png or gif, by default the format of an image is determined during decoding")
	var quality = flag.Int("q", downsize.DefaultQuality, "desired output jpeg quality, ranges from 1 to 100 inclusive, higher is better")
	var infile = flag.String("i", "", "input file name, required")
	var outfile = flag.String("o", "", "output file name, required")
	flag.Parse()

	if *infile == "" || *outfile == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(*infile)
	if err != nil {
		log.Fatalln("Cannot open input file: ", *infile, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Cannot close input file: ", *infile, err)
		}
	}()

	img, decodedFormat, err := image.Decode(file)
	if err != nil {
		log.Fatalln("Cannot decode file: ", file.Name(), err)
	}

	out, err := os.Create(*outfile)
	if err != nil {
		log.Fatalln("Cannot create output file: ", *outfile, err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Println("Cannot close output file: ", *outfile, err)
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
