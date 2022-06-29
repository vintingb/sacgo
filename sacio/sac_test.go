package sacio

import "testing"

func TestSac_Read(t *testing.T) {
	sac := new(Sac)
	err := sac.Read("test.SAC")
	if err != nil {
		t.Error("read file error", err)
	}
	if len(sac.sacData.Data) != 803120 {
		t.Error("error")
	}
	if sac.sacHead.Npts != 803116 {
		t.Error("error")
	}
}

func TestSac_Sac2asc(t *testing.T) {
	sac := new(Sac)
	err := sac.Read("test.SAC")
	if err != nil {
		t.Error("read file error", err)
	}
	err = sac.Sac2asc("test.ASC")
	if err != nil {
		t.Error("error")
	}
}

func TestSac_NewSac(t *testing.T) {
	sac := new(Sac)
	sacHead := new(SacHead)
	sacData := new(SacData)
	err := sacHead.Read("test.SAC")
	if err != nil {
		t.Error("read file error", err)
	}
	err = sacData.Read("test.SAC")
	if err != nil {
		t.Error("read file error", err)
	}
	sac.NewSac(sacHead, sacData)
	if len(sac.sacData.Data) != 803120 {
		t.Error("error")
	}
	if sac.sacHead.Npts != 803116 {
		t.Error("error")
	}
}
