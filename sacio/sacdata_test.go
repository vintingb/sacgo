package sacio

import (
	"testing"
)

func TestSacData_ReadData(t *testing.T) {
	d := new(SacData)
	d.Read("test.SAC")
	if len(d.Data) != 803120 {
		t.Error("error")
	}
}
