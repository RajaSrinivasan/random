package generate

import (
	"fmt"
	"log"
	"os"
)

func (samples SamplesType) Show() {
	if len(Output) >= 1 {
		outfile, err := os.Create(Output)
		if err != nil {
			log.Fatalf("%s", err)
		}
		defer outfile.Close()
		for _, s := range samples {
			fmt.Fprintf(outfile, "%f\n", s)
		}
	}
}
