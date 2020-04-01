package generate

import (
	"math/rand"
)

func Normal(mean, stdev float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)

	rand.Seed(int64(Seed))
	for idx, _ := range samples {
		next := rand.NormFloat64()
		samples[idx] = next*stdev + mean
	}
	return samples
}
