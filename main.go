package main

import (
	"github.com/ivansukach/binomial-and-geometry-distribution/distributions"
	"github.com/ivansukach/binomial-and-geometry-distribution/generators"
	"github.com/sirupsen/logrus"
	"math"
)

func BiasOfAnEstimator(variates []float64) (float64, float64) {
	n := len(variates)
	logrus.Info("n=", n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += variates[i]
	}
	E := sum / float64(n)
	sum = 0.0
	for i := 0; i < n; i++ {
		sum += math.Pow(variates[i]-E, 2)
	}
	D := sum / float64(n-1)
	return E, D
}

func main() {
	a01 := 296454621
	a02 := 302711857
	c1 := 48840859
	c2 := 37330745
	M := int(math.Pow(2, 31))
	K := 64
	n := 1000
	pGeometry := 0.3
	mBinomial := 4
	pBinomial := 0.2
	logrus.Info("M: ", M)
	aSequence2 := *generators.LinearCongruential(a02, c2, M, n)
	aSequence1SpecialSize := *generators.LinearCongruential(a01, c1, M, n+K)
	sequenceByMacLarenMarsaglia := *generators.MacLarenMarsaglia(aSequence1SpecialSize, aSequence2, K, n)
	aSequence22 := *generators.LinearCongruential(a02, c2, M, n+mBinomial)
	aSequence1SpecialSize2 := *generators.LinearCongruential(a01, c1, M, n+K+mBinomial)
	sequenceByMacLarenMarsaglia2 := *generators.MacLarenMarsaglia(aSequence1SpecialSize2, aSequence22, K, n+mBinomial)
	geometryDistributionVariates := *distributions.GeometryDistributionVariates(pGeometry, sequenceByMacLarenMarsaglia)
	logrus.Info("First 10 variates of geometry distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(geometryDistributionVariates[i])
	}
	logrus.Info("Expected value: ", 1.0/pGeometry)
	logrus.Info("Variance: ", (1-pGeometry)/(pGeometry*pGeometry))
	unbiasedGeometryEV, unbiasedGeometryV := BiasOfAnEstimator(geometryDistributionVariates)
	logrus.Info("Bias of an estimator expected value: ", unbiasedGeometryEV)
	logrus.Info("Bias of an estimator variance: ", unbiasedGeometryV)
	binomialDistributionVariates := *distributions.BinomialDistributionVariates(mBinomial, pBinomial, sequenceByMacLarenMarsaglia2)
	logrus.Info("First 10 variates of binomial distribution")
	for i := 0; i < 10; i++ {
		logrus.Info(binomialDistributionVariates[i])
	}
	logrus.Info("Expected value: ", float64(mBinomial)*pBinomial)
	logrus.Info("Variance: ", float64(mBinomial)*pBinomial*(1-pBinomial))
	unbiasedBinomialEV, unbiasedBinomialV := BiasOfAnEstimator(binomialDistributionVariates)
	logrus.Info("Bias of an estimator expected value: ", unbiasedBinomialEV)
	logrus.Info("Bias of an estimator variance: ", unbiasedBinomialV)
}
