package generate

import (
	"fmt"
	"log"
	"os"
)

func (samples SamplesType) Show() {
	if len(Table) >= 1 {
		fmt.Printf("Dumping values to %s\n", Table)
		outfile, err := os.Create(Table)
		if err != nil {
			log.Fatalf("%s", err)
		}
		defer outfile.Close()
		for _, s := range samples {
			fmt.Fprintf(outfile, "%f\n", s)
		}
	}
}
