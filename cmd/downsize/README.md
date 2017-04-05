# downsize

Reduces an image to a specified file size in bytes.

# Installation

```bash
$ go get -u github.com/lelenanam/downsize/...
```

# Usage

You can specify the size in bytes and the format for the output file. For `jpeg` format, you can specify the quality.

```
Usage of downsize:
downsize [-s=size] [-f=format] [-q=jpeg quality] [-i=infile] [-o=outfile]
  -f string
    	format: jpeg, png or gif, by default the format of an image is determined during decoding
  -i string
    	input file name, required
  -o string
    	output file name, required
  -q int
    	desired output jpeg quality, ranges from 1 to 100 inclusive, higher is better (default 80)
  -s int
    	desired output file size in bytes (default 204800)
```


## Example

Resize the file `image.jpg` to size `1 MB` and save the result in `jpeg` format file `resized.jpg`.

```sh
$ downsize -s=1048576 -f=jpeg image.jpg resized.jpg
```

## Sample 1

The original image `2.4 MB`:

![flower](https://cloud.githubusercontent.com/assets/4003503/24270582/f352a102-0fd2-11e7-852e-7ea77c4eae82.jpg)

Downsize to `1 MB`, auto determine format for result image:

```sh
$ downsize -s=1048576 flower.jpg flower1mb.jpg
```

Resized result:

![flower1mb](https://cloud.githubusercontent.com/assets/4003503/24625151/f6576e30-1862-11e7-89cd-aa6ebbc21e3f.jpg)

Downsize to `200 KB`, `jpeg` format for result image:

```sh
$ downsize -s=204800 -f=jpeg flower.jpg flower200kb.jpg
```

Resized result:

![flower200kb](https://cloud.githubusercontent.com/assets/4003503/24625184/120b66fe-1863-11e7-9cab-42af6bb2aa71.jpg)

Downsize to `200 KB`, `png` format for result image:

```sh
$ downsize -s=204800 -f=png flower.jpg flower200kb.png
```

Resized result:

![flower200kb](https://cloud.githubusercontent.com/assets/4003503/24625215/26a34bfe-1863-11e7-9d5f-3258a8aa71ce.png)


## Sample 2

The original image `3.4 MB`:

![leaves](https://cloud.githubusercontent.com/assets/4003503/24270590/ffc8b070-0fd2-11e7-949f-3f76364ac252.jpg)

Downsize to `200 KB`, auto determine format for result image, default quality:

```sh
$ downsize -s=204800 leaves.jpg leaves200kb.jpg
```

Resized result:

![leaves200kb](https://cloud.githubusercontent.com/assets/4003503/24625297/690b42d0-1863-11e7-86f3-bb90358b009d.jpg)

Downsize to `200 KB`, auto determine format for result image, quality `50`:

```sh
$ downsize -s=204800 -q=50 leaves.jpg leaves200kbQ50.jpg
```

Resized result:

![leaves200kbq50](https://cloud.githubusercontent.com/assets/4003503/24625339/8c90db3e-1863-11e7-9a9d-227980e19464.jpg)

Downsize to `100 KB`, auto determine format for result image:

```sh
$ downsize -s=102400 leaves.jpg leaves100kb.jpg
```

Resized result:

![leaves100kb](https://cloud.githubusercontent.com/assets/4003503/24625357/9f83193c-1863-11e7-99c7-2cc912f5b723.jpg)