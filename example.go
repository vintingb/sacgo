package main

import (
	"fmt"
	"goSac/sacio"
)

func main() {
	// read SacHead
	h := &sacio.SacHead{}
	h.ReadHead("sacio/test.SAC")
	fmt.Println(string(h.Kstnm[:]))
	fmt.Println(h.Nzyear)
	fmt.Println(h.Nzjday)
	fmt.Println(h.Nzhour)
	fmt.Println(h.Nzmin)
	fmt.Println(h.Nzsec)
	fmt.Println(h.Nzmsec)
	fmt.Println(h.Npts)
	// read SacData
	d := &sacio.SacData{}
	d.ReadData("sacio/test.SAC")
	fmt.Println(d.Data)
}
