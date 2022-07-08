package sac

import (
	"bytes"
	"encoding/binary"
	"github.com/vintingb/sacgo/line"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
	"io"
	"io/ioutil"
)

type sacData struct {
	Pts []float64
}

func (d *sacData) decode(r io.Reader) error {
	rp, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	data := make([]byte, 4)
	n := 0
	var k float32
	for {
		if n >= len(rp) {
			break
		}
		data = rp[n : n+4]
		n += 4
		_ = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &k)
		d.Pts = append(d.Pts, float64(k))
	}
	return nil
}

func (d *sacData) Plot() {

}

func (d *sacData) plot(name string) *plot.Plot {
	p := plot.New()
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	line.AddLine(p, name, d.Pts)
	return p
}
