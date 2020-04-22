package generate

import (
	"math"
)

var u, v, x, y, q float64

func CryptoNormalVariate(mean, stdev float64) float64 {
	cycles := 100

	for {
		u = CryptoRandom()
		v = 1.7156 * (CryptoRandom() - 0.5)
		x = u - 0.449871
		y = math.Abs(v) + 0.386595
		q = x*x + y*(0.19600*y-0.25472*x)
		if (q > 0.27597) && (q > 0.27846) || v*v > -4*math.Log(u)*u*u {
			cycles = cycles - 1
			if cycles > 0 {
				continue
			} else {
				cycles = cycles - 1
				break
			}
		}
	}
	return mean + stdev*(v/u)
}

func BoxMuller(mean, stdev float64) (float64, float64) {
	u1 := CryptoRandom()
	u2 := CryptoRandom()
	//fmt.Printf("%f %f\n", u1, u2)
	r := math.Sqrt(-2.0 * math.Log(u1))

	z1 := r * math.Cos(2*math.Pi*u2)
	z2 := r * math.Sin(2*math.Pi*u2)
	return z1*stdev + mean, z2*stdev + mean

}
func CryptoNormal(mean, stdev float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)

	for idx, _ := range samples {
		x1, _ := BoxMuller(mean, stdev)
		samples[idx] = x1
	}
	return samples
}
