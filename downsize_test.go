package downsize

import (
	"bytes"
	"image"
	"os"
	"testing"
)

func TestDownsize(t *testing.T) {
	table := []struct {
		img     string
		maxSize int
	}{
		{
			img:     "./testdata/cat.jpg",
			maxSize: 20000,
		},
		{
			img:     "./testdata/grape.png",
			maxSize: 20000,
		},
		{
			img:     "./testdata/fry.gif",
			maxSize: 20000,
		},
	}
	for _, test := range table {
		file, err := os.Open(test.img)
		if err != nil {
			t.Errorf("Error: %v, cannot open file %v\n", err, test.img)
		}
		defer file.Close()

		resBuf := bytes.NewBuffer(nil)
		img, format, err := image.Decode(file)
		if err != nil {
			t.Errorf("Error: %v, cannot decode file %v\n", err, test.img)
		}

		if err = Encode(resBuf, img, &Options{Size: test.maxSize, Format: format}); err != nil {
			t.Errorf("Error: %v, cannot downsize file %v\n", err, test.img)
		}
		resSize := resBuf.Len()
		resAccur := 1 - float64(resSize)/float64(test.maxSize)
		if resAccur > Accuracy {
			t.Errorf("[FAIL] For file: %v, size: %v, accuracy: %.4f, should be: %v\n",
				test.img, resSize, resAccur, Accuracy)
		} else {
			t.Logf("[OK] File: %v, size: %v, accuracy: %.4f\n", test.img, resSize, resAccur)
		}
	}
}
