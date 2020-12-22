package generators

import "math"

func MacLarenMarsaglia(aSequence1 []float64, aSequence2 []float64, K int, n int) *[]float64 {
	V := make([]float64, K)
	a := make([]float64, n)
	for i := 0; i < K; i++ {
		V[i] = aSequence1[i]
	}
	for i := 0; i < n; i++ {
		s := int(math.Floor(aSequence2[i] * float64(K)))
		a[i] = V[s]
		V[s] = aSequence1[i+K]
	}
	return &a
}
