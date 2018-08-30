package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type shape struct {
	bg  draw.Image
	new color.Color
	f   func(x, y int, rect image.Rectangle) bool
}

func (h shape) Convert(c color.Color) color.Color {
	return h.new
}

func (h shape) ColorModel() color.Model {
	return h
}

func (h shape) Bounds() image.Rectangle {
	return h.bg.Bounds()
}

func (h shape) At(x int, y int) color.Color {
	if h.f(x, y, h.bg.Bounds()) {
		return h.new
	}
	return h.bg.At(x, y)
}

func Right(old draw.Image, new color.Color) *shape {
	return &shape{
		bg:  old,
		new: new,
		f:   righthalf,
	}
}

func Left(old draw.Image, new color.Color) *shape {
	return &shape{
		bg:  old,
		new: new,
		f:   lefthalf,
	}
}

func Upper(old draw.Image, new color.Color) *shape {
	return &shape{
		bg:  old,
		new: new,
		f:   upperhalf,
	}
}

func Lower(old draw.Image, new color.Color) *shape {
	return &shape{
		bg:  old,
		new: new,
		f:   lowerhalf,
	}
}

func Circle(old draw.Image, new color.Color) *shape {
	return &shape{
		bg:  old,
		new: new,
		f:   circle,
	}
}

func righthalf(x, y int, rect image.Rectangle) bool {
	return x > rect.Bounds().Dx()/2
}

func lefthalf(x, y int, rect image.Rectangle) bool {
	return x < rect.Bounds().Dx()/2
}

func upperhalf(x, y int, rect image.Rectangle) bool {
	return y < rect.Bounds().Dy()/2
}

func lowerhalf(x, y int, rect image.Rectangle) bool {
	return y > rect.Bounds().Dy()/2
}

func circle(x, y int, rect image.Rectangle) bool {
	//
	// point:  (x,y)
	// center: (a,b)
	// radius: r
	//
	// (x-a)^2+(y-b)^2 < r^2
	p := image.Pt(x, y)
	center := image.Pt(rect.Dx()/2, rect.Dy()/2)
	return math.Pow(float64(p.X-center.X), 2)+math.Pow(float64(p.Y-center.Y), 2) < math.Pow(float64(rect.Dx()/3), 2)
}
