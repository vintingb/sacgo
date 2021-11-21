package sacpic

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgpdf"
	"image/color"
	"os"
	"sacgo/sacio"
)

type PicData struct {
	sums int32
	xys  plotter.XYs
	time string
}

func newFftPicData(sacHead []sacio.SacHead, sacData []sacio.SacData) []PicData {
	picData := make([]PicData, len(sacHead))
	for k := range picData {
		picData[k].time = kzdate(sacHead[k].Nzyear, sacHead[k].Nzjday)
		npts := len(sacData[k].Abs)
		picData[k].sums = int32(npts)
		fs := 1 / sacHead[k].Delta
		var i int
		var xys = make(plotter.XYs, npts)
		for i = 0; i < npts; i++ {
			xys[i].X = float64(fs) * float64(i) / float64(npts)
			xys[i].Y = sacData[k].Abs[i]
		}
		picData[k].xys = xys
	}
	return picData
}

func NewPicData(sacHead []sacio.SacHead, sacData []sacio.SacData) []PicData {
	picData := make([]PicData, len(sacHead))
	for k := range picData {
		picData[k].time = kzdate(sacHead[k].Nzyear, sacHead[k].Nzjday)
		beginTime := sacHead[k].B
		npts := sacHead[k].Npts
		picData[k].sums = npts
		dt := sacHead[k].Delta
		var i int32
		var xys = make(plotter.XYs, npts)
		for i = 0; i < npts; i++ {
			xys[i].X = float64(beginTime + dt*float32(i))
			xys[i].Y = sacData[k].Data[i]
		}
		picData[k].xys = xys
	}

	return picData
}

func SavePic(d *[]PicData, fileNames []string, m mod) {
	switch m {
	case plot0:
		//  plot
		for k := range fileNames {
			go func(k int) {
				wg.Add(1)
				p := plot.New()
				var fn = []string{fileNames[k]}
				var xys = []plotter.XYs{(*d)[k].xys}
				err := AddLines(p, fn, xys)
				if err != nil {
					panic(err)
				}
				p.Title.Text = (*d)[k].time
				W, H := newPicSize()
				p.X.Tick.Marker = DefaultTicks{}
				p.Y.Tick.Marker = DefaultTicks{}
				xText := Magnitude(p.X.Min, p.X.Max)
				yText := Magnitude(p.Y.Min, p.Y.Max)
				p.X.Label.Text = fmt.Sprintf("X:10+%d", xText)
				p.Y.Label.Text = fmt.Sprintf("Y:10+%d", yText)
				if err := p.Save(W, H, fileNames[k]+FileType); err != nil {
					panic(err)
				}
				wg.Done()
			}(k)

		}
	case plot1:
		//	plot1
		var rows = len(fileNames)
		plots := make([][]*plot.Plot, rows)
		for j := 0; j < rows; j++ {
			plots[j] = make([]*plot.Plot, 1)
			p := plot.New()
			var fn = []string{fileNames[j]}
			var xys = []plotter.XYs{(*d)[j].xys}
			err := AddLines(p, fn, xys)
			if err != nil {
				panic(err)
			}
			p.Title.Text = (*d)[j].time
			p.X.Tick.Marker = DefaultTicks{}
			p.Y.Tick.Marker = DefaultTicks{}
			xText := Magnitude(p.X.Min, p.X.Max)
			if xText > 0 {
				p.X.Label.Text = fmt.Sprintf("X:10+%d", xText)
			} else if xText < 0 {
				p.X.Label.Text = fmt.Sprintf("X:10%d", xText)
			}
			yText := Magnitude(p.Y.Min, p.Y.Max)
			p.Y.Label.Text = fmt.Sprintf("Y:10+%d", yText)
			plots[j][0] = p
		}
		W, H := newPicSize()
		pdf := vgpdf.New(W, H*vg.Length(rows))
		dc := draw.New(pdf)
		t := draw.Tiles{
			Rows:      rows,
			Cols:      1,
			PadX:      vg.Millimeter,
			PadY:      vg.Millimeter,
			PadTop:    vg.Points(2),
			PadBottom: vg.Points(2),
			PadLeft:   vg.Points(2),
			PadRight:  vg.Points(2),
		}

		canvases := plot.Align(plots, t, dc)
		for j := 0; j < rows; j++ {
			plots[j][0].Draw(canvases[j][0])
		}

		w, err := os.Create("P1" + FileType)
		if err != nil {
			panic(err)
		}
		defer w.Close()
		if _, err := pdf.WriteTo(w); err != nil {
			panic(err)
		}
	case plot2:
		//	plot2
		p := plot.New()

		var xys = make([]plotter.XYs, len(fileNames))
		for i := range *d {
			xys[i] = (*d)[i].xys
		}
		err := AddLines(p, fileNames, xys)
		p.Title.Text = "Plot2"
		if err != nil {
			panic(err)
		}
		W, H := newPicSize()
		p.X.Tick.Marker = DefaultTicks{}
		p.Y.Tick.Marker = DefaultTicks{}
		xText := Magnitude(p.X.Min, p.X.Max)
		yText := Magnitude(p.Y.Min, p.Y.Max)
		p.X.Label.Text = fmt.Sprintf("X:10+%d", xText)
		p.Y.Label.Text = fmt.Sprintf("Y:10+%d", yText)
		if err := p.Save(W, H, "P2"+FileType); err != nil {
			panic(err)
		}
	case plotFft0:
		for k := range fileNames {
			go func(k int) {
				wg.Add(1)
				p := plot.New()
				var fn = []string{fileNames[k]}
				var xys = []plotter.XYs{(*d)[k].xys}
				var ps []plot.Plotter
				for i := range fn {
					l := &line{
						XYs: xys[i],
						LineStyle: draw.LineStyle{
							Color:    color.Black,
							Width:    vg.Points(0.5),
							DashOffs: 0,
						},
					}
					l.Color = Color(i)
					ps = append(ps, l)
					p.Legend.Top = true
					p.Legend.Add(fn[i]+"  FFT", l)
				}
				p.Add(ps...)
				p.Title.Text = (*d)[k].time
				W, H := newPicSize()
				p.X.Tick.Marker = DefaultTicks{}
				p.Y.Tick.Marker = DefaultTicks{}
				xText := Magnitude(p.X.Min, p.X.Max)
				yText := Magnitude(p.Y.Min, p.Y.Max)
				p.X.Label.Text = fmt.Sprintf("X:10+%d", xText)
				p.Y.Label.Text = fmt.Sprintf("Y:10+%d", yText)
				if err := p.Save(W, H, fileNames[k]+"_FFT"+FileType); err != nil {
					panic(err)
				}
				wg.Done()
			}(k)

		}
	case plotFft1:
		var rows = len(fileNames)
		plots := make([][]*plot.Plot, rows)
		for j := 0; j < rows; j++ {
			plots[j] = make([]*plot.Plot, 1)
			p := plot.New()
			var fn = []string{fileNames[j]}
			var xys = []plotter.XYs{(*d)[j].xys}
			var ps []plot.Plotter
			for i := range fn {
				l := &line{
					XYs: xys[i],
					LineStyle: draw.LineStyle{
						Color:    color.Black,
						Width:    vg.Points(0.5),
						DashOffs: 0,
					},
				}
				l.Color = Color(i)
				ps = append(ps, l)
				p.Legend.Top = true
				p.Legend.Add(fn[i]+"  FFT", l)
			}
			p.Add(ps...)
			p.Title.Text = (*d)[j].time
			p.X.Tick.Marker = DefaultTicks{}
			p.Y.Tick.Marker = DefaultTicks{}
			xText := Magnitude(p.X.Min, p.X.Max)
			if xText > 0 {
				p.X.Label.Text = fmt.Sprintf("X:10+%d", xText)
			} else if xText < 0 {
				p.X.Label.Text = fmt.Sprintf("X:10%d", xText)
			}
			yText := Magnitude(p.Y.Min, p.Y.Max)
			p.Y.Label.Text = fmt.Sprintf("Y:10+%d", yText)
			plots[j][0] = p
		}
		W, H := newPicSize()
		pdf := vgpdf.New(W, H*vg.Length(rows))
		dc := draw.New(pdf)
		t := draw.Tiles{
			Rows:      rows,
			Cols:      1,
			PadX:      vg.Millimeter,
			PadY:      vg.Millimeter,
			PadTop:    vg.Points(2),
			PadBottom: vg.Points(2),
			PadLeft:   vg.Points(2),
			PadRight:  vg.Points(2),
		}

		canvases := plot.Align(plots, t, dc)
		for j := 0; j < rows; j++ {
			plots[j][0].Draw(canvases[j][0])
		}

		w, err := os.Create("PF1" + FileType)
		if err != nil {
			panic(err)
		}
		defer w.Close()
		if _, err := pdf.WriteTo(w); err != nil {
			panic(err)
		}
		wg.Wait()
	}
}
