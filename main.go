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

var (
	useCircle   = flag.Bool("circle", false, "paint circle for favicon")
	whiteBG     = flag.Bool("whiteBG", false, "white background")
	transparent = flag.Bool("t", false, "use transparent background")
)

// You should check out go blog.
// https://blog.golang.org/go-imagedraw-package
func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	favicon := createFavicon()
	err := png.Encode(os.Stdout, favicon)
	if err != nil {
		log.Fatal("failed to write favicon: ", err)
	}
}

func createFavicon() image.Image {
	picture := background(bgColor())
	if *useCircle {
		fillCircle(picture, randomColor())
		return picture
	}
	fillRight(picture, randomColor())

	return picture
}

func bgColor() color.Color {
	if *whiteBG {
		return color.White
	}
	return randomColor()
}

func background(c color.Color) draw.Image {
	rect := faviconRect()
	m := image.NewRGBA(rect)
	if !*transparent {
		draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	}
	return m
}

func faviconRect() image.Rectangle {
	return image.Rect(faviconX0, faviconY0, *faviconLength, *faviconLength)
}

func fillRight(bg draw.Image, c color.Color) {
	fillshape(bg, Right(bg, c))
}

func fillLower(bg draw.Image, c color.Color) {
	fillshape(bg, Lower(bg, c))
}

func fillCircle(bg draw.Image, c color.Color) {
	fillshape(bg, Circle(bg, c))
}

func fillshape(bg draw.Image, s *shape) {
	draw.Draw(bg, bg.Bounds(), s, image.ZP, draw.Src)
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
