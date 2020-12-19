package sacpic

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"sacgo/sacio"
	"sync"
)

var Wg sync.WaitGroup

type PicData struct {
	Sums int32
	X    []float32
	Y    []float32
}

func (d *PicData) NewPicData(sacHead *sacio.SacHead, sacData *sacio.SacData) {
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
	err = plotutil.AddLines(p, fileName+".SAC", pts)
	if err != nil {
		panic(err)
	}
	if err := p.Save(10*vg.Inch, 5*vg.Inch, fileName+".pdf"); err != nil {
		panic(err)
	}
	defer Wg.Done()
}
