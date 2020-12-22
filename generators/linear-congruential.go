package generators

import "math"

func LinearCongruential(alphaTempLast int, c1 int, M int, n int) *[]float64 {
	beta := int(math.Max(float64(c1), float64(M-c1)))
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		alphaTemp := (beta * alphaTempLast) - M*(beta*alphaTempLast/M)
		a[i] = float64(alphaTemp) / float64(M)
		alphaTempLast = alphaTemp
	}
	return &a
}
