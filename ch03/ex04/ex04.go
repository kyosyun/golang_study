package main

import (
	"net/http"
	"log"
	"fmt"
	"math"
	"net/url"
)

const (
	cells         = 100
	xyrange       = 30.0
	angle         = math.Pi / 6
)

var width  = 500.0
var height = 300.0
var xyscale  = width / 2 / xyrange
var zscale  = height * 0.5

var sin30, cos30 = math.Sin(angle), math.Cos(angle)


func main() {
	// アクセスをhandlerに飛ばす
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.Path)
	query := u.Query()
	if(query["height"] != nil) {
		height =
	}
	w.Header().Set("Content-Type","image/svg+xml")
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	printSvgData(w)
}


func printSvgData(w http.ResponseWriter) {
	fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsInf(ax, 0) ||
				math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) ||
				math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0) ||
				math.IsNaN(ax) ||
				math.IsNaN(ay) ||
				math.IsNaN(bx) ||
				math.IsNaN(by) ||
				math.IsNaN(cx) ||
				math.IsNaN(cy) ||
				math.IsNaN(dx) ||
				math.IsNaN(dy) {
				continue
			}

			fmt.Fprintf(w,"<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w,"</svg>")
}


func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Copysign(x, y)
	return math.Sin(r)
}
