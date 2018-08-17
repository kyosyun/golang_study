package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)


func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, sampledMandelbrot(x,y))
		}
	}
	png.Encode(os.Stdout, img)
}

func sampledMandelbrot(x, y float64) color.Color {
	var rdSum, grSum, blSum uint32
	for i := -1 ; i < 2; i++ {
		for j := -1; j < 2; j++ {
			z := complex(x + float64(i/ width * (xmax - xmin)), y + float64(j/ width * (xmax - xmin)))
			r, g, b, _ := mandelbrot(z).RGBA()
			rdSum += uint32(r)
			grSum += uint32(g)
			blSum += uint32(b)
		}
	}
	//mandelbrotColor := mandelbrot(complex(x,y))
	//r,g,b,_ := mandelbrotColor.RGBA()
	//return color.RGBA{uint8(r)/2,uint8(g)/2,uint8(b)/2, 0xff}
	return color.RGBA{ uint8(rdSum/9), uint8(grSum/9), uint8(blSum/9), 0xff}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast5 = 5
	const contrast10 = 10
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{n * contrast5, 255 - n*contrast10, n * contrast5, 0xff}
		}
	}
	return color.Black
}
