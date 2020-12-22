package main

import (
	"fmt"
	"github.com/peterh/liner"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sacgo/sacio"
	"sacgo/sacpic"
	"strings"
)

var (
	historyFn   = filepath.Join(os.TempDir(), ".liner_history")
	completions []string
)

func init() {
	updateFileList()
}

func main() {
	var sacHead []sacio.SacHead
	var sacData []sacio.SacData
	var inputFileNames []string
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		updateFileList()
		lines := strings.Fields(line)
		s := lines[len(lines)-1]
		line = line[:len(line)-len(s)]
		for _, n := range completions {
			if strings.HasPrefix(n, strings.ToLower(s)) {
				c = append(c, line+n)
			}
			if strings.HasPrefix(n, strings.ToUpper(s)) {
				c = append(c, line+n)
			}
		}
		return
	})

	if f, err := os.Open(historyFn); err == nil {
		_, _ = line.ReadHistory(f)
		f.Close()
	}
	loop := true
	for loop {
		if command, err := line.Prompt("sacgo--> "); err == nil {
			line.AppendHistory(command)
			commands := strings.Split(command, ";")
			for index := 0; index < len(commands); index++ {
				optionAndArguments := strings.Fields(commands[index])
				option := optionAndArguments[0]
				arguments := optionAndArguments[1:]
				switch option {
				case "kill", "k", "quit", "q":
					loop = false
				case "read", "r":
					err := readSACFile(&inputFileNames, arguments, &sacHead, &sacData)
					if err != nil {
						fmt.Println(err)
						continue
					}
				case "lh":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					for k, tmp := range sacHead {
						fmt.Printf("file: %s\n--------------------\n", inputFileNames[k])
						for _, headString := range arguments {
							headString := strFirstToUpper(strings.ToLower(headString))
							immutable := reflect.ValueOf(&tmp).Elem()
							switch headString[0] {
							case 'K':
								hs := tmp.HeadToString()
								immutable := reflect.ValueOf(hs).Elem()
								if headString == "Kzdate" {
									fmt.Println(kzdate(tmp.Nzyear, tmp.Nzjday))
									continue
								}
								if headString == "Kztime" {
									fmt.Println(kztime(tmp.Nzhour, tmp.Nzmin, tmp.Nzsec, tmp.Nzmsec))
									continue
								}
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								fmt.Printf("   %s = %v\n", headString, val)
							case 'I':
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								v := int32(val.Int())
								if v == sacio.IDefault {
									continue
								}
								if headString == "Iftype" {
									fmt.Printf("   %s = %v\n", headString, sacio.IFTYPE[v])
								}
								if headString == "Idep" {
									fmt.Printf("   %s = %v\n", headString, sacio.IDEP[v])
								}
								if headString == "Iztype" {
									fmt.Printf("   %s = %v\n", headString, sacio.IZTYPE[v])
								}
							case 'L':
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								v := int32(val.Int())
								if v == sacio.LDefault {
									fmt.Printf("   %s = %v\n", headString, false)
									continue
								} else {
									fmt.Printf("   %s = %v\n", headString, true)
									continue
								}
							default:
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								fmt.Printf("   %s = %v\n", headString, val)
							}
						}
					}
				case "w", "write":
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
				case "ls":
					updateFileList()
					for _, v := range completions {
						fmt.Printf("%v ", v)
					}
					fmt.Println()
				case string('p'), "plot1", "p1":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					for k := range inputFileNames {
						sacpic.Wg.Add(1)
						picData := new(sacpic.PicData)
						picData.NewPicData(&sacHead[k], &sacData[k])
						inputFileName := strings.Split(inputFileNames[k], ".SAC")
						if len(inputFileName) == 1 {
							inputFileName = strings.Split(inputFileNames[k], ".sac")
						}
						go picData.SavePic(inputFileName[0])
					}
					sacpic.Wg.Wait()
					fmt.Println("Picture saved successfully")
				default:
					fmt.Println("Command not currently supported")
				}
			}
		}
		if f, err := os.Create(historyFn); err != nil {
			log.Print("Error writing history file: ", err)
		} else {
			_, _ = line.WriteHistory(f)
			f.Close()
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
