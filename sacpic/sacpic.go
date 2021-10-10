package sacpic

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
	"os/exec"
	"runtime"
	"sacgo/sacio"
	"sync"
)

var (
	Wg        sync.WaitGroup
	PicHeight vg.Length = 5
	PicWeight vg.Length = 20

	LineSize = 0.1

	// FileType .eps, .jpg, .jpeg, .pdf, .png, .svg, .tex, .tif and .tiff.
	FileType = ".pdf"
)

func newSize() (vg.Length, vg.Length) {
	return PicWeight * vg.Centimeter, PicHeight * vg.Centimeter
}

func newLineStyle() draw.LineStyle {
	return draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(LineSize),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}
}

type PicData struct {
	Sums int32
	X    []float32
	Y    []float32
	time string
}

func (d *PicData) NewPicData(sacHead *sacio.SacHead, sacData *sacio.SacData) {
	d.time = kzdate(sacHead.Nzyear, sacHead.Nzjday)
	beginTime := sacHead.B
	overTime := sacHead.E
	npts := sacHead.Npts
	d.Sums = npts
	dt := (overTime - beginTime) / float32(npts-1)
	var i int32
	for i = 0; i < npts; i++ {
		d.X = append(d.X, beginTime+dt*float32(i))
		d.Y = sacData.Data[:npts]
	}
}

func (d *PicData) SavePic(fileName string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	pts := make(plotter.XYs, d.Sums)
	for i := range pts {
		pts[i].Y = float64(d.Y[i])
		pts[i].X = float64(d.X[i])
	}
	err = AddLines(p, fileName+".SAC", pts)
	if err != nil {
		panic(err)
	}
	p.Title.Text = d.time
	W, H := newSize()
	if err := p.Save(W, H, fileName+FileType); err != nil {
		panic(err)
	}
	defer Wg.Done()
	OpenBrowser(fileName + FileType)

}

func AddLines(plt *plot.Plot, name string, pts plotter.XYs) error {
	var ps []plot.Plotter

	l, err := NewLine(pts)
	if err != nil {
		return err
	}
	l.Color = color.RGBA{A: 255}
	l.Dashes = plotutil.Dashes(0)
	ps = append(ps, l)

	plt.Add(ps...)
	plt.Legend.Add(name, l)
	return nil
}
func NewLine(xys plotter.XYer) (*plotter.Line, error) {
	data, err := plotter.CopyXYs(xys)
	if err != nil {
		return nil, err
	}
	return &plotter.Line{
		XYs:       data,
		LineStyle: newLineStyle(),
	}, nil
}

func kzdate(nzyear, nzjday int32) string {
	mons := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	days := []int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (nzyear%4 == 0 && nzyear%100 != 0) || nzyear%400 == 0 {
		days[1] = 29
	}
	var daySum int32
	for index, value := range days {
		daySum += value
		if daySum > nzjday {
			daySum -= value
			return fmt.Sprintf("   kzdate = %v %v (%v), %v", mons[index], nzjday-daySum, nzjday, nzyear)
		}

	}
	return ""
}

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
