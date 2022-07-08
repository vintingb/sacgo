package sacpic

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"log"
	"os"
	"testing"
)

func TestTick(t *testing.T) {
	p := plot.New()
	p.Title.Text = "Title"
	p.X.Label.Text = "X [mm]"
	p.Y.Label.Text = "Ys [A.U.]"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	p.X.Min = 100
	p.X.Max = 300
	p.Y.Min = 1e8
	p.Y.Max = 2e9

	p.X.Tick.Marker = DefaultTicks{}
	p.Y.Tick.Marker = DefaultTicks{}

	err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "DefaultTicks.png")
	if err != nil {
		log.Fatalf("could not save line: %+v", err)
	}
	ExampleAlign()
}

func ExampleAlign() {
	const rows, cols = 4, 2
	plots := make([][]*plot.Plot, rows)
	for j := 0; j < rows; j++ {
		plots[j] = make([]*plot.Plot, cols)
		for i := 0; i < cols; i++ {
			p := plot.New()
			plots[j][i] = p
		}
	}

	img := vgimg.New(vg.Points(500), vg.Points(700))
	dc := draw.New(img)

	t := draw.Tiles{
		Rows:      rows,
		Cols:      cols,
		PadX:      vg.Millimeter,
		PadY:      vg.Millimeter,
		PadTop:    vg.Points(2),
		PadBottom: vg.Points(2),
		PadLeft:   vg.Points(2),
		PadRight:  vg.Points(2),
	}

	canvases := plot.Align(plots, t, dc)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if plots[j][i] != nil {
				plots[j][i].Draw(canvases[j][i])
			}
		}
	}

	w, err := os.Create("align.png")
	if err != nil {
		panic(err)
	}
	defer w.Close()
	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		panic(err)
	}
}
