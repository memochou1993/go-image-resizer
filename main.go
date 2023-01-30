package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

var (
	config Config
)

type Config struct {
	In     string
	Out    string
	X0     int
	Y0     int
	Width  int
	Height int
	Size   int
}

func init() {
	flag.StringVar(&config.In, "i", "in.png", "input image")
	flag.StringVar(&config.Out, "o", "out.png", "output image")
	flag.IntVar(&config.X0, "x", 0, "x of the image")
	flag.IntVar(&config.Y0, "y", 0, "y of the image")
	flag.IntVar(&config.Width, "w", 0, "width of the image")
	flag.IntVar(&config.Height, "h", 0, "height of the image")
	flag.IntVar(&config.Size, "s", 480, "size of the image")
	flag.Parse()
}

func main() {
	in, err := os.Open(config.In)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	src, _, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	width := config.Width
	if width == 0 {
		width = src.Bounds().Max.X
	}
	height := config.Height
	if height == 0 {
		height = src.Bounds().Max.Y
	}
	transformed := image.NewRGBA(image.Rect(config.X0, config.Y0, config.X0+width, config.Y0+height))
	draw.Copy(transformed, image.Point{0, 0}, src, src.Bounds(), draw.Over, nil)
	draw.BiLinear.Transform(transformed, f64.Aff3{}, src, src.Bounds(), draw.Over, nil)
	scaled := image.NewRGBA(image.Rect(0, 0, config.Size, config.Size))
	draw.BiLinear.Scale(scaled, scaled.Rect, transformed, transformed.Bounds(), draw.Over, nil)
	out, err := os.Create(config.Out)
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
