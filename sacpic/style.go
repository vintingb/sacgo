package sacpic

import (
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"sync"
)

var (
	wg        sync.WaitGroup
	PicHeight vg.Length = 5
	PicWeight vg.Length = 20

	LineSize = 0.1

	// FileType .eps, .jpg, .jpeg, .pdf, .png, .svg, .tex, .tif and .tiff.
	FileType = ".pdf"
)

type mod int

const (
	plot0 mod = iota
	plot1
	plot2
	plotFft0
	plotFft1
)

func newPicSize() (vg.Length, vg.Length) {
	return PicWeight * vg.Centimeter, PicHeight * vg.Centimeter
}

func newLineStyle() draw.LineStyle {
	return draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(LineSize),
		DashOffs: 0,
	}
}
