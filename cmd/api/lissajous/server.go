package lissajous

import (
	"github.com/spf13/cobra"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var (
	startCmd = &cobra.Command{
		Use: "lissajous",
		Short: "l",
		Example: "mogo lissajous > out.gif",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	palette = []color.Color{color.White, color.Black}
)

const (
	whiteIndex = 0
	blackIndex = 1
)

func run() error {

	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)

	return nil
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
