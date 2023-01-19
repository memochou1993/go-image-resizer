package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

func main() {
	in, err := os.Open("in.png")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	src, _, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	transformed := image.NewRGBA(image.Rect(50, 50, 150, 150))
	draw.Copy(transformed, image.Point{0, 0}, src, src.Bounds(), draw.Over, nil)
	draw.ApproxBiLinear.Transform(transformed, f64.Aff3{}, src, src.Bounds(), draw.Over, nil)
	scaled := image.NewRGBA(image.Rect(0, 0, 120, 120))
	draw.ApproxBiLinear.Scale(scaled, scaled.Rect, transformed, transformed.Bounds(), draw.Over, nil)
	out, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	err = png.Encode(out, scaled)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image has bounds %v.\n", scaled.Bounds())
}
