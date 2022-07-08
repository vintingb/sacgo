package line

import "image/color"

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
