package generate

import (
	"math/rand"
)

func Uniform(minimum, maximum float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)
	reqint := maximum - minimum

	rand.Seed(int64(Seed))
	for idx, _ := range samples {
		next := rand.Float64() * reqint
		samples[idx] = minimum + next
	}
	return samples
}
