# downsize
Reduces an image to a specified file size.

Usage
-----

```go
import "github.com/morelena/downsize"
```

The downsize package provides a function `downsize.Downsize` which:

* reads image from reader `r`
* reduces an image's dimensions to achieve a specified file size `maxSize` in kilobytes
* writes result image to writer `w`
 
```go
downsize.Downsize(maxSize int, r io.Reader, w io.Writer) error 
```

Sample usage:

```go
package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/morelena/downsize"
)

func main() {
	file, err := os.Open("img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	out, err := os.Create("resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = downsize.Downsize(1048576, file, out)
	if err != nil {
		fmt.Println("Error while downsizing", err)
	}
}
```

License
-------

[MIT License](LICENSE.md)
