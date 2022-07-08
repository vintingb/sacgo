package line

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type line struct {
	Ys
	draw.LineStyle
}

func NewLine(ys Ys) *line {
	return &line{
		Ys:        ys,
		LineStyle: newLineStyle(),
	}
}

func (line *line) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	ps := make([]vg.Point, line.Len())
	for i := range line.Ys {
		x, y := line.Ys.XY(i)
		ps[i].X = trX(x)
		ps[i].Y = trY(y)
	}

	lines := c.ClipLinesXY(ps)
	if line.LineStyle.Width != 0 && len(lines) != 0 {
		c.SetLineStyle(line.LineStyle)
		for _, l := range lines {
			if len(l) == 0 {
				continue
			}
			var p vg.Path
			prev := l[0]
			p.Move(prev)
			for _, pt := range l[1:] {
				p.Line(pt)
				prev = pt
			}
			c.Stroke(p)
		}
	}
}

func (line *line) DataRange() (xmin, xmax, ymin, ymax float64) {
	return plotter.XYRange(line)
}

func (line *line) Thumbnail(c *draw.Canvas) {

	if line.LineStyle.Width != 0 {
		y := c.Center().Y
		c.StrokeLine2(line.LineStyle, c.Min.X, y, c.Max.X, y)
	}
}

func AddLine(plt *plot.Plot, name string, pts Ys) {
	var ps []plot.Plotter
	l := NewLine(pts)
	l.Color = Color(0)
	ps = append(ps, l)
	plt.Legend.Add(name, l)
	plt.Add(ps...)
}

func AddLines(plt *plot.Plot, name []string, pts []Ys) error {
	var ps []plot.Plotter
	for k := range name {
		l := NewLine(pts[k])
		l.Color = Color(k)
		ps = append(ps, l)
		plt.Legend.Add(name[k], l)
	}
	plt.Add(ps...)
	return nil
}
