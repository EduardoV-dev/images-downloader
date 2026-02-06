# Images Downloader

A small CLI that reads image URLs from a text file and downloads each image to a local directory.

## Requirements

- Go 1.20+ (module-based build)

## Usage

From the project root:

```bash
go run .
```

By default it reads `images.txt` and writes to the `images/` folder.

### Flags

- `-f`, `--file`  Path to the text file containing image URLs (default: `images.txt`)
- `-o`, `--output`  Directory to save downloaded images (default: `images`)

### Examples

```bash
# Use defaults

go run .

# Custom input file

go run . --file my-images.txt

# Custom output directory

go run . --output downloads

# Custom input and output

go run . --file my-images.txt --output downloads
```

## Input file format

Place one image URL per line. Lines with multiple URLs are rejected.

```text
https://example.com/image1.jpg
https://example.com/image2.png
```

## Notes

- The downloader creates the output directory if it does not exist.
- Images are saved using the URL base name.
