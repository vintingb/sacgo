package sacio

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
)

type SacData struct {
	Data []float32
}

func (d *SacData) ReadData(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
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
		d.Data = append(d.Data, k)
	}
	return nil
}
