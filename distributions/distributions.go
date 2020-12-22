package distributions

import "math"

func E(ksi func(x float64) float64, n int, a []float64) float64 {
	sum := 0.
	for i := 0; i < n; i++ {
		sum += ksi(a[i])
	}
	return sum / float64(n)
}
func one(x float64) float64 {
	if x <= 0 {
		return 0.0
	}
	return 1.0
}
func BinomialDistributionVariates(m int, p float64, basicVariate []float64) *[]float64 {
	n := len(basicVariate) - m
	binomialVariates := make([]float64, n)
	var tmp float64
	for i := 0; i < n; i++ {
		tmp = 0.0
		for j := 0; j < m; j++ {
			tmp += one(p - basicVariate[j+i])
		}
		binomialVariates[i] = tmp
	}
	return &binomialVariates
}
func GeometryDistributionVariates(p float64, basicVariate []float64) *[]float64 {
	n := len(basicVariate)
	q := 1 - p
	logQ := math.Log(q)
	geometryVariates := make([]float64, n)
	for i := 0; i < n; i++ {
		geometryVariates[i] = math.Ceil(math.Log(basicVariate[i]) / logQ)
	}
	return &geometryVariates
}
