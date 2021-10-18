package sacpic

import (
	"fmt"
	"image/color"
	"log"
	"os/exec"
	"runtime"
)

var SoftColors = []color.Color{
	rgb(0, 0, 0),
	rgb(241, 90, 96),
	rgb(122, 195, 106),
	rgb(90, 155, 212),
	rgb(250, 167, 91),
	rgb(158, 103, 171),
	rgb(206, 112, 88),
	rgb(215, 127, 180),
}

func rgb(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func Color(i int) color.Color {
	n := len(SoftColors)
	if i < 0 {
		return SoftColors[i%n+n]
	}
	return SoftColors[i%n]
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

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
