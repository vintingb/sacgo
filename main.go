package main

import (
	"fmt"
	"github.com/peterh/liner"
	"os"
	"reflect"
	"sacgo/sacio"
	"strings"
)

func main() {
	var sacHead []sacio.SacHead
	var sacData []sacio.SacData
	var inputFileNames []string
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	loop := true
	for loop {
		if command, err := line.Prompt("sacgo--> "); err == nil {
			commands := strings.Split(command, ";")
			for index := 0; index < len(commands); index++ {
				optionAndArguments := strings.Fields(commands[index])
				option := optionAndArguments[0]
				arguments := optionAndArguments[1:]
				switch option {
				case "kill":
					loop = false
				case "k":
					loop = false
				case "read":
					inputFileNames = arguments
					for _, inputFileName := range arguments {
						if !isFileExist(inputFileName) {
							panic("File Not Found")
						}
						tmpHead := new(sacio.SacHead)
						tmpData := new(sacio.SacData)
						err := tmpHead.ReadHead(inputFileName)
						if err != nil {
							fmt.Println(err)
							continue
						}
						err = tmpData.ReadData(inputFileName)
						if err != nil {
							fmt.Println(err)
							continue
						}
						sacHead = append(sacHead, *tmpHead)
						sacData = append(sacData, *tmpData)
					}
				case "r":
					inputFileNames = arguments
					for _, inputFileName := range arguments {
						if !isFileExist(inputFileName) {
							panic("File Not Found")
						}
						tmpHead := new(sacio.SacHead)
						tmpData := new(sacio.SacData)
						err := tmpHead.ReadHead(inputFileName)
						if err != nil {
							fmt.Println(err)
							continue
						}
						err = tmpData.ReadData(inputFileName)
						if err != nil {
							fmt.Println(err)
							continue
						}
						sacHead = append(sacHead, *tmpHead)
						sacData = append(sacData, *tmpData)
					}
				case "lh":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					for k, tmp := range sacHead {
						for _, headString := range arguments {
							immutable := reflect.ValueOf(&tmp).Elem()
							headString := strFirstToUpper(headString)
							val := immutable.FieldByName(headString)
							if headString[0] == 'K' {
								hs := tmp.HeadToString()
								immutable := reflect.ValueOf(hs).Elem()
								val := immutable.FieldByName(headString)
								fmt.Printf("   fileName:%s : %s = %v\n", inputFileNames[k], headString, val)
							} else {
								fmt.Printf("   fileName:%s : %s = %v\n", inputFileNames[k], headString, val)
							}
						}
					}
				case "w":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					for k, outputFileName := range arguments {
						sac := new(sacio.Sac)
						sac.NewSac(&sacHead[k], &sacData[k])
						err := sac.Sac2asc(outputFileName)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}

	}
}

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
