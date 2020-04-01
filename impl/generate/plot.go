package generate

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func (samples SamplesType) plotSeries() {

}

func (samples SamplesType) plotHistogram() {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"

	h, err := plotter.NewHist(plotter.Values(samples), Slices)
	if err != nil {
		panic(err)
	}

	//h.Normalize(1)
	p.Add(h)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, Output); err != nil {
		panic(err)
	}

}

func (samples SamplesType) Plot() {
	if Series {
		samples.plotSeries()
	} else {
		samples.plotHistogram()
	}
}
