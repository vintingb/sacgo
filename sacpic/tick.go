package sacpic

import (
	"gonum.org/v1/plot"
	"math"
	"strconv"
)

type DefaultTicks struct{}

func (DefaultTicks) Ticks(min, max float64) []plot.Tick {
	if max <= min {
		panic("illegal range")
	}

	const suggestedTicks = 3

	labels, _, _, magnitude := talbotLinHanrahan(min, max, suggestedTicks, withinData, nil, nil, nil)
	step := math.Pow10(magnitude)
	ticks := make([]plot.Tick, len(labels))
	for i, v := range labels {
		ticks[i] = plot.Tick{Value: v, Label: strconv.FormatFloat(v/step, 'f', -1, 64)}
	}
	return ticks
}
func Magnitude(min, max float64) int {
	_, _, _, magnitude := talbotLinHanrahan(min, max, 3, withinData, nil, nil, nil)
	return magnitude

}
