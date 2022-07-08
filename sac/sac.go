package sac

import (
	"gonum.org/v1/plot/vg"
	"io"
	"log"
)

type Sac struct {
	SacHead *sacHead
	SacData *sacData
}

func Decode(r io.Reader) *Sac {
	h := new(sacHead)
	d := new(sacData)
	err := h.decode(r)
	if err != nil {
		log.Fatalln(err)
	}
	err = d.decode(r)
	if err != nil {
		log.Fatalln(err)
	}
	return &Sac{SacHead: h, SacData: d}
}

func (s *Sac) Plot(name string) (io.WriterTo, error) {
	p := s.SacData.plot(name)
	c, err := p.WriterTo(20*vg.Inch, 5*vg.Inch, "pdf")
	if err != nil {
		return nil, err
	}
	return c, nil
}
