package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sacgo/sacio"
	"strings"
)

func isFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArray := []rune(str)
	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}
	return string(strArray)
}

func readSACFile(inputFileNames *[]string, arguments []string, sacHead *[]sacio.SacHead, sacData *[]sacio.SacData) error {
	if arguments[0] == "*.sac" || arguments[0] == "*.SAC" {
		updateFileList()
		for _, inputFileName := range completions {
			if strings.HasSuffix(inputFileName, ".SAC") || strings.HasSuffix(inputFileName, ".sac") {
				*inputFileNames = append(*inputFileNames, inputFileName)
				tmpHead := new(sacio.SacHead)
				tmpData := new(sacio.SacData)
				err := tmpHead.ReadHead(inputFileName)
				if err != nil {
					return err
				}
				err = tmpData.ReadData(inputFileName)
				if err != nil {
					return err
				}
				*sacHead = append(*sacHead, *tmpHead)
				*sacData = append(*sacData, *tmpData)
			}
		}
	} else {
		*inputFileNames = append(*inputFileNames, arguments...)
		for _, inputFileName := range arguments {
			if !isFileExist(inputFileName) {
				panic("File Not Found")
			}
			tmpHead := new(sacio.SacHead)
			tmpData := new(sacio.SacData)
			err := tmpHead.ReadHead(inputFileName)
			if err != nil {
				return err
			}
			err = tmpData.ReadData(inputFileName)
			if err != nil {
				return err
			}
			*sacHead = append(*sacHead, *tmpHead)
			*sacData = append(*sacData, *tmpData)
		}
	}
	return nil
}

func kzdate(nzyear, nzjday int32) string {
	mons := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	days := []int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (nzyear%4 == 0 && nzyear%100 != 0) || nzyear%400 == 0 {
		days[1] = 29
	}
	var daySum int32
	for index, value := range days {
		daySum += value
		if daySum > nzjday {
			daySum -= value
			return fmt.Sprintf("   kzdate = %v %v (%v), %v", mons[index], nzjday-daySum, nzjday, nzyear)
		}

	}
	return ""
}

func kztime(nzhour, nzmin, nzsec, nzmsec int32) string {
	return fmt.Sprintf("   kztime = %v:%v:%v.%v", nzhour, nzmin, nzsec, nzmsec)

}

func updateFileList() {
	completions = []string{}
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fileInfoList {
		completions = append(completions, fileInfoList[i].Name())
	}
}
