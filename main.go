package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	faviconX0 = 0
	faviconY0 = 0
)

var (
	faviconLength = flag.Int("l", 32, "length for favicon")
)

// You should check out go blog.
// https://blog.golang.org/go-imagedraw-package
func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	rect := faviconRect()
	favicon := image.NewRGBA(rect)
	c := randomColor()
	draw.Draw(favicon, favicon.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	drawRight(favicon, randomColor())
	err := png.Encode(os.Stdout, favicon)
	if err != nil {
		log.Fatal("failed to write favicon: ", err)
	}
}

func faviconRect() image.Rectangle {
	return image.Rect(faviconX0, faviconY0, *faviconLength, *faviconLength)
}

func drawRight(bg draw.Image, c color.Color) {
	for x := 0; x < bg.Bounds().Dx(); x++ {
		for y := 0; y < bg.Bounds().Dy(); y++ {
			if x > bg.Bounds().Dx()/2 {
				bg.Set(x, y, c)
			}
		}
	}
}

func randomColor() color.Color {
	// 0 <= uint8 < 255
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
