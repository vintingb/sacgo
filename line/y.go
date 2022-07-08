package line

var Delta float64 = 1

type Ys []float64

func (y Ys) Len() int {
	return len(y)
}

func (y Ys) XY(i int) (float64, float64) {
	return float64(i) * Delta, y[i]
}
