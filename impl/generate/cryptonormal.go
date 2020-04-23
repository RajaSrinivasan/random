package generate

import (
	"math"
)

// Boxmuller(mean,stdev) (float64,float64)
//   generates normally distributed random variables with the mean and standard deviation
//   provided. Two random variates are returned.
// Reference:
// https://en.wikipedia.org/wiki/Box–Muller_transform
func BoxMuller(mean, stdev float64) (float64, float64) {
	u1 := CryptoRandom()
	u2 := CryptoRandom()

	r := math.Sqrt(-2.0 * math.Log(u1))

	z1 := r * math.Cos(2*math.Pi*u2)
	z2 := r * math.Sin(2*math.Pi*u2)
	return z1*stdev + mean, z2*stdev + mean

}
func CryptoNormal(mean, stdev float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)

	for idx := 0; idx < Samples/2; idx = idx + 1 {
		x1, x2 := BoxMuller(mean, stdev)
		samples[idx*2] = x1
		samples[idx*2+1] = x2
	}
	if Samples%2 != 0 {
		x1, _ := BoxMuller(mean, stdev)
		samples[Samples-1] = x1
	}
	return samples
}
