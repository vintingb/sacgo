package sacpic

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
)

type line struct {
	plotter.XYs
	draw.LineStyle
	FillColor color.Color
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

	if pts.FillColor != nil && len(ps) > 0 {
		minY := trY(plt.Y.Min)
		fillPoly := []vg.Point{{X: ps[0].X, Y: minY}}
		fillPoly = append(fillPoly, ps...)
		fillPoly = append(fillPoly, vg.Point{X: ps[len(ps)-1].X, Y: minY})
		fillPoly = c.ClipPolygonXY(fillPoly)
		if len(fillPoly) > 0 {
			c.SetColor(pts.FillColor)
			var pa vg.Path
			prev := fillPoly[0]
			pa.Move(prev)
			for _, pt := range fillPoly[1:] {
				pa.Line(pt)
				prev = pt
			}
			pa.Close()
			c.Fill(pa)
		}
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
	if pts.FillColor != nil {
		var topY vg.Length
		if pts.LineStyle.Width == 0 {
			topY = c.Max.Y
		} else {
			topY = (c.Min.Y + c.Max.Y) / 2
		}
		points := []vg.Point{
			{X: c.Min.X, Y: c.Min.Y},
			{X: c.Min.X, Y: topY},
			{X: c.Max.X, Y: topY},
			{X: c.Max.X, Y: c.Min.Y},
		}
		poly := c.ClipPolygonY(points)
		c.FillPolygon(pts.FillColor, poly)
	}

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
