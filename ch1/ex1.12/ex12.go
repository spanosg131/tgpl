// ex12 creates a webserver and displays a lissajous animated gif.
// Passing http://localhost:8000/?cycles=NN sets the number of cycles
// to NN. The default cycle count is 5
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, c int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // images canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	cycles := c // number of complete x oscillator revolutions
	if c == 0 {
		cycles = 5 // set the default value
	}
	freq := rand.Float64() * 3.0 // relative fequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Form.Get("cycles") != "" {
		cycles, _ := strconv.Atoi(r.Form.Get("cycles"))
		lissajous(w, cycles)

	} else {
		lissajous(w, 5)
	}
}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
