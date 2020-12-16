package sacio

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
)

func (d *SacData) ReadData(fileName string) {
	fp, _ := os.Open(fileName)
	defer fp.Close()
	r, _ := ioutil.ReadAll(fp)
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
}
