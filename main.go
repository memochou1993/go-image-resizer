package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	input, _ := os.Open("example.png")
	defer input.Close()

	src, err := png.Decode(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	output, _ := os.Create("example_resized.png")
	defer output.Close()

	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	if err = png.Encode(output, dst); err != nil {
		log.Fatal(err)
	}
}
