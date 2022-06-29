package main

import (
	"fmt"
	"github.com/peterh/liner"
	"github.com/vintingb/sacgo/sacio"
	"github.com/vintingb/sacgo/sacpic"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot/vg"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	historyFn   = filepath.Join(os.TempDir(), ".liner_history")
	completions []string
)

func init() {
	updateFileList()
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Print("not find config")
		return
	}

	sacpic.FileType = cfg.Section("").Key("FileType").In("pdf", []string{
		".eps", ".jpg", ".jpeg", ".pdf", ".png", ".svg", ".tex", ".tif", ".tiff",
	})
	sacpic.PicHeight = vg.Length(cfg.Section("").Key("PicHeight").RangeFloat64(5, 5, 20))
	sacpic.PicWeight = vg.Length(cfg.Section("").Key("PicWeight").RangeFloat64(20, 10, 40))
	sacpic.LineSize = cfg.Section("").Key("LineSize").RangeFloat64(0.1, 0.01, 2)
	log.Printf("\nPicHeight: %v\nPicWeight: %v\nLineSize: %v", sacpic.PicHeight, sacpic.PicWeight, sacpic.LineSize)
}

func main() {
	var sacHead []sacio.SacHead
	var sacData []sacio.SacData
	var inputFileNames []string
	line := liner.NewLiner()
	defer func(line *liner.State) {
		err := line.Close()
		if err != nil {

		}
	}(line)
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
		err := f.Close()
		if err != nil {
			return
		}
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
				case "fft":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					for index := range sacData {
						sacData[index].Fft()
					}
				case "rmean":
					for index := range sacHead {
						var mean = floats.Sum(sacData[index].Data) / float64(len(sacData[index].Data))
						for i := range sacData[index].Data {
							sacData[index].Data[i] -= mean
						}
					}
				case "kill", "k", "quit", "q", "exit":
					loop = false
				case "read", "r":
					err := readSACFile(&inputFileNames, arguments, &sacHead, &sacData)
					if err != nil {
						fmt.Println(err)
						continue
					}
					for index := range sacHead {
						sacHead[index].Depmax = float32(floats.Max(sacData[index].Data))
						sacHead[index].Depmin = float32(floats.Min(sacData[index].Data))
						sacHead[index].Depmen = float32(floats.Sum(sacData[index].Data) / float64(len(sacData[index].Data)))
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
							case sacio.K.K:
								hs := tmp.Format()
								immutable := reflect.ValueOf(hs).Elem()
								if headString == sacio.K.Kzdate {
									fmt.Println(kzdate(tmp.Nzyear, tmp.Nzjday))
									continue
								}
								if headString == sacio.K.Kztime {
									fmt.Println(kztime(tmp.Nzhour, tmp.Nzmin, tmp.Nzsec, tmp.Nzmsec))
									continue
								}
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								fmt.Printf("   %s = %v\n", headString, val)
							case sacio.I.I:
								val := immutable.FieldByName(headString)
								if val.Kind() == reflect.Invalid {
									fmt.Printf("no found %s in SACHead\n", headString)
									continue
								}
								v := int32(val.Int())
								if v == sacio.IDefault {
									continue
								}
								switch headString {
								case sacio.I.Iftype:
									fmt.Printf("   %s = %v\n", headString, sacio.Iftype(v))
								case sacio.I.Idep:
									fmt.Printf("   %s = %v\n", headString, sacio.Idep(v))
								case sacio.I.Iztype:
									fmt.Printf("   %s = %v\n", headString, sacio.Iztype(v))
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
								log.Println(headString)
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
				case string('p'), "plot":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					sacpic.Plot(inputFileNames, sacHead, sacData)
					fmt.Println("Picture saved successfully")
				case "p1", "plot1":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					sacpic.Plot1(inputFileNames, sacHead, sacData)
					fmt.Println("Picture saved successfully")
					sacpic.OpenBrowser("P1" + sacpic.FileType)
				case "p2", "plot2":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					sacpic.Plot2(inputFileNames, sacHead, sacData)
					fmt.Println("Picture saved successfully")
					sacpic.OpenBrowser("P2" + sacpic.FileType)
				case "plotfft", "pf":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					sacpic.PlotFft(inputFileNames, sacHead, sacData)
					fmt.Println("Picture saved successfully")
				case "plotfft1", "pf1":
					if len(sacHead) == 0 {
						fmt.Println("no data, please read SAC file first")
						continue
					}
					sacpic.PlotFft1(inputFileNames, sacHead, sacData)
					sacpic.OpenBrowser("PF1" + sacpic.FileType)
				default:
					fmt.Println("Command not currently supported")
				}
			}
		}
		if f, err := os.Create(historyFn); err != nil {
			log.Print("Error writing history file: ", err)
		} else {
			_, _ = line.WriteHistory(f)
			err := f.Close()
			if err != nil {
				return
			}
		}

	}

}
