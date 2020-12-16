/*
@Time : 2020/12/13 23:33
@Author : vintingb
@File : setting
@Software: GoLand
*/

package sacio

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

type SacData struct {
	Data []float32
}

type Sac struct {
	sacHead SacHead
	sacData SacData
}
