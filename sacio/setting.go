/*
@Time : 2020/12/13 23:33
@Author : vintingb
@File : setting
@Software: GoLand
*/

package sacio

var (
	IFTYPE = map[int32]string{
		1: "ITIME {Time series file}",
		2: "IRLIM {Spectral file---real and imaginary}",
		3: "IAMPH {Spectral file---amplitude and phase}",
		4: "IXY {General x versus y data}",
		5: "IXYZ {General XYZ (3-D) file}",
	}
	IDEP = map[int32]string{
		1: "IUNKN (Unknown)",
		2: "IDISP (Displacement in nm)",
		3: "IVEL (Velocity in nm/sec)",
		4: "IACC (Acceleration in nm/sec/sec)",
		5: "NOTE Others use SI units for IDISP, IVEL, IACC.",
		6: "IVOLTS (Velocity in volts)",
	}
	IZTYPE = map[int32]string{
		1: "IUNKN (Unknown)",
		2: "IB (Begin time)",
		3: "IDAY (Midnight of refernece GMT day)",
		4: "IO (Event origin time)",
		5: "IA (First arrival time)",
		6: "ITn (User defined time pick n, n=0,9)",
	}
	FDefault float32 = -12345.0
	NDefault int32   = -12345
	IDefault int32   = -12345 // 逻辑型
	LDefault int32   = 0      // FALSE
	KDefault string  = "-12345  "
)
