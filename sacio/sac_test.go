package sacio

import "testing"

func TestSac_Read(t *testing.T) {
	sac := new(Sac)
	sac.Read("test.SAC")
	if len(sac.sacData.Data) != 803120 {
		t.Error("error")
	}
	if sac.sacHead.Npts != 803116 {
		t.Error("error")
	}
}

func TestSac_Sac2asc(t *testing.T) {
	sac := new(Sac)
	sac.Read("test.SAC")
	err := sac.Sac2asc("test.ASC")
	if err != nil {
		t.Error("error")
	}
}
