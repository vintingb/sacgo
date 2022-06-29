package sacio

import (
	"fmt"
	"github.com/lunixbochs/struc"
	"os"
	"reflect"
	"strings"
)

type SacHead struct {
	Delta, Depmin, Depmax, Scale, Odelta             float32 `struc:"little"` //0
	B, E, O, A, Internal1                            float32 `struc:"little"` //20
	T0, T1, T2, T3, T4                               float32 `struc:"little"` //40
	T5, T6, T7, T8, T9                               float32 `struc:"little"` //60
	F, Resp0, Resp1, Resp2, Resp3                    float32 `struc:"little"` //80
	Resp4, Resp5, Resp6, Resp7, Resp8                float32 `struc:"little"` //100
	Resp9, Stla, Stlo, Stel, Stdp                    float32 `struc:"little"` //120
	Evla, Evlo, Evel, Evdp, Mag                      float32 `struc:"little"` //140
	User0, User1, User2, User3, User4                float32 `struc:"little"` //160
	User5, User6, User7, User8, User9                float32 `struc:"little"` //180
	Dist, Az, Baz, Gcarc, Internal2                  float32 `struc:"little"` //200
	Internal3, Depmen, Cmpaz, Cmpinc, Xminimum       float32 `struc:"little"` //220
	Xmaximum, Yminimum, Ymaximum, Unused1, Unused2   float32 `struc:"little"` //240
	Unused3, Unused4, Unused5, Unused6, Unused7      float32 `struc:"little"` //260
	Nzyear, Nzjday, Nzhour, Nzmin, Nzsec             int32   `struc:"little"` //280
	Nzmsec, Nvhdr, Norid, Nevid, Npts                int32   `struc:"little"` //300
	Internal, Nwfid, Nxsize, Nysize, Unused8         int32   `struc:"little"` //320
	Iftype, Idep, Iztype, Unused9, Iinst             int32   `struc:"little"` //340
	Istreg, Ievreg, Ievtyp, Iqual, Isynth            int32   `struc:"little"` //360
	Imagtyp, Imagsrc, Unused10, Unused11, Unused12   int32   `struc:"little"` //380
	Unused13, Unused14, Unused15, Unused16, Unused17 int32   `struc:"little"` //400
	Leven, Lpspol, Lovrok, Lcalda, Unused18          int32   `struc:"little"` //420
	Kstnm                                            string  `struc:"little,[8]byte"`
	Kevnm                                            string  `struc:"little,[16]byte"`
	Khole, Ko, Ka                                    string  `struc:"little,[8]byte"`
	Kt0, Kt1, Kt2                                    string  `struc:"little,[8]byte"`
	Kt3, Kt4, Kt5                                    string  `struc:"little,[8]byte"`
	Kt6, Kt7, Kt8                                    string  `struc:"little,[8]byte"`
	Kt9, Kf, KUser0                                  string  `struc:"little,[8]byte"`
	Kuser1, Kuser2, Kcmpnm                           string  `struc:"little,[8]byte"`
	KnetwK, Kdatrd, Kinst                            string  `struc:"little,[8]byte"`
}

func (h *SacHead) Read(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
		}
	}(fp)
	err = struc.Unpack(fp, h)
	if err != nil {
		return err
	}
	return nil
}

func (h *SacHead) Format() string {
	var tmp strings.Builder
	v := reflect.ValueOf(*h)
	for k := 0; k < 70; k += 5 {
		r := fmt.Sprintf("%12.2f     %12.2f      %12.2f      %12.2f      %12.2f\n", v.Field(k).Interface(), v.Field(k+1).Interface(), v.Field(k+2).Interface(), v.Field(k+3).Interface(), v.Field(k+4).Interface())
		tmp.WriteString(r)
	}
	for k := 70; k < 110; k += 5 {
		r := fmt.Sprintf("%10v   %10v    %10v    %10v    %10v\n", v.Field(k).Interface(), v.Field(k+1).Interface(), v.Field(k+2).Interface(), v.Field(k+3).Interface(), v.Field(k+4).Interface())
		tmp.WriteString(r)
	}

	tmp.WriteString(fmt.Sprintf("%v     %v\n", h.Kstnm, h.Kevnm))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Khole, h.Ko, h.Ka))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Kt0, h.Kt1, h.Kt2))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Kt3, h.Kt4, h.Kt5))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Kt6, h.Kt6, h.Kt8))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Kt9, h.Kf, h.KUser0))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.Kuser1, h.Kuser2, h.Kcmpnm))
	tmp.WriteString(fmt.Sprintf("%v     %v     %v\n", h.KnetwK, h.Kdatrd, h.Kinst))
	return tmp.String()
}
