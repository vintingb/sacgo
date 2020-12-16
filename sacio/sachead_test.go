package sacio

import (
	"fmt"
	"testing"
)

func TestSacHead_ReadHead(t *testing.T) {
	h := new(SacHead)
	h.ReadHead("test.SAC")
	if h.Npts != 803116 {
		t.Error("error")
	}
}

func TestSacHead_Format(t *testing.T) {
	h := new(SacHead)
	h.ReadHead("test.SAC")
	t.Log(h.Format())
}

func TestSacHead_HeadToString(t *testing.T) {
	h := new(SacHead)
	h.ReadHead("test.SAC")
	hs := h.HeadToString()
	fmt.Println(hs.Kstnm)
	if hs.Kstnm != "PKD     " {
		t.Error("error")
	}
}
