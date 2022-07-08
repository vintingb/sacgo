package sacpic

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type line struct {
	plotter.XYs
	draw.LineStyle
}

func NewLine(data plotter.XYs) (*line, error) {
	return &line{
		XYs:       data,
		LineStyle: newLineStyle(),
	}, nil
}

func (pts *line) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	ps := make([]vg.Point, len(pts.XYs))

	for i, p := range pts.XYs {
		ps[i].X = trX(p.X)
		ps[i].Y = trY(p.Y)
	}

	lines := c.ClipLinesXY(ps)
	if pts.LineStyle.Width != 0 && len(lines) != 0 {
		c.SetLineStyle(pts.LineStyle)
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

func (pts *line) DataRange() (xmin, xmax, ymin, ymax float64) {
	return plotter.XYRange(pts)
}

func (pts *line) Thumbnail(c *draw.Canvas) {
	if pts.LineStyle.Width != 0 {
		y := c.Center().Y
		c.StrokeLine2(pts.LineStyle, c.Min.X, y, c.Max.X, y)
	}
}

func AddLines(plt *plot.Plot, name []string, pts []plotter.XYs) error {
	var ps []plot.Plotter
	for k := range name {
		l, err := NewLine(pts[k])
		if err != nil {
			return err
		}
		l.Color = Color(k)
		ps = append(ps, l)
		plt.Legend.Add(name[k], l)
	}
	plt.Add(ps...)
	return nil
}
