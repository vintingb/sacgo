/*
@Time : 2020/12/13 23:33
@Author : vintingb
@File : setting
@Software: GoLand
*/

package sacio

type Iftype int32
type Idep int32
type Iztype int32

const (
	ITIME Iftype = iota + 1
	IRLIM
	IAMPH
	IXY
	IXYZ
)

const (
	IUNKN Idep = iota + 1
	IDISP
	IVEL
	IACC
	NOTE
	IVOLTS
)

const (
	IUNKN_z Iztype = iota + 1
	IB
	IDAY
	IO
	IA
	ITn
)

var (
	K = struct {
		K              uint8
		Kzdate, Kztime string
	}{
		'K', "Kzdate", "Kztime",
	}

	I = struct {
		I                    uint8
		Iftype, Idep, Iztype string
	}{
		'I', "Iftype", "Idep", "Iztype",
	}
)

var (
	FDefault float32 = -12345.0
	NDefault int32   = -12345
	IDefault int32   = -12345 // 逻辑型
	LDefault int32   = 0      // FALSE
	KDefault string  = "-12345  "
)

func (i Iftype) String() string {
	switch i {
	case ITIME:
		return "ITIME {Time series file}"
	case IRLIM:
		return "IRLIM {Spectral file---real and imaginary}"
	case IAMPH:
		return "IAMPH {Spectral file---amplitude and phase}"
	case IXY:
		return "IXY {General x versus y data}"
	case IXYZ:
		return "IXYZ {General XYZ (3-D) file}"
	default:
		return ""
	}
}
func (i Idep) String() string {
	switch i {
	case IUNKN:
		return "IUNKN (Unknown)"
	case IDISP:
		return "IDISP (Displacement in nm)"
	case IVEL:
		return "IVEL (Velocity in nm/sec)"
	case IACC:
		return "IACC (Acceleration in nm/sec/sec)"
	case NOTE:
		return "NOTE Others use SI units for IDISP, IVEL, IACC."
	case IVOLTS:
		return "IVOLTS (Velocity in volts)"
	default:
		return ""
	}
}
func (i Iztype) String() string {
	switch i {
	case IUNKN_z:
		return "IUNKN (Unknown)"
	case IB:
		return "IB (Begin time)"
	case IDAY:
		return "IDAY (Midnight of refernece GMT day)"
	case IO:
		return "IO (Event origin time)"
	case IA:
		return "IA (First arrival time)"
	case ITn:
		return "ITn (User defined time pick n, n=0,9)"
	default:
		return ""
	}
}
