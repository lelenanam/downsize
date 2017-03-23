# downsize

[![Build Status](https://travis-ci.org/lelenanam/downsize.svg?branch=master)](https://travis-ci.org/lelenanam/downsize)
[![GoDoc](https://godoc.org/github.com/lelenanam/downsize?status.svg)](https://godoc.org/github.com/lelenanam/downsize)

Reduces an image to a specified file size in bytes.

# Installation

```bash
$ go get -u github.com/lelenanam/downsize/...
```

# Usage

## Command Line Usage

```sh
$ # You can specify the size in bytes and the format for the output file.
$ ds [-s size] [-f format] infile outfile
$ ds [-help]
```

By default, output file size is 200 KB and the format is determined during decoding.

## Example

Resize the file `image.jpg` to size = `1 MB` and save the result in `jpeg` format file `resized.jpg`.

```sh
$ downsize -s=1048576 -f=jpeg image.jpg resized.jpg
```

# Sample 1

The original image:

![flower](https://cloud.githubusercontent.com/assets/4003503/24270582/f352a102-0fd2-11e7-852e-7ea77c4eae82.jpg)

Downsize to `1 MB`, auto determine format for result image:

```sh
$ downsize -s=1048576 flower.jpg flower1mb.jpg
```

Resized result:

![flower1mb](https://cloud.githubusercontent.com/assets/4003503/24270847/031ddab0-0fd4-11e7-8c59-704ddeab9fe0.jpg)

Downsize to `200 KB`, `jpeg` format for result image:

```sh
$ downsize -s=204800 -f=jpeg flower.jpg flower200kb.jpg
```

Resized result:

![flower200kb](https://cloud.githubusercontent.com/assets/4003503/24270835/f6f6d728-0fd3-11e7-9429-cef375b1e969.jpg)

Downsize to `200 KB`, `png` format for result image:

```sh
$ downsize -s=204800 -f=png flower.jpg flower200kb.png
```

Resized result:

![flower200kb](https://cloud.githubusercontent.com/assets/4003503/24270862/1126aace-0fd4-11e7-8c06-769162a93abe.png)

Downsize to `100 KB`, `png` format for result image:

```sh
$ downsize -s=102400 -f=png flower.jpg flower100kb.png
```

Resized result:

![flower100kb](https://cloud.githubusercontent.com/assets/4003503/24270871/1b7d8e7a-0fd4-11e7-8b27-8b055a60201b.png)

# Sample 2

The original image:

![leaves](https://cloud.githubusercontent.com/assets/4003503/24270590/ffc8b070-0fd2-11e7-949f-3f76364ac252.jpg)

Downsize to `500 KB`, auto determine format for result image:

```sh
$ downsize -s=512000 leaves.jpg leaves500kb.jpg
```

Resized result:

![leaves500kb](https://cloud.githubusercontent.com/assets/4003503/24270890/2b3de260-0fd4-11e7-97b3-1d70d9f3874e.jpg)

Downsize to `200 KB`, auto determine format for result image:

```sh
$ downsize -s=204800 leaves.jpg leaves200kb.jpg
```

Resized result:

![leaves200kb](https://cloud.githubusercontent.com/assets/4003503/24270881/245cb76e-0fd4-11e7-86a4-b3547010e4f6.jpg)

Downsize to `100 KB`, auto determine format for result image:

```sh
$ downsize -s=102400 leaves.jpg leaves100kb.jpg
```

Resized result:

![leaves100kb](https://cloud.githubusercontent.com/assets/4003503/24271855/02c5bcfa-0fd8-11e7-8bbc-b1cf86751350.jpg)
