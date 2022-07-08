package sacio

import (
	"bytes"
	"encoding/binary"
	"gonum.org/v1/gonum/dsp/fourier"
	"io/ioutil"
	"math/cmplx"
	"os"
)

type SacData struct {
	Data []float64
	Abs  []float64
}

func (d *SacData) Read(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
		}
	}(fp)
	r, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	r = r[632:]
	data := make([]byte, 4)
	n := 0
	var k float32
	for {
		if n >= len(r) {
			break
		}
		data = r[n : n+4]
		n += 4
		_ = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &k)
		d.Data = append(d.Data, float64(k))
	}
	return nil
}

func (d *SacData) Fft() {
	var lens = len(d.Data)
	fft := fourier.NewFFT(lens)
	coeff := fft.Coefficients(nil, d.Data)
	for _, c := range coeff {
		d.Abs = append(d.Abs, cmplx.Abs(c))
	}
}
