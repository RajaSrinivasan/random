package generate

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
)

const uintmax uint64 = 0xffffffffffffffff
const uintfloatmax float64 = float64(uintmax)

func CryptoUniform(minimum, maximum float64) (samples SamplesType) {
	samples = make(SamplesType, Samples)
	reqint := maximum - minimum

	ivar := make([]uint8, 8)

	var intvar uint64
	fmt.Printf("Creating %d samples", Samples)
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
