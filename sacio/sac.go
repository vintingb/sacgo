package sacio

import (
	"bufio"
	"fmt"
	"os"
)

type Sac struct {
	sacHead SacHead
	sacData SacData
}

func (sac *Sac) Sac2asc(filename string) error {
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()
	w := bufio.NewWriter(fd)
	_, err2 := w.WriteString(sac.sacHead.Format())
	if err2 != nil {
		return err2
	}
	data := sac.sacData.Data
	npts := int(sac.sacHead.Npts)
	for k := 0; k < npts-5; k += 5 {
		r := fmt.Sprintf("%10.2f     %10.2f      %10.2f      %10.2f      %10.2f\n", data[k], data[k+1], data[k+2], data[k+3], data[k+4])
		_, err3 := w.WriteString(r)
		if err3 != nil {
			return err3
		}
	}
	for k := npts % 5; k >= 1; k-- {
		r := fmt.Sprintf("%10.2f     ", data[npts-k])
		_, _ = w.WriteString(r)
	}
	_ = w.Flush()
	_ = fd.Sync()
	return nil
}

func (sac *Sac) Read(fileName string) error {
	err := sac.sacData.Read(fileName)
	if err != nil {
		return err
	}
	err = sac.sacHead.Read(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (sac *Sac) NewSac(head *SacHead, data *SacData) {
	sac.sacHead = *head
	sac.sacData = *data
}
