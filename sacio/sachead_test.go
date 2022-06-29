package sacio

import (
	"testing"
)

func TestSacHead_ReadHead(t *testing.T) {
	h := new(SacHead)
	h.Read("test.SAC")
	if h.Npts != 803116 {
		t.Error("error")
	}
}

func TestSacHead_Format(t *testing.T) {
	h := new(SacHead)
	h.Read("test.SAC")
	t.Log(h.Format())
}
