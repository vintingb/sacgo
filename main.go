package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sacgo/sacio"
)

func main() {
	var inputFileName string
	var outputFileName string
	var headString string
	flag.StringVar(&inputFileName, "i", "", "输入的SAC文件名")
	flag.StringVar(&outputFileName, "o", "", "输出的SAC文件名,默认为空字符串,不输出")
	flag.StringVar(&headString, "h", "", "是否输出head中的值,默认为不输出")
	flag.Parse()
	if inputFileName != "" {
		if !isFileExist(inputFileName) {
			panic("File Not Found")
		}
	}
	if outputFileName != "" {
		sac := new(sacio.Sac)
		sac.Read(inputFileName)
		err := sac.Sac2asc(outputFileName)
		if err != nil {
			panic(err)
		}
	}
	if headString != "" {
		h := new(sacio.SacHead)
		h.ReadHead("test.SAC")
		immutable := reflect.ValueOf(h).Elem()
		val := immutable.FieldByName(headString)
		if headString[0] == 'K' {
			hs := h.HeadToString()
			immutable := reflect.ValueOf(hs).Elem()
			val := immutable.FieldByName(headString)
			fmt.Println(val)
			return
		}
		fmt.Println(val)
	}
}

func isFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
