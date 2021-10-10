package sacio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

type SacHead struct {
	Delta, Depmin, Depmax, Scale, Odelta             float32 //0
	B, E, O, A, Internal1                            float32 //20
	T0, T1, T2, T3, T4                               float32 //40
	T5, T6, T7, T8, T9                               float32 //60
	F, Resp0, Resp1, Resp2, Resp3                    float32 //80
	Resp4, Resp5, Resp6, Resp7, Resp8                float32 //100
	Resp9, Stla, Stlo, Stel, Stdp                    float32 //120
	Evla, Evlo, Evel, Evdp, Mag                      float32 //140
	User0, User1, User2, User3, User4                float32 //160
	User5, User6, User7, User8, User9                float32 //180
	Dist, Az, Baz, Gcarc, Internal2                  float32 //200
	Internal3, Depmen, Cmpaz, Cmpinc, Xminimum       float32 //220
	Xmaximum, Yminimum, Ymaximum, Unused1, Unused2   float32 //240
	Unused3, Unused4, Unused5, Unused6, Unused7      float32 //260
	Nzyear, Nzjday, Nzhour, Nzmin, Nzsec             int32   //280
	Nzmsec, Nvhdr, Norid, Nevid, Npts                int32   //300
	Internal, Nwfid, Nxsize, Nysize, Unused8         int32   //320
	Iftype, Idep, Iztype, Unused9, Iinst             int32   //340
	Istreg, Ievreg, Ievtyp, Iqual, Isynth            int32   //360
	Imagtyp, Imagsrc, Unused10, Unused11, Unused12   int32   //380
	Unused13, Unused14, Unused15, Unused16, Unused17 int32   //400
	Leven, Lpspol, Lovrok, Lcalda, Unused18          int32   //420
	Kstnm                                            [8]byte
	Kevnm                                            [16]byte
	Khole, Ko, Ka                                    [8]byte
	Kt0, Kt1, Kt2                                    [8]byte
	Kt3, Kt4, Kt5                                    [8]byte
	Kt6, Kt7, Kt8                                    [8]byte
	Kt9, Kf, KUser0                                  [8]byte
	Kuser1, Kuser2, Kcmpnm                           [8]byte
	KnetwK, Kdatrd, Kinst                            [8]byte
}

type SacHeadString struct {
	Kstnm                  string
	Kevnm                  string
	Khole, Ko, Ka          string
	Kt0, Kt1, Kt2          string
	Kt3, Kt4, Kt5          string
	Kt6, Kt7, Kt8          string
	Kt9, Kf, KUser0        string
	Kuser1, Kuser2, Kcmpnm string
	KnetwK, Kdatrd, Kinst  string
}

func (h *SacHead) ReadHead(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {

		}
	}(fp)
	// 创建byte slice, 以读取bytes.
	dataBytes := make([]byte, unsafe.Sizeof(*h))
	n, err := fp.Read(dataBytes)
	if err != nil {
		return err
	}
	dataBytes = dataBytes[:n]
	// 将bytes转成对应的struct
	_ = binary.Read(bytes.NewBuffer(dataBytes), binary.LittleEndian, h)
	return nil
}

func (h *SacHead) Format() string {
	var tmp string
	v := reflect.ValueOf(*h)
	for k := 0; k < 70; k += 5 {
		r := fmt.Sprintf("%12.2f     %12.2f      %12.2f      %12.2f      %12.2f\n", v.Field(k).Interface(), v.Field(k+1).Interface(), v.Field(k+2).Interface(), v.Field(k+3).Interface(), v.Field(k+4).Interface())
		tmp += r
	}
	for k := 70; k < 110; k += 5 {
		r := fmt.Sprintf("%10v   %10v    %10v    %10v    %10v\n", v.Field(k).Interface(), v.Field(k+1).Interface(), v.Field(k+2).Interface(), v.Field(k+3).Interface(), v.Field(k+4).Interface())
		tmp += r
	}
	tmp += fmt.Sprintf("%v     %v\n", string(h.Kstnm[:]), string(h.Kevnm[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Khole[:]), string(h.Ko[:]), string(h.Ka[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Kt0[:]), string(h.Kt1[:]), string(h.Kt2[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Kt3[:]), string(h.Kt4[:]), string(h.Kt5[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Kt6[:]), string(h.Kt6[:]), string(h.Kt8[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Kt9[:]), string(h.Kf[:]), string(h.KUser0[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.Kuser1[:]), string(h.Kuser2[:]), string(h.Kcmpnm[:]))
	tmp += fmt.Sprintf("%v     %v     %v\n", string(h.KnetwK[:]), string(h.Kdatrd[:]), string(h.Kinst[:]))
	return tmp
}

func (h *SacHead) HeadToString() *SacHeadString {
	headString := new(SacHeadString)
	headString.Kstnm = string(h.Kstnm[:])
	headString.Kevnm = string(h.Kevnm[:])
	headString.Khole = string(h.Khole[:])
	headString.Ko = string(h.Ko[:])
	headString.Ka = string(h.Ka[:])
	headString.Kt0 = string(h.Kt0[:])
	headString.Kt1 = string(h.Kt1[:])
	headString.Kt2 = string(h.Kt2[:])
	headString.Kt3 = string(h.Kt3[:])
	headString.Kt4 = string(h.Kt4[:])
	headString.Kt5 = string(h.Kt5[:])
	headString.Kt6 = string(h.Kt6[:])
	headString.Kt7 = string(h.Kt7[:])
	headString.Kt8 = string(h.Kt8[:])
	headString.Kt9 = string(h.Kt9[:])
	headString.Kf = string(h.Kf[:])
	headString.KUser0 = string(h.KUser0[:])
	headString.Kuser1 = string(h.Kuser1[:])
	headString.Kuser2 = string(h.Kuser2[:])
	headString.Kcmpnm = string(h.Kcmpnm[:])
	headString.KnetwK = string(h.KnetwK[:])
	headString.Kdatrd = string(h.Kdatrd[:])
	headString.Kinst = string(h.Kinst[:])
	return headString
}
