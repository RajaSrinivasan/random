package generate

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func (samples SamplesType) convert() plotter.XYs {
	pts := make(plotter.XYs, len(samples))
	for i, _ := range pts {
		pts[i].X = float64(i)
		pts[i].Y = samples[i]
	}
	return pts
}

func (samples SamplesType) plotSeries() {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Random Numbers - Series"

	p.X.Label.Text = "No"
	p.Y.Label.Text = "Random No"

	err = plotutil.AddLinePoints(p, samples.convert())
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, Output); err != nil {
		panic(err)
	}
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
