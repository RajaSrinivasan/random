package generate

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
)

const uintmax uint64 = 0xffffffffffffffff
const uintfloatmax float64 = float64(uintmax)

func CryptoRandom() float64 {
	ivar := make([]uint8, 8)
	_, err := crand.Read(ivar)
	if err != nil {
		log.Fatal(err)
	}

	var intvar uint64
	intvar = binary.LittleEndian.Uint64(ivar)
	rndflt := float64(intvar) / uintfloatmax
	return rndflt
}

func CryptoUniform(minimum, maximum float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)
	reqint := maximum - minimum

	ivar := make([]uint8, 8)

	var intvar uint64

	for idx, _ := range samples {
		_, err := crand.Read(ivar)
		if err != nil {
			log.Fatal(err)
		}
		intvar = binary.LittleEndian.Uint64(ivar)
		rndflt := float64(intvar)
		next := (rndflt / uintfloatmax) * reqint
		samples[idx] = minimum + next
	}
	return samples
}
