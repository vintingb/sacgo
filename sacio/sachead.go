package sacio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func (h *SacHead) ReadHead(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
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
