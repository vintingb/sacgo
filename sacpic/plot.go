package sacpic

import (
	"sacgo/sacio"
)

func Plot(fileNames []string, sacHead []sacio.SacHead, sacData []sacio.SacData) {
	var picData = NewPicData(sacHead, sacData)
	SavePic(&picData, fileNames, plot0)
}

func Plot1(fileNames []string, sacHead []sacio.SacHead, sacData []sacio.SacData) {
	var picData = NewPicData(sacHead, sacData)
	SavePic(&picData, fileNames, plot1)
}

func Plot2(fileNames []string, sacHead []sacio.SacHead, sacData []sacio.SacData) {
	var picData = NewPicData(sacHead, sacData)
	SavePic(&picData, fileNames, plot2)
}
